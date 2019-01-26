# vessel-service/Dockerfile
FROM golang:1.11.4 as builder

WORKDIR /go/src/github.com/grelol/shipper/vessel-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 go build -a -installsuffix cgo .


FROM debian:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/Egrelol/shipper/vessel-service/vessel-service .

CMD ["./vessel-service"]