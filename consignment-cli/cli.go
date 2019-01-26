package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	pb "github.com/grelol/shipper/consignment-service/proto/consignment"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
)

const (
	filename = "consignment.json"
)

func parseFile(filename string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, nil
}

func main() {

	cmd.Init()

	client := pb.NewConsignmentServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	consignment, err := parseFile(filename)
	if err != nil {
		log.Fatalf("Failed to parse file: %v", err)
	}

	resp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Failed to create: %v", err)
	}
	log.Printf("Created: %t\n", resp.Created)

	consignments, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Failed to get consignments: %v", err)
	}
	for _, v := range consignments.Consignments {
		log.Println(v)
	}
}
