package main

import (
	"fmt"
	"os"
	"time"

	"github.com/johnkelly/dynamotest/database"
	_ "github.com/joho/godotenv/autoload"
)

type query struct {
	id        string
	question  string
	createdAt int64
	updatedAt int64
}

func main() {
	db := database.New(os.Getenv("AWS_DYNAMO_ACCESS_KEY"), os.Getenv("AWS_DYNAMO_SECRET_KEY"))

	resp, err := db.CreateQuery("Jessica can do what?")

	if err != nil {
		fmt.Printf("CreateQuery Error: %v\n", err.Error())
	}

	q := &query{
		id:        resp.ID,
		question:  resp.Question,
		createdAt: resp.CreatedAt,
		updatedAt: resp.UpdatedAt,
	}

	fmt.Println(q)

	resp, err = db.ShowQuery(q.id)

	if err != nil {
		fmt.Printf("ShowQuery Error: %v\n", err.Error())
	}

	q = &query{
		id:        resp.ID,
		question:  resp.Question,
		createdAt: resp.CreatedAt,
		updatedAt: resp.UpdatedAt,
	}

	fmt.Println(q)

	time.Sleep(2 * time.Second)

	newQuestion := "The computer said to the chicken, wat?"
	timestamp, err := db.UpdateQuery(q.id, "question", newQuestion)
	if err != nil {
		fmt.Printf("UpdateQuery Error: %v\n", err.Error())
	}

	q.question = newQuestion
	q.updatedAt = timestamp

	fmt.Println(q)
	fmt.Println("Checking against db...")

	resp, err = db.ShowQuery(q.id)

	if err != nil {
		fmt.Printf("ShowQuery Error: %v\n", err.Error())
	}

	q = &query{
		id:        resp.ID,
		question:  resp.Question,
		createdAt: resp.CreatedAt,
		updatedAt: resp.UpdatedAt,
	}

	fmt.Println(q)
	fmt.Printf("UpdatedAt matches: %t", (q.updatedAt == timestamp))
}
