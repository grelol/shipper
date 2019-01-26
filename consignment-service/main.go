package main

import (
	"context"
	"log"

	pb "github.com/grelol/shipper/consignment-service/proto/consignment"
	micro "github.com/micro/go-micro"
)

const (
	port = ":50051"
)

// Repository is an interface
type Repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// ConsignmentRepository is a temporary datastore
type ConsignmentRepository struct {
	consignments []*pb.Consignment
}

// Create creates consignments
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

// GetAll returns all consignments
func (repo *ConsignmentRepository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo ConsignmentRepository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	log.Printf("Created consignment: %v\n", consignment)
	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	log.Println("Starting Consignment Service")

	repo := &ConsignmentRepository{}

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterConsignmentServiceHandler(srv.Server(), &service{*repo})

	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
