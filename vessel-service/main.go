package main

import (
	"log"
	"os"

	pb "github.com/grelol/shipper/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
)

const (
	defaultHost = "localhost:27017"
)

func CreateDummy(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vess001", Name: "Salty Secrets", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}

func main() {

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	if err != nil {
		log.Fatalf("Failed to connect to the datastore: %v", err)
	}
	defer session.Close()

	repo := &VesselRepository{session.Copy()}

	CreateDummy(repo)

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{session})

	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to run vessel service: %v", err)
	}

}
