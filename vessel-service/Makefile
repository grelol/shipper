build:
	protoc -I. --go_out=plugins=micro:$(HOME)/go/src/github.com/grelol/shipper/vessel-service/ \
		proto/vessel/vessel.proto
	CGO_ENABLED=0 go build -o vessel-service -a -installsuffix cgo main.go repository.go handler.go datastore.go
	sudo docker build -t vessel-service .

run:
	sudo docker run -p 50052:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns vessel-service