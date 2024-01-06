package main

import (
	"context"
	"fmt"
	"net"

	"github.com/google/uuid"
	"github.com/trainbooking/train_grpc/train"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type trainServer struct {
	tickets      map[string]*train.TicketReceipt
	seats        map[string]string
	userSeatsMap map[string]*train.UserSeat
}

func NewTrainServer() *trainServer {
	return &trainServer{
		tickets:      make(map[string]*train.TicketReceipt),
		seats:        make(map[string]string),
		userSeatsMap: make(map[string]*train.UserSeat),
	}
}

func generateUniqueTicketID() string {

	id := uuid.New()

	return id.String()
}

func (s *trainServer) PurchaseTicket(ctx context.Context, req *train.TicketRequest) (*train.TicketReceipt, error) {

	if req == nil || req.From == "" || req.To == "" || req.User == nil || req.User.FirstName == "" || req.User.LastName == "" || req.User.Email == "" || req.Price <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid ticket purchase request")
	}
	
        price := float32(20.0)

	ticketID := generateUniqueTicketID()

	ticket := &train.TicketReceipt{
		Id:        ticketID,
		From:      req.From,
		To:        req.To,
		User:      req.User,
		PricePaid: req.Price,
	}

	s.tickets[ticketID] = ticket

	return ticket, nil
}

func (s *trainServer) ViewReceipt(ctx context.Context, req *train.UserRequest) (*train.TicketReceipt, error) {

	if req == nil || req.UserId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user request")
	}

	for _, ticket := range s.tickets {
		if ticket.User != nil && ticket.User.UserId == req.UserId {

			return ticket, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "No ticket receipt found for the given user ID")
}

func (s *trainServer) ViewSeatsBySection(ctx context.Context, req *train.SectionRequest) (*train.SeatList, error) {

	if req == nil || req.Section == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid section request")
	}

	filteredUserSeats := make([]*train.UserSeat, 0)

	for _, userSeat := range s.userSeatsMap {
		if userSeat.Seat == req.Section {
			filteredUserSeats = append(filteredUserSeats, userSeat)
		}
	}

	if len(filteredUserSeats) == 0 {
		return nil, status.Errorf(codes.NotFound, "No seats found for the given section")
	}

	seatList := &train.SeatList{
		UserSeats: filteredUserSeats,
	}

	return seatList, nil
}
func (s *trainServer) RemoveUser(ctx context.Context, req *train.UserRequest) (*train.Empty, error) {

	if req == nil || req.User == nil || req.User.UserId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user request")
	}

	userIDToRemove := req.User.UserId

	for _, userSeat := range s.userSeatsMap {
		if userSeat.User != nil && userSeat.User.UserId == userIDToRemove {
			userSeat.User = nil
		}
	}

	return &train.Empty{}, nil
}
func (s *trainServer) ModifySeat(ctx context.Context, req *train.SeatModificationRequest) (*train.Empty, error) {

	if req == nil || req.User == nil || req.User.UserId == "" || req.NewSeat == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid seat modification request")
	}

	seatNumber := req.NewSeat

	userSeat, seatExists := s.userSeatsMap[seatNumber]
	if !seatExists {
		return nil, status.Errorf(codes.NotFound, "Seat not found")
	}

	userSeat.Seat = req.NewSeat
	userSeat.User = req.User

	return &train.Empty{}, nil
}
func (s *trainServer) mustEmbedUnimplementedTrainServiceServer() {

}

func main() {

	grpcServer := grpc.NewServer()

	server := &trainServer{}

	train.RegisterTrainServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", ":5500")
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}

	fmt.Println("gRPC server is running...")
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Println("Failed to serve:", err)
	}
}
