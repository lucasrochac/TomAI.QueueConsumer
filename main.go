package main

import (
	"encoding/json"
	"fmt"
	"log"

	"TomAI.QueueConsumer/dataaccess"
	"TomAI.QueueConsumer/domain"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var rabbit dataaccess.RabbitMQ
	rabbit.Init("amqp://guest:guest@localhost/", "ratebeerreview")
	msgs, err := rabbit.Consume()

	if err != nil {
		log.Fatal("Connection Error...")
	}

	//var beerReviews []domain.BeerReview

	go func() {
		for d := range msgs {
			review := domain.BeerReview{}
			err := json.Unmarshal(d.Body, &review)

			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
				continue
			}

			//beerReviews = append(beerReviews, review)

			var id, erro = dataaccess.InsertReview(review)
			if erro != nil {
				log.Printf("Error checking beer: %s", err)
				continue
			} else {
				print(id)
			}
		}
	}()

	fmt.Println("Pressione Enter para sair...")
	fmt.Scanln()

	defer rabbit.Close()
}

func print(beerId int) {
	fmt.Printf("| BeerId! #%d\n", beerId)
}
