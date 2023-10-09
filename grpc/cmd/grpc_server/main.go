package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	grpcUser "github.com/prostasmosta/auth/grpc/pkg/user_v1"
)

const grpcPort = 50051

type server struct {
	grpcUser.UnimplementedUserV1Server
}

func (s *server) Create(ctx context.Context, req *grpcUser.CreateRequest) (*grpcUser.CreateResponse, error) {
	fakeId := gofakeit.Int64()
	log.Printf("User id: %d", fakeId)

	return &grpcUser.CreateResponse{
		Id: fakeId,
	}, nil
}

func (s *server) Get(ctx context.Context, req *grpcUser.GetRequest) (*grpcUser.GetResponse, error) {
	log.Printf("User id: %d", req.GetId())

	return &grpcUser.GetResponse{
		Id: req.GetId(),
		Info: &grpcUser.UserInfo{
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
			Role:  0,
		},
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (s *server) Update(ctx context.Context, req *grpcUser.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("User id: %d", req.GetId())

	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *grpcUser.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("User id: %d", req.GetId())

	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	grpcUser.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at: %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
