package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	userAPI "github.com/prostasmosta/auth/grpc/internal/api/user"
	userRepository "github.com/prostasmosta/auth/grpc/internal/repository/user"
	userService "github.com/prostasmosta/auth/grpc/internal/service/user"
	grpcUser "github.com/prostasmosta/auth/grpc/pkg/user_v1"
)

const (
	grpcPort = 50051
	dbDSN    = "host=localhost port=54321 dbname=auth user=auth password=pass sslmode=disable"
)

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
	grpcUser.RegisterUserV1Server(s, userAPI.NewServer(userServ))

	log.Printf("server listening at: %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
