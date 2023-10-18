package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"

	"github.com/prostasmosta/auth/grpc/internal/converter"
	userRepository "github.com/prostasmosta/auth/grpc/internal/repository/user"
	"github.com/prostasmosta/auth/grpc/internal/service"
	userService "github.com/prostasmosta/auth/grpc/internal/service/user"
	grpcUser "github.com/prostasmosta/auth/grpc/pkg/user_v1"
)

const (
	grpcPort = 50051
	dbDSN    = "host=localhost port=54321 dbname=auth user=auth password=pass sslmode=disable"
)

type server struct {
	grpcUser.UnimplementedUserV1Server
	userService service.UserService
}

func (s *server) Create(ctx context.Context, req *grpcUser.CreateRequest) (*grpcUser.CreateResponse, error) {
	id, err := s.userService.Create(ctx, converter.ToCreateUserFromProto(req))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted user with id: %d", id)

	return &grpcUser.CreateResponse{
		Id: id,
	}, nil
}

func (s *server) Get(ctx context.Context, req *grpcUser.GetRequest) (*grpcUser.GetResponse, error) {
	userObj, err := s.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, name: %s, email: %s, role: %v, created_at: %v, updated_at: %v\n",
		userObj.Id, userObj.Info.Name, userObj.Info.Email, userObj.Info.Role, userObj.CreatedAt, userObj.UpdatedAt)

	convertedUser := converter.ToGetUserFromService(userObj)

	return &grpcUser.GetResponse{
		Id: convertedUser.Id,
		Info: &grpcUser.UserInfo{
			Name:  convertedUser.Info.Name,
			Email: convertedUser.Info.Email,
			Role:  convertedUser.Info.Role,
		},
		CreatedAt: convertedUser.CreatedAt,
		UpdatedAt: convertedUser.UpdatedAt,
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

	userRepo := userRepository.NewRepository(pool)
	userServ := userService.NewService(userRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	grpcUser.RegisterUserV1Server(s, &server{userService: userServ})

	log.Printf("server listening at: %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
