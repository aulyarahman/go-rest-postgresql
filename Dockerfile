FROM golang:latest as builder

COPY go.mod go.sum /go/src/github.com/aulyarahman/bucketeer/

WORKDIR /go/src/github.com/aulyarahman/bucketeer

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/bucketeer github.com/aulyarahman/bucketeer

FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=builder /go/src/github.com/aulyarahman/bucketeer/build/bucketeer /usr/bin/bucketeer

EXPOSE 8080 8080

ENTRYPOINT ["/usr/bin/bucketeer"]
