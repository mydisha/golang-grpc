package main

import (
	"context"
	"fmt"
	"github.com/mydisha/learn-grpc/greet/model"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *model.GreetRequest) (*model.GreetResponse, error) {
	fmt.Printf("Grpc executed with payload %v", req)
	firstName := req.GetGreeting().GetFirstName()
	result := fmt.Sprintf("Hello my name is %s", firstName)

	response := model.GreetResponse{
		Result: result,
	}

	return &response, nil
}

func (*server) Sum(ctx context.Context, req *model.SumRequest) (*model.SumResponse, error) {
	fmt.Printf("Grpc sum executed with payload %v", req)
	firstNumber := req.GetSum().GetFirstNumber()
	secondNumber := req.GetSum().GetSecondNumber()

	total := firstNumber + secondNumber

	response := model.SumResponse{
		Total: total,
	}

	return &response, nil
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	model.RegisterGreetServiceServer(s, &server{})
	model.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
