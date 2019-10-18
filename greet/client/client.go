package main

import (
	"context"
	"github.com/mydisha/learn-grpc/greet/model"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := model.NewGreetServiceClient(conn)

	doUnary(client)
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
