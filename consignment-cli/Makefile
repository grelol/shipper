build:
	CGO_ENABLED=0 go build
	sudo docker build -t consignment-cli .

run:
	sudo docker run -e MICRO_REGISTRY=mdns consignment-cli