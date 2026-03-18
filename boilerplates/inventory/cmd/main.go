package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	inventoryAPI "inventory/internal/api/inventory/v1"
	inventoryRepository "inventory/internal/repository/part"
	inventoryService "inventory/internal/service/part"
	inventory "shared/pkg/proto/inventory/v1"
)

const grpcPort = 50051


func main(){
	ctx := context.Background()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return
	}
	defer func() {
		if cerr := lis.Close(); cerr != nil {
			log.Printf("failed to close listener: %v\n", cerr)
		}
	}()

	err = godotenv.Load(".env")
	if err!=nil{
		log.Printf("Load .env failed: %v\n", err)
		return
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI ==""{
		log.Printf("MONGO_URI is nil\n")
		return
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err!= nil{
		log.Printf("Mongo connect failed: %v\n", err)
		return
	}

	defer func(){
		if cerr:=client.Disconnect(ctx); cerr!=nil{
			log.Printf("Disconnect from MongoDB failed: %v", cerr)
		}
	}()

	err = client.Ping(ctx, nil)
	if err!=nil{
		log.Printf("MongoDB is not available: %v", err)
		return
	}
	log.Printf("Successful connection to MongoDB")

	mongoDB := client.Database("inventory")

	s := grpc.NewServer()

	repo := inventoryRepository.NewInventoryRepository(mongoDB)
	service := inventoryService.NewService(repo)
	api := inventoryAPI.NewAPI(service)

	inventory.RegisterInventoryServiceServer(s, api)

	reflection.Register(s)

	go func(){
		log.Printf("🚀 gRPC server listening on %d\n", grpcPort)
		err = s.Serve(lis)
		if err!=nil{
			log.Printf("failed to serve: %v\n", err)
			return
		}
	}()
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down gRPC server...")
	s.GracefulStop()
	log.Println("✅ Server stopped")
}