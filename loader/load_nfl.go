package loader

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/Shopify/sarama"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var weekCounter prometheus.Counter
var weekErrorCounter prometheus.Counter
var playsCounter prometheus.Counter
var playsErrorCounter prometheus.Counter
var gamesCounter prometheus.Counter

func init() {

	weekCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "nfl",
			Name:      "weeks_records_out",
			Help:      "The number of records sent to the weeks topic",
		})

	playsCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "nfl",
			Name:      "plays_records_out",
			Help:      "The number of records sent to the plays topic",
		})

	gamesCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "nfl",
			Name:      "games_records_out",
			Help:      "The number of records sent to the games topic",
		})

	playsErrorCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "nfl",
			Name:      "plays_kafka_errors",
			Help:      "The number of records sent to the plays topic",
		})

	weekErrorCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "nfl",
			Name:      "weeks_kafka_errors",
			Help:      "The number of records sent to the weeks topic",
		})
}

// LoadWeeks load the week CSV files into the topic
func LoadWeeks(directory string, producer sarama.AsyncProducer, topic string) error {
	// func LoadWeeks(directory string, producer sarama.SyncProducer, topic string) error {

	log.Printf("Directory: %s", directory)

	files, err := ioutil.ReadDir(directory)

	regex := regexp.MustCompile(`week\d+\.csv`)

	if err != nil {
		log.Printf("Error reading directory: %s. %+v", directory, err)

		return err
	}

	// gocsv.SetCSVReader(gocsv.DefaultCSVReader)

	for _, f := range files {
		if !f.IsDir() {
			log.Printf("Working with file: %s/%s. %t", directory, f.Name(), regex.Match([]byte(f.Name())))

			if regex.Match([]byte(f.Name())) {
				log.Printf("CSV file %s/%s processing", directory, f.Name())

				csvFile, err := os.Open(fmt.Sprintf("%s/%s", directory, f.Name()))

				if err != nil {
					return err
				}

				defer csvFile.Close()

				r := bufio.NewReader(csvFile)

				csvReader := csv.NewReader(r)

				var line []string
				var week *JWeek
				var i int = 0

				for {
					line, err = csvReader.Read()

					if err != nil && err != io.EOF {
						log.Printf("CSV error: %+v", err)
						return err
					} else if err == io.EOF {
						break
					}

					// log.Printf("The record: %+v", line)
					week = newJWeek(line)

					bytes, err := json.Marshal(week)

					if err != nil {
						log.Printf("JSON Marshal error: %+v", err)
						return err
					}

					// log.Printf("The record: %s\n", string(bytes))
					// key, _ := uuid.NewRandom()

					// producer.SendMessage(&sarama.ProducerMessage{
					// 	Topic: topic,
					// 	Value: sarama.StringEncoder(bytes),
					// })

					producer.Input() <- &sarama.ProducerMessage{
						Topic: topic,
						// Key:   sarama.StringEncoder(key.String()),
						Value: sarama.StringEncoder(bytes),
					}

					weekCounter.Inc()

					i++

					// if (i % 100) == 0 {
					// 	time.Sleep(50 * time.Millisecond)
					// }

					// log.Printf("The record: %+v", *week)

				}

			}
		}
	}

	return nil
}

