package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	orderAPI "order/internal/api/order/v1"
	orderRepository "order/internal/repository/order"
	orderService "order/internal/service/order"
	orderv1 "shared/pkg/openapi/order/v1"
)

const (
	httpPort          = "8080"
	inventoryAddr     = "localhost:50051"
	paymentAddr       = "localhost:50052"
	readHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 10 * time.Second
)


func main() {
	if err := run(); err != nil {
		log.Fatalf("order service failed: %v", err)
	}
}

func run() error {
	invConn, err := grpc.NewClient(inventoryAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer func() {
		if cerr := invConn.Close(); cerr != nil {
			log.Printf("inventory conn close error: %v", cerr)
		}
	}()

	payConn, err := grpc.NewClient(paymentAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer func() {
		if cerr := payConn.Close(); cerr != nil {
			log.Printf("payment conn close error: %v", cerr)
		}
	}()


	repository := orderRepository.NewOrderRepository()
	service := orderService.NewOrderService(repository)
	api := orderAPI.NewAPI(service)


	orderServer, err := orderv1.NewServer(api)
	if err != nil {
		return err
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))
	r.Mount("/", orderServer)

	srv := &http.Server{
		Addr:              net.JoinHostPort("localhost", httpPort),
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	go func() {
		log.Printf("order http server on %s", httpPort)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("http serve error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("http shutdown error: %v", err)
	}

	return nil
}
