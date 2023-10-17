# Use the official Golang image to create a build artifact.
FROM golang:1.21 as builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /app
WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

RUN GOOS=linux go build -o main ./cmd/main.go

FROM public.ecr.aws/amazonlinux/amazonlinux:latest

COPY --from=builder /app/main /app/main

EXPOSE 8080

# Set command for the container
CMD ["/app/main"]