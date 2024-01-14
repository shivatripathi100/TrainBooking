package main

import (
	"context"
	"fmt"
	"log"

	"github.com/trainbooking/train_grpc/train"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:5500", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := train.NewTrainServiceClient(conn)

	purchaseTicket(client)
	viewReceipt(client)
	viewSeatsBySection(client)
	removeUser(client)
	modifySeat(client)
	mustEmbedUnimplementedTrainServiceServer()

}

func purchaseTicket(client train.TrainServiceClient) {

	resp, err := client.PurchaseTicket(context.Background(), &train.TicketRequest{
		From:  "London",
		To:    "Paris",
		Price: 20,
		User: &train.User{
			UserId:    "user123",
			FirstName: "Shiva",
			LastName:  "Tripathi",
			Email:     "Shiva@example.com",
		},
	})
	if err != nil {
		log.Fatalf("PurchaseTicket failed: %v", err)
	}
	fmt.Println("PurchaseTicket response:", resp)
}

func viewReceipt(client train.TrainServiceClient) {

	resp, err := client.ViewReceipt(context.Background(), &train.UserRequest{
		UserId: "user123",
	})
	if err != nil {
		log.Fatalf("ViewReceipt failed: %v", err)
	}
	fmt.Println("ViewReceipt response:", resp)
}

func viewSeatsBySection(client train.TrainServiceClient) {

	resp, err := client.ViewSeatsBySection(context.Background(), &train.SectionRequest{
		Section: "YourSectionName",
	})
	if err != nil {
		log.Fatalf("ViewSeatsBySection failed: %v", err)
	}
	fmt.Println("ViewSeatsBySection response:", resp)
}

func removeUser(client train.TrainServiceClient) {

	resp, err := client.RemoveUser(context.Background(), &train.UserRequest{
		User: &train.User{
			UserId: "user123",
		},
	})
	if err != nil {
		log.Fatalf("RemoveUser failed: %v", err)
	}
	fmt.Println("RemoveUser response:", resp)
}

func modifySeat(client train.TrainServiceClient) {

	resp, err := client.ModifySeat(context.Background(), &train.SeatModificationRequest{
		User: &train.User{
			UserId: "user123",
		},
		NewSeat: "NewSeatNumber",
	})
	if err != nil {
		log.Fatalf("ModifySeat failed: %v", err)
	}
	fmt.Println("ModifySeat response:", resp)
}

func mustEmbedUnimplementedTrainServiceServer() {

}