// LoadPlays load the plays CSV files into the topic
func LoadPlays(directory string, producer sarama.AsyncProducer, topic string) error {
	// func LoadPlays(directory string, producer sarama.SyncProducer, topic string) error {

	log.Printf("Directory: %s", directory)

	files, err := ioutil.ReadDir(directory)

	regex := regexp.MustCompile(`plays\d*\.csv`)

	if err != nil {
		log.Printf("Error reading directory: %s. %+v", directory, err)

		return err
	}

	// gocsv.SetCSVReader(gocsv.DefaultCSVReader)

	for _, f := range files {
		if !f.IsDir() {
			log.Printf("Working with file: %s/%s. %t", directory, f.Name(), regex.Match([]byte(f.Name())))

			if regex.Match([]byte(f.Name())) {
				log.Printf("CSV file %s/%s processing", directory, f.Name())

				csvFile, err := os.Open(fmt.Sprintf("%s/%s", directory, f.Name()))

				if err != nil {
					return err
				}

				defer csvFile.Close()

				r := bufio.NewReader(csvFile)

				csvReader := csv.NewReader(r)

				var line []string
				var plays *JPlays
				var i int = 0

				for {
					line, err = csvReader.Read()

					if err != nil && err != io.EOF {
						log.Printf("CSV error: %+v", err)
						return err
					} else if err == io.EOF {
						break
					}

					// log.Printf("The record: %+v", line)
					plays = NewJPlays(line)

					bytes, err := json.Marshal(plays)

					if err != nil {
						log.Printf("JSON Marshal error: %+v", err)
						return err
					}

					// log.Printf("The record: %s\n", string(bytes))
					// key, _ := uuid.NewRandom()

					producer.Input() <- &sarama.ProducerMessage{
						Topic: topic,
						// Key:   sarama.StringEncoder(key.String()),
						Value: sarama.StringEncoder(bytes),
					}

					// producer.SendMessage(&sarama.ProducerMessage{
					// 	Topic: topic,
					// 	Value: sarama.StringEncoder(bytes),
					// })

					playsCounter.Inc()

					i++

					// if (i % 100) == 0 {
					// 	time.Sleep(50 * time.Millisecond)
					// }

					// log.Printf("The record: %+v", *week)

				}

			}
		}
	}

	return nil
}

// LoadGames load the games CSV files into the topic
// func LoadGames(directory string, producer sarama.SyncProducer, topic string) error {
func LoadGames(directory string, producer sarama.AsyncProducer, topic string) error {

	log.Printf("Directory: %s", directory)

	files, err := ioutil.ReadDir(directory)

	regex := regexp.MustCompile(`games\d*\.csv`)

	if err != nil {
		log.Printf("Error reading directory: %s. %+v", directory, err)

		return err
	}

	// gocsv.SetCSVReader(gocsv.DefaultCSVReader)

	for _, f := range files {
		if !f.IsDir() {
			log.Printf("Working with file: %s/%s. %t", directory, f.Name(), regex.Match([]byte(f.Name())))

			if regex.Match([]byte(f.Name())) {
				log.Printf("CSV file %s/%s processing", directory, f.Name())

				csvFile, err := os.Open(fmt.Sprintf("%s/%s", directory, f.Name()))

				if err != nil {
					return err
				}

				defer csvFile.Close()

				r := bufio.NewReader(csvFile)

				csvReader := csv.NewReader(r)

				var line []string
				var games *JGames
				var i int = 0

				for {
					line, err = csvReader.Read()

					if err != nil && err != io.EOF {
						log.Printf("CSV error: %+v", err)
						return err
					} else if err == io.EOF {
						break
					}

					// log.Printf("The record: %+v", line)
					games = NewJGames(line)

					bytes, err := json.Marshal(games)

					if err != nil {
						log.Printf("JSON Marshal error: %+v", err)
						return err
					}

					// log.Printf("The record: %s\n", string(bytes))
					// key, _ := uuid.NewRandom()

					producer.Input() <- &sarama.ProducerMessage{
						Topic: topic,
						// Key:   sarama.StringEncoder(key.String()),
						Value: sarama.StringEncoder(bytes),
					}

					// _, _, err = producer.SendMessage(&sarama.ProducerMessage{
					// 	Topic: topic,
					// 	Value: sarama.StringEncoder(bytes),
					// })

					if err != nil {
						log.Printf("Error sending data to kafka: %+v\n", err)
					}

					gamesCounter.Inc()

					i++

					// if (i % 100) == 0 {
					// 	time.Sleep(50 * time.Millisecond)
					// }

				}

			}
		}
	}

	return nil
}
