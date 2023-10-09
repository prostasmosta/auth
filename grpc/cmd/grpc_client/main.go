package main

import (
	"context"
	"log"
	"time"

	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	grpcUser "github.com/prostasmosta/auth/grpc/pkg/user_v1"
)

const (
	address = "localhost:50051"
	userID  = 5
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect ot server: %v", err)
	}
	defer conn.Close()

	c := grpcUser.NewUserV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &grpcUser.GetRequest{Id: userID})
	if err != nil {
		log.Fatalf("failed to get user by id: %v", err)
	}

	log.Printf(color.RedString("User info:\n"), color.GreenString("%+v", r.GetInfo()))
}
