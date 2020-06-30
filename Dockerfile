# First stage: build the application
FROM golang:1.14-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/netoax/slave-test

# Copy the source code from the current directory to $WORKDIR (inside the container)
COPY . .

# Download go modules to local cache
RUN go mod download

# Build the soimulator
RUN go build slave.go

ENTRYPOINT ["./slave"]
