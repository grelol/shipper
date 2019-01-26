FROM golang:1.11.4 as builder

WORKDIR /go/src/github.com/grelol/shipper/consignment-cli

COPY , ,

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure

RUN CGO_ENABLED = 0 go build -a -installsuffix cgo .

FROM debian:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

ADD consignment.json /app/consignment.json
COPY --from=builder /go/src/github.com/grelol/shipper/consignment-cli/consignment-cli .

CMD ["./consignment-cli"]