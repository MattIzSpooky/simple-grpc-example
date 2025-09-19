package main

import (
	"context"
	"database/sql"
	"github.com/mattizspooky/simple-grpc-example/v2/internal/db"
	"github.com/mattizspooky/simple-grpc-example/v2/internal/pb"
	"github.com/mattizspooky/simple-grpc-example/v2/internal/service"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/notes?sslmode=disable"
	}

	dbConn, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// HTTP server
	r := chi.NewRouter()
	svc := service.NewService(queries)
	svc.RegisterHTTP(r)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// gRPC server
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	grpcSvc := service.NewGRPCServer(queries)
	pb.RegisterNotesServiceServer(grpcServer, grpcSvc)

	// run servers
	go func() {
		log.Println("HTTP listening on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http: %v", err)
		}
	}()

	go func() {
		log.Println("gRPC listening on :9090")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("grpc: %v", err)
		}
	}()

	// wait for signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	// graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	grpcServer.GracefulStop()
}
