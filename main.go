package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/EdSwArchitect/go-nfl-process/loader"
	"github.com/EdSwArchitect/go-nfl-process/reader"
	"github.com/Shopify/sarama"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var directory string
var weeksTopic string
var playsTopic string
var gamesTopic string
var bootstrap string
var verbose bool
var doMetrics bool
var readerOnly bool

var producer sarama.AsyncProducer

// var producer sarama.SyncProducer
var kafkaErrorCounter prometheus.Counter
var kafkaSuccessCounter prometheus.Counter

func init() {
	log.Printf("main Init called")

	flag.StringVar(&directory, "dir", ".", "The directory of CSV files")
	flag.StringVar(&weeksTopic, "weeks-topic", "test", "The Kafka topic to put weeks data")
	flag.StringVar(&playsTopic, "plays-topic", "test", "The Kafka topic to put plays data")
	flag.StringVar(&gamesTopic, "games-topic", "test", "The Kafka topic to put games data")
	flag.StringVar(&bootstrap, "bootstrap-server", "localhost:9092", "The Kafka bootstrap servers")
	flag.BoolVar(&verbose, "verbose", true, "Verbose Kafka information")
	flag.BoolVar(&doMetrics, "metrics", true, "Enable Prometheus metrics")
	flag.BoolVar(&readerOnly, "reader-only", true, "Run Kafka consumer")

	flag.Parse()

	log.Printf("Directory: %s\n", directory)
	log.Printf("Weeks topic: %s\n", weeksTopic)
	log.Printf("Plays topic: %s\n", playsTopic)
	log.Printf("Games topic: %s\n", gamesTopic)
	log.Printf("Bootstrap-Server: %s\n", bootstrap)
	log.Printf("Verbose: %t", verbose)
	log.Printf("Metrics enabled: %t", doMetrics)
	log.Printf("Reader only: %t", readerOnly)

	initKafka()

	if doMetrics {
		go initMetrics()
	}

}

func initMetrics() {

	http.Handle("/metrics", promhttp.Handler())

	if !readerOnly {
		http.ListenAndServe(":18080", nil)
	} else {
		http.ListenAndServe(":19080", nil)
	}

	kafkaErrorCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "kafka",
			Name:      "producer_errors",
			Help:      "The number of Kafka producer errors",
		})

	kafkaSuccessCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "kafka",
			Name:      "producer_successes",
			Help:      "The number of Kafka producer successes",
		})

}

func initKafka() {
	if verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	// version, err := sarama.ParseKafkaVersion("2.1.1")
	// if err != nil {
	// 	log.Panicf("Error parsing Kafka version: %v", err)
	// }

	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	config := sarama.NewConfig()
	config.Version = sarama.V0_11_0_2
	config.ClientID = "edwin"
	config.Producer.RequiredAcks = 1
	config.Producer.Flush.MaxMessages = 50
	config.Producer.Retry.Max = 3
	config.Net.TLS.Enable = false
	config.Producer.Return.Successes = false
	config.Producer.Return.Errors = true

	// if doMetrics {
	// 	config.MetricRegistry.
	// }

	brokers := strings.Split(bootstrap, ",")

	var err error

	producer, err = sarama.NewAsyncProducer(brokers, config)
	// producer, err = sarama.NewSyncProducer(brokers, config)

	if err != nil {
		log.Fatalf("Unable to create Kafka producer: %+v", err)
	}

	go func() {
		for err := range producer.Errors() {
			log.Printf("Producer error: %+v", err)
		}
	}()
}

func catcher(signals <-chan os.Signal) {
	for {

		select {
		case <-signals:
			log.Printf("Shutdown requested")

			log.Printf("Close producer returned: %s", producer.Close().Error())

			// producer.AsyncClose()

			os.Exit(0)
		}
	}
}

// func readStuff() {
// 	for {

