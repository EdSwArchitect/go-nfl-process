package reader

import (
	"log"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var weekCounter prometheus.Counter
var weekErrorCounter prometheus.Counter
var playsCounter prometheus.Counter
var playsErrorCounter prometheus.Counter
var gamesCounter prometheus.Counter
var gamesErrorCounter prometheus.Counter

func init() {

	weekCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "nfl",
			Name:      "weeks_records_in",
			Help:      "The number of records received from the weeks topic",
		})

	playsCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "nfl",
			Name:      "plays_records_in",
			Help:      "The number of records received from the plays topic",
		})

	gamesCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "nfl",
			Name:      "games_records_in",
			Help:      "The number of records received from the games topic",
		})

	playsErrorCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "nfl",
			Name:      "plays_kafka_consumer_errors",
			Help:      "The number of records received from the plays topic",
		})

	weekErrorCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "nfl",
			Name:      "weeks_kafka_consumer_errors",
			Help:      "The number of records received from the weeks topic",
		})

	gamesErrorCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "nfl",
			Name:      "games_kafka_consumer_errors",
			Help:      "The number of records received from the weeks topic",
		})
}

// ReadTopic the topic consumers
func ReadTopic(broker string, topics []string) {

	config := sarama.NewConfig()
	config.ClientID = "go-kafka-consumer"
	config.Consumer.Return.Errors = true

	brokers := []string{broker}

	// Create new consumer
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Printf("Unable to create consumer: %+v\n", err)
		panic(err)
	}

	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	consumer, errors := consume(topics, master)

	// Count how many message processed
	msgCount := 0

	// Get signal for finish
	doneCh := make(chan struct{})
	var errorCount int = 0

	go func() {
		for {
			select {

			case msg := <-consumer:
				msgCount++
				log.Printf("Received messages: %s\n", string(msg.Value))
			case consumerError := <-errors:
				msgCount++
				log.Printf("Received consumerError %s. %s. %+v\n", string(consumerError.Topic), string(consumerError.Partition), consumerError.Err)

				errorCount++

				if errorCount >= 100 {
					doneCh <- struct{}{}
				}
			}
		}
	}()

	<-doneCh
	log.Printf("Processed %d messages", msgCount)

}

func consume(topics []string, master sarama.Consumer) (chan *sarama.ConsumerMessage, chan *sarama.ConsumerError) {
	consumers := make(chan *sarama.ConsumerMessage)
	errors := make(chan *sarama.ConsumerError)
	for _, topic := range topics {
		if strings.Contains(topic, "__consumer_offsets") {
			continue
		}
		partitions, _ := master.Partitions(topic)

		for index := range partitions {

			consumer, err := master.ConsumePartition(topic, partitions[index], sarama.OffsetOldest)
			if nil != err {
				log.Printf("Error consuming topic. %v Partitions: %v", topic, partitions)
				panic(err)
			}

			log.Printf(" Start consuming topic %s\n", topic)

			go func(topic string, consumer sarama.PartitionConsumer) {

				for {
					select {
					case consumerError := <-consumer.Errors():
						errors <- consumerError
						log.Printf("consumerError: %+v\n", consumerError.Err)

						if strings.Contains(topic, "games") {
							gamesErrorCounter.Inc()
						} else if strings.Contains(topic, "play") {
							playsErrorCounter.Inc()
						} else if strings.Contains(topic, "week") {
							weekErrorCounter.Inc()
						}

					case msg := <-consumer.Messages():
						consumers <- msg
						log.Printf("Got message on topic %s. %v", topic, string(msg.Value))

						if strings.Contains(topic, "games") {
							gamesCounter.Inc()
						} else if strings.Contains(topic, "play") {
							playsCounter.Inc()
						} else if strings.Contains(topic, "week") {
							weekCounter.Inc()
						}

					}
				}
			}(topic, consumer)

		}
	}

	return consumers, errors
}
