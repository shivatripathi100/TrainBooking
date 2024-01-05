package main

import (
	"context"
	"fmt"
	"log"

	"github.com/yourusername/yourproject/protos" // Import your generated protobuf package
	"google.golang.org/grpc"
)

func main() {
	// Set up a gRPC connection to the server
	conn, err := grpc.Dial("localhost:5500", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a client instance using the connection
	client := protos.NewTrainServiceClient(conn)

	// Use client methods to interact with the gRPC server
	// Implement functions to call gRPC methods defined in the service
	// For example:
	purchaseTicket(client)
	viewReceipt(client)
	viewSeatsBySection(client)
	removeUser(client)
	modifySeat(client)

}

func purchaseTicket(client protos.TrainServiceClient) {
	// Implement logic to call the PurchaseTicket gRPC method
	// Use client.PurchaseTicket(context, request) to make the call
	resp, err := client.PurchaseTicket(context.Background(), &protos.TicketRequest{
		From:  "Source",
		To:    "Destination",
		Price: 100,
		User: &protos.User{
			UserId:     "user123",
			FirstName:  "John",
			LastName:   "Doe",
			Email:      "john@example.com",
			Additional: "Additional info",
		},
	})
	if err != nil {
		log.Fatalf("PurchaseTicket failed: %v", err)
	}
	fmt.Println("PurchaseTicket response:", resp)
}

func viewReceipt(client protos.TrainServiceClient) {
	// Implement logic to call the ViewReceipt gRPC method
	// Use client.ViewReceipt(context, request) to make the call
	resp, err := client.ViewReceipt(context.Background(), &protos.UserRequest{
		UserId: "user123",
	})
	if err != nil {
		log.Fatalf("ViewReceipt failed: %v", err)
	}
	fmt.Println("ViewReceipt response:", resp)
}

func viewSeatsBySection(client protos.TrainServiceClient) {
	// Implement logic to call the ViewSeatsBySection gRPC method
	resp, err := client.ViewSeatsBySection(context.Background(), &protos.SectionRequest{
		Section: "YourSectionName", // Replace with the section name you want to query
	})
	if err != nil {
		log.Fatalf("ViewSeatsBySection failed: %v", err)
	}
	fmt.Println("ViewSeatsBySection response:", resp)
}

func removeUser(client protos.TrainServiceClient) {
	// Implement logic to call the RemoveUser gRPC method
	resp, err := client.RemoveUser(context.Background(), &protos.UserRequest{
		User: &protos.User{
			UserId: "user123", // Replace with the user ID you want to remove
		},
	})
	if err != nil {
		log.Fatalf("RemoveUser failed: %v", err)
	}
	fmt.Println("RemoveUser response:", resp)
}

func modifySeat(client protos.TrainServiceClient) {
	// Implement logic to call the ModifySeat gRPC method
	resp, err := client.ModifySeat(context.Background(), &protos.SeatModificationRequest{
		User: &protos.User{
			UserId: "user123", // Replace with the user ID you want to modify the seat for
		},
		NewSeat: "NewSeatNumber", // Replace with the new seat number
	})
	if err != nil {
		log.Fatalf("ModifySeat failed: %v", err)
	}
	fmt.Println("ModifySeat response:", resp)
}
