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
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	orderAPI "order/internal/api/order/v1"
	orderInventoryClientV1 "order/internal/client/grpc/inventory/v1"
	orderPaymentClientV1 "order/internal/client/grpc/payment/v1"
	"order/internal/migrator"
	orderRepository "order/internal/repository/order"
	orderService "order/internal/service/order"
	orderv1 "shared/pkg/openapi/order/v1"
	inventory_v1 "shared/pkg/proto/inventory/v1"
	payment_v1 "shared/pkg/proto/payment/v1"
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

	inventoryServiceClient := inventory_v1.NewInventoryServiceClient(invConn)
	paymentServiceClient := payment_v1.NewPaymentServiceClient(payConn)

	invClient := orderInventoryClientV1.NewInventoryClient(inventoryServiceClient)
	payClient := orderPaymentClientV1.NewPaymentClient(paymentServiceClient)

	err = godotenv.Load(".env")
	if err!=nil{
		log.Printf("failed to load .env file: %v\n", err)
	}

	ctx := context.Background()
	var cancel context.CancelFunc

	dbURI := os.Getenv("POSTGRES_URI")
	if dbURI == "" {
		log.Println("POSTGRES_URI not set in environment")
		return errors.New("POSTGRES_URI not set")
	}

	con, err := pgx.Connect(ctx, dbURI)
	if err != nil {
		log.Printf("failed to connect to database: %v\n", err)
		return err
	}
	defer func() {
		if cerr := con.Close(ctx); cerr != nil {
			log.Printf("database connection close error: %v", cerr)
		}
	}()

	err = con.Ping(ctx)
	if err != nil {
		log.Printf("failed to ping database: %v\n", err)
		return err
	}
	sqlDB := stdlib.OpenDB(*con.Config().Copy())
	migrationsDIR := os.Getenv("MIGRATIONS_DIR")
	if migrationsDIR == "" {
		log.Println("MIGRATIONS_DIR not set in environment")
		return errors.New("MIGRATIONS_DIR not set")
	}
	migratorRunner := migrator.NewMigrator(sqlDB, migrationsDIR)

	err = migratorRunner.Up()
	if err != nil {
		log.Printf("migration failed: %v\n", err)
		return err
	}

	repository := orderRepository.NewOrderRepository(sqlDB)
	service := orderService.NewOrderService(repository, invClient, payClient)
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

	ctx, cancel = context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("http shutdown error: %v", err)
	}

	return nil
}
