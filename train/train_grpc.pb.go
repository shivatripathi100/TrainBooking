package train

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion7

const (
	TrainService_PurchaseTicket_FullMethodName     = "/TrainService/PurchaseTicket"
	TrainService_ViewReceipt_FullMethodName        = "/TrainService/ViewReceipt"
	TrainService_ViewSeatsBySection_FullMethodName = "/TrainService/ViewSeatsBySection"
	TrainService_RemoveUser_FullMethodName         = "/TrainService/RemoveUser"
	TrainService_ModifySeat_FullMethodName         = "/TrainService/ModifySeat"
)


type TrainServiceClient interface {
	PurchaseTicket(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*TicketReceipt, error)
	ViewReceipt(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*TicketReceipt, error)
	ViewSeatsBySection(ctx context.Context, in *SectionRequest, opts ...grpc.CallOption) (*SeatList, error)
	RemoveUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Empty, error)
	ModifySeat(ctx context.Context, in *SeatModificationRequest, opts ...grpc.CallOption) (*Empty, error)
}

type trainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrainServiceClient(cc grpc.ClientConnInterface) TrainServiceClient {
	return &trainServiceClient{cc}
}

func (c *trainServiceClient) PurchaseTicket(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*TicketReceipt, error) {
	out := new(TicketReceipt)
	err := c.cc.Invoke(ctx, TrainService_PurchaseTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) ViewReceipt(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*TicketReceipt, error) {
	out := new(TicketReceipt)
	err := c.cc.Invoke(ctx, TrainService_ViewReceipt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) ViewSeatsBySection(ctx context.Context, in *SectionRequest, opts ...grpc.CallOption) (*SeatList, error) {
	out := new(SeatList)
	err := c.cc.Invoke(ctx, TrainService_ViewSeatsBySection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) RemoveUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, TrainService_RemoveUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) ModifySeat(ctx context.Context, in *SeatModificationRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, TrainService_ModifySeat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type TrainServiceServer interface {
	PurchaseTicket(context.Context, *TicketRequest) (*TicketReceipt, error)
	ViewReceipt(context.Context, *UserRequest) (*TicketReceipt, error)
	ViewSeatsBySection(context.Context, *SectionRequest) (*SeatList, error)
	RemoveUser(context.Context, *UserRequest) (*Empty, error)
	ModifySeat(context.Context, *SeatModificationRequest) (*Empty, error)
	mustEmbedUnimplementedTrainServiceServer()
}


type UnimplementedTrainServiceServer struct {
}

func (UnimplementedTrainServiceServer) PurchaseTicket(context.Context, *TicketRequest) (*TicketReceipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseTicket not implemented")
}
func (UnimplementedTrainServiceServer) ViewReceipt(context.Context, *UserRequest) (*TicketReceipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewReceipt not implemented")
}
func (UnimplementedTrainServiceServer) ViewSeatsBySection(context.Context, *SectionRequest) (*SeatList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewSeatsBySection not implemented")
}
func (UnimplementedTrainServiceServer) RemoveUser(context.Context, *UserRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedTrainServiceServer) ModifySeat(context.Context, *SeatModificationRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySeat not implemented")
}
func (UnimplementedTrainServiceServer) mustEmbedUnimplementedTrainServiceServer() {}


type UnsafeTrainServiceServer interface {
	mustEmbedUnimplementedTrainServiceServer()
}

func RegisterTrainServiceServer(s grpc.ServiceRegistrar, srv TrainServiceServer) {
	s.RegisterService(&TrainService_ServiceDesc, srv)
}

func _TrainService_PurchaseTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).PurchaseTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_PurchaseTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).PurchaseTicket(ctx, req.(*TicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_ViewReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).ViewReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_ViewReceipt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).ViewReceipt(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_ViewSeatsBySection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).ViewSeatsBySection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_ViewSeatsBySection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).ViewSeatsBySection(ctx, req.(*SectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_RemoveUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).RemoveUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_ModifySeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeatModificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).ModifySeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainService_ModifySeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).ModifySeat(ctx, req.(*SeatModificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var TrainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TrainService",
	HandlerType: (*TrainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PurchaseTicket",
			Handler:    _TrainService_PurchaseTicket_Handler,
		},
		{
			MethodName: "ViewReceipt",
			Handler:    _TrainService_ViewReceipt_Handler,
		},
		{
			MethodName: "ViewSeatsBySection",
			Handler:    _TrainService_ViewSeatsBySection_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _TrainService_RemoveUser_Handler,
		},
		{
			MethodName: "ModifySeat",
			Handler:    _TrainService_ModifySeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "train.proto",
}
