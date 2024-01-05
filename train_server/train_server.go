package main

import (
	"context"
	"fmt"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type trainServer struct {
	tickets      map[string]*protos.TicketReceipt
	seats        map[string]string
	userSeatsMap map[string]*protos.UserSeat
}

func NewTrainServer() *trainServer {
	return &trainServer{
		tickets:      make(map[string]*protos.TicketReceipt),
		seats:        make(map[string]string),
		userSeatsMap: make(map[string]*protos.UserSeat),
	}
}

func generateUniqueTicketID() string {
	// Generate a new UUID
	id := uuid.New()
	// Convert UUID to string representation
	return id.String()
}

func (s *trainServer) PurchaseTicket(ctx context.Context, req *protos.TicketRequest) (*protos.TicketReceipt, error) {
	// Assuming you have a storage mechanism for ticket receipts
	tickets := make(map[string]*protos.TicketReceipt)

	// Validate the request data
	if req == nil || req.From == "" || req.To == "" || req.User == nil || req.User.FirstName == "" || req.User.LastName == "" || req.User.Email == "" || req.Price <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid ticket purchase request")
	}

	// Logic to generate a unique ticket ID (You can implement this function)
	ticketID := generateUniqueTicketID()

	// Create the ticket receipt
	ticket := &protos.TicketReceipt{
		Id:        ticketID,
		From:      req.From,
		To:        req.To,
		User:      req.User,
		PricePaid: req.Price,
	}

	// Store the ticket in the map
	tickets[ticketID] = ticket

	return ticket, nil
}

func (s *trainServer) ViewReceipt(ctx context.Context, req *protos.UserRequest) (*protos.TicketReceipt, error) {
	// Validate the request data
	if req == nil || req.UserId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user request")
	}

	// Assuming tickets are stored in a map with ticket IDs as keys
	// Check if the requested user ID matches any ticket in the map
	for _, ticket := range s.tickets {
		if ticket.User != nil && ticket.User.UserId == req.UserId {
			// If the user ID matches, return the ticket receipt
			return ticket, nil
		}
	}

	// If no matching ticket is found, return an error
	return nil, status.Errorf(codes.NotFound, "No ticket receipt found for the given user ID")
}

func (s *trainServer) ViewSeatsBySection(ctx context.Context, req *protos.SectionRequest) (*protos.SeatList, error) {
	// Validate the request data
	if req == nil || req.Section == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid section request")
	}

	// Initialize the filtered user seats
	filteredUserSeats := make([]*protos.UserSeat, 0)

	// Iterate over the userSeatsMap and filter based on the section
	for _, userSeat := range s.userSeatsMap {
		if userSeat.Seat == req.Section {
			filteredUserSeats = append(filteredUserSeats, userSeat)
		}
	}

	// If no seats are found for the given section, return an error
	if len(filteredUserSeats) == 0 {
		return nil, status.Errorf(codes.NotFound, "No seats found for the given section")
	}

	// Create a SeatList containing the filtered seats
	seatList := &protos.SeatList{
		UserSeats: filteredUserSeats,
	}

	return seatList, nil
}
func (s *trainServer) RemoveUser(ctx context.Context, req *protos.UserRequest) (*protos.Empty, error) {
	// Validate the request data
	if req == nil || req.User == nil || req.User.UserId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user request")
	}

	// Get the UserID to remove
	userIDToRemove := req.User.UserId

	// Iterate over the userSeatsMap and remove the user from all seats
	for _, userSeat := range s.userSeatsMap {
		if userSeat.User != nil && userSeat.User.UserId == userIDToRemove {
			userSeat.User = nil // Removing the user from the seat
		}
	}

	// Assuming you want to clear other user-related data or perform additional cleanup

	// Return an empty response as confirmation
	return &protos.Empty{}, nil
}
func (s *trainServer) ModifySeat(ctx context.Context, req *protos.SeatModificationRequest) (*protos.Empty, error) {
	// Validate the request data
	if req == nil || req.User == nil || req.User.UserId == "" || req.NewSeat == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid seat modification request")
	}

	// Assuming you have a map of seats where the key is the seat number or identifier
	seatNumber := req.NewSeat

	// Check if the seat exists in your data structure
	userSeat, seatExists := s.userSeatsMap[seatNumber]
	if !seatExists {
		return nil, status.Errorf(codes.NotFound, "Seat not found")
	}

	// Modify the seat according to the request
	userSeat.Seat = req.NewSeat // Update seat number
	userSeat.User = req.User    // Assign the user to the seat

	// Assuming you want to perform other modifications or additional actions related to the seat modification

	// Return an empty response as confirmation
	return &protos.Empty{}, nil
}

// Implement other gRPC service methods similarly

func main() {
	// Initialize your in-memory storage if needed

	// Create a gRPC server
	grpcServer := grpc.NewServer()

	// Create an instance of your gRPC service implementation
	server := &trainServer{}

	// Register your service implementation with the gRPC server
	protos.RegisterTrainServiceServer(grpcServer, server)

	// Start listening on a specific port
	listener, err := net.Listen("tcp", ":5500")
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}

	// Start the gRPC server
	fmt.Println("gRPC server is running...")
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Println("Failed to serve:", err)
	}
}