// 		select {
// 		case /*err := */ <-producer.Errors():
// 			// log.Printf("*** Producer error: %s", err.Err.Error())
// 			kafkaErrorCounter.Inc()

// 		case /* succ := */ <-producer.Successes():
// 			kafkaSuccessCounter.Inc()
// 			// log.Printf("**** Sent %s", succ.Topic)
// 		}
// 	}

// }

func main() {
	log.Printf("Main executing")

	/*
			      - name: POD_NAME
		        valueFrom:
		            fieldRef:
		              fieldPath: metadata.name
		      - name: NODE_NAME
		        valueFrom:
		            fieldRef:
		              fieldPath: spec.nodeName
		      - name: POD_IP
	*/

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go catcher(signals)

	log.Printf("**** POD_NAME: '%s'\n", os.Getenv("POD_NAME"))
	log.Printf("**** NODE_NAME: '%s'\n", os.Getenv("NODE_NAME"))
	log.Printf("**** POD_IP: %s\n", os.Getenv("POD_IP"))

	hostName, err := os.Hostname()

	if err == nil {
		log.Printf("The hostname is: '%s'", hostName)
	} else {
		log.Printf("***** Unable to get hostname: %+v", err)
	}

	ips, err := net.LookupIP(hostName)

	if err == nil {
		for _, ip := range ips {
			log.Printf("Hostname %s lookup IP address: %s", hostName, ip.String())
		}
	} else {
		log.Printf("**** Lookup error: %+v\n", err)
	}

	ips, err = net.LookupIP("kafka01")

	if err == nil {
		for _, ip := range ips {
			log.Printf("Hostname kafka01 lookup IP address: %s", ip.String())
		}
	} else {
		log.Printf("**** Lookup error kafka01: %+v\n", err)
	}

	ips, err = net.LookupIP("kafka-svc")

	if err == nil {
		for _, ip := range ips {
			log.Printf("Hostname kafka-svc lookup IP address: %s", ip.String())
		}
	} else {
		log.Printf("**** Lookup error kafka-svc: %+v\n", err)
	}

	ifaces, err := net.Interfaces()

	if err == nil {
		log.Println("**** IP Addresses ****")
		for _, i := range ifaces {
			addrs, err := i.Addrs()

			if err == nil {

				// handle err
				for _, addr := range addrs {
					var ip net.IP
					switch v := addr.(type) {
					case *net.IPNet:
						ip = v.IP
					case *net.IPAddr:
						ip = v.IP
					}
					// process IP address
					log.Printf("\t***ip address: %s\n", ip.String())

				}
			} else {
				log.Printf("Unable to get i.Addrs: %+v", err)
			}
		}
	} else {
		log.Println("Unable to get net.Interfaces()")
	}

	log.Println("Waiting 1 minute before starting")

	time.Sleep(1 * time.Minute)

	if !readerOnly {

		log.Println("Starting the process. Weeks")

		err = loader.LoadWeeks(directory, producer, weeksTopic)

		if err != nil {
			log.Printf("Error loading weeks data: %+v", err)
			os.Exit(1)
		}

		log.Println("Starting the process. Plays")

		err = loader.LoadPlays(directory, producer, playsTopic)

		if err != nil {
			log.Printf("Error loading plays data: %+v", err)
			os.Exit(1)
		}

		log.Println("Starting the process, Games")

		err = loader.LoadGames(directory, producer, gamesTopic)

		if err != nil {
			log.Printf("Error loading games data: %+v", err)
			os.Exit(1)
		}
	} else {
		var consumerTopics []string

		consumerTopics = make([]string, 3)

		consumerTopics[0] = gamesTopic
		consumerTopics[1] = playsTopic
		consumerTopics[2] = weeksTopic

		reader.ReadTopic(bootstrap, consumerTopics)
	}

	log.Println("Sleeping 20 minutes before exit")

	time.Sleep(20 * time.Minute)

	os.Exit(0)
}
