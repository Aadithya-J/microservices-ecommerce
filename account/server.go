package account

import (
	"context"
	"net"
	"strconv"

	pb "github.com/Aadithya-J/microservices-ecommerce/account/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	service Service
}

func NewGRPCServer(service Service) *grpc.Server {
	srv := &grpcServer{service: service}
	s := grpc.NewServer()
	pb.RegisterAccountServiceServer(s, srv)
	reflection.Register(s)
	return s
}

func ListenAndServeGRPC(service Service, port int) error {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	s := NewGRPCServer(service)
	return s.Serve(lis)
}

func (s *grpcServer) PostAccount(ctx context.Context, r *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
	if r.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name cannot be empty")
	}

	a, err := s.service.PostAccount(ctx, r.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.PostAccountResponse{
		Account: &pb.Account{
			Id:   a.ID,
			Name: a.Name,
		},
	}, nil
}

func (s *grpcServer) GetAccount(ctx context.Context, r *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	if r.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id cannot be empty")
	}

	a, err := s.service.GetAccount(ctx, r.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "account not found")
	}

	return &pb.GetAccountResponse{
		Account: &pb.Account{
			Id:   a.ID,
			Name: a.Name,
		},
	}, nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
	if r.Take > 100 {
		r.Take = 100
	}

	accs, err := s.service.GetAccounts(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	accounts := make([]*pb.Account, len(accs))
	for i, acc := range accs {
		accounts[i] = &pb.Account{
			Id:   acc.ID,
			Name: acc.Name,
		}
	}

	return &pb.GetAccountsResponse{
		Accounts: accounts,
	}, nil
}
