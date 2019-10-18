package main

import (
	"context"
	"github.com/mydisha/learn-grpc/greet/model"
	"google.golang.org/grpc"
	"log"
)

type numbers struct {
	FirstNumber  uint64
	SecondNumber uint64
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := model.NewGreetServiceClient(conn)
	sumClient := model.NewSumServiceClient(conn)

	doUnary(client)
	sumTwoNumbers(sumClient, numbers{
		FirstNumber:  10,
		SecondNumber: 20,
	})
}

func sumTwoNumbers(client model.SumServiceClient, payload numbers) {
	request := &model.SumRequest{
		Sum: &model.Sum{
			FirstNumber:          payload.FirstNumber,
			SecondNumber:         payload.SecondNumber,
		},
	}

	resp, err := client.Sum(context.Background(), request)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response %v", resp.Total)
}

func doUnary(client model.GreetServiceClient) {
	request := &model.GreetRequest{
		Greeting: &model.Greeting{
			FirstName: "Dias Taufik",
			LastName:  "Rahman",
		},
	}
	resp, err := client.Greet(context.Background(), request)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response %v", resp.Result)
}
