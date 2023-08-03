package main

import (
	"encoding/json"
	"log"

	"TomAI.QueueConsumer/dataaccess"
	"TomAI.QueueConsumer/domain"
)

func main() {
	var rabbit dataaccess.RabbitMQ
	rabbit.Init("amqp://guest:guest@localhost/", "ratebeerreview")
	msgs, err := rabbit.Consume()

	if err != nil {
		log.Fatal("Connection Error...")
	}

	go func() {
		for d := range msgs {
			// Create a new Person object
			review := domain.BeerReview{}

			// Unmarshal the message into a Person object
			err := json.Unmarshal(d.Body, review)

			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}

			// Now you can access the fields of the message as properties of the person object
			log.Printf("Name: %s", review.Beer.Name)
		}
	}()

	// Don't forget to close the connection and channel when you're done
	defer rabbit.Close()
}
