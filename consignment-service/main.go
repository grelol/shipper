package main

import (
	"log"
	"os"

	pb "github.com/grelol/shipper/consignment-service/proto/consignment"
	vesselProto "github.com/grelol/shipper/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	if err != nil {
		log.Panic("Failed to connect to the datastore with host %s - %v", host, err)
	}
	defer session.Close()

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	srv.Init()

	pb.RegisterConsignmentServiceHandler(srv.Server(), &service{session, vesselClient})

	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
