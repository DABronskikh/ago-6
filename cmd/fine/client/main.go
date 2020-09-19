package main

import (
	"context"
	fineV1Pb "github.com/DABronskikh/ago-6/pkg/fine/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"time"
)

const defaultPort = "9999"
const defaultHost = "0.0.0.0"

func main() {
	port, ok := os.LookupEnv("APP_PORT")
	if !ok {
		port = defaultPort
	}

	host, ok := os.LookupEnv("APP_HOST")
	if !ok {
		host = defaultHost
	}

	if err := execute(net.JoinHostPort(host, port)); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func execute(addr string) (err error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			if err == nil {
				err = cerr
				return
			}
			log.Print(err)
		}
	}()

	client := fineV1Pb.NewAutopayServiceClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	responseFindByAll, err := client.FindByAll(ctx, &fineV1Pb.AutopaysRequest{UserId: 2})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			log.Print(st.Code())
			log.Print(st.Message())
		}
		return err
	}
	log.Print("FindByAll:", responseFindByAll)

	responseFindById, err := client.FindById(ctx, &fineV1Pb.AutopayRequestById{AutopayId: 1})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			log.Print(st.Code())
			log.Print(st.Message())
		}
		return err
	}
	log.Print("FindById:", responseFindById)

	responseCreate, err := client.Create(ctx, &fineV1Pb.Autopay{
		Name:  "demo-3",
		Phone: "9001001020",
	})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			log.Print(st.Code())
			log.Print(st.Message())
		}
		return err
	}
	log.Print("Create:", responseCreate)

	responseUpdate, err := client.Update(ctx, &fineV1Pb.Autopay{
		Id:    1,
		Name:  "new name",
		Phone: "1111111111",
	})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			log.Print(st.Code())
			log.Print(st.Message())
		}
		return err
	}
	log.Print("Update:", responseUpdate)

	responseDelete, err := client.Delete(ctx, &fineV1Pb.Autopay{Id: 2})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			log.Print(st.Code())
			log.Print(st.Message())
		}
		return err
	}
	log.Print("Delete:", responseDelete)

	return nil
}
