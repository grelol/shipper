build:
	protoc -I. --go_out=plugins=micro:$(HOME)/go/src/github.com/grelol/shipper/consignment-service/ \
		proto/consignment/consignment.proto
	CGO_ENABLED=0 go build -o consignment-service -a -installsuffix cgo main.go repository.go handler.go datastore.go
	sudo docker build -t consignment-service .

run:
	sudo docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns consignment-service