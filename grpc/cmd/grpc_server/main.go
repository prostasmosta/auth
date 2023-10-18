package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/prostasmosta/auth/grpc/internal/repository"
	"github.com/prostasmosta/auth/grpc/internal/repository/user"
	"log"
	"net"

	grpcUser "github.com/prostasmosta/auth/grpc/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	grpcPort = 50051
	dbDSN    = "host=localhost port=54321 dbname=auth user=auth password=pass sslmode=disable"
)

type server struct {
	grpcUser.UnimplementedUserV1Server
	userRepository repository.UserRepository
}

func (s *server) Create(ctx context.Context, req *grpcUser.CreateRequest) (*grpcUser.CreateResponse, error) {
	id, err := s.userRepository.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	log.Printf("inserteduser with id: %d", id)

	return &grpcUser.CreateResponse{
		Id: id,
	}, nil
}

func (s *server) Get(ctx context.Context, req *grpcUser.GetRequest) (*grpcUser.GetResponse, error) {
	userObj, err := s.userRepository.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, name: %s, email: %s, role: %v, created_at: %v, updated_at: %v\n",
		userObj.Id, userObj.Info.Name, userObj.Info.Email, userObj.Info.Role, userObj.CreatedAt, userObj.UpdatedAt)

	return &grpcUser.GetResponse{
		Id: userObj.Id,
		Info: &grpcUser.UserInfo{
			Name:  userObj.Info.Name,
			Email: userObj.Info.Email,
			Role:  userObj.Info.Role,
		},
		CreatedAt: userObj.CreatedAt,
		UpdatedAt: userObj.UpdatedAt,
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
	ctx := context.Background()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer pool.Close()

	userRepo := user.NewRepository(pool)

	s := grpc.NewServer()
	reflection.Register(s)
	grpcUser.RegisterUserV1Server(s, &server{userRepository: userRepo})

	log.Printf("server listening at: %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
