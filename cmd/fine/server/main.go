package main

import (
	"context"
	"github.com/DABronskikh/ago-6/cmd/fine/server/app"
	fineV1Pb "github.com/DABronskikh/ago-6/pkg/fine/v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const (
	defaultPort = "9999"
	defaultHost = "0.0.0.0"
	defaultDSN  = "postgres://app:pass@localhost:5432/db"
)

func main() {
	port, ok := os.LookupEnv("APP_PORT")
	if !ok {
		port = defaultPort
	}

	host, ok := os.LookupEnv("APP_HOST")
	if !ok {
		host = defaultHost
	}

	log.Println(host)
	log.Println(port)

	dsn, ok := os.LookupEnv("APP_DSN")
	if !ok {
		dsn = defaultDSN
	}

	if err := execute(net.JoinHostPort(host, port), dsn); err != nil {
		os.Exit(1)
	}
}

func execute(addr string, dsn string) (err error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Print(err)
		return err
	}

	businessSvc := app.NewService(pool)

	grpcServer := grpc.NewServer()
	server := app.NewServer(businessSvc)

	fineV1Pb.RegisterAutopayServiceServer(grpcServer, server)

	return grpcServer.Serve(listener)
}
