package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	inventoryAPI "inventory/internal/api/inventory/v1"
	inventoryRepository "inventory/internal/repository/part"
	inventoryService "inventory/internal/service/part"
	inventory "shared/pkg/proto/inventory/v1"
)

const grpcPort = 50051


func main(){
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

	s := grpc.NewServer()

	repo := inventoryRepository.NewInventoryRepository()
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