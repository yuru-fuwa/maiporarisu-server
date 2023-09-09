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
COPY ./app/go.mod .
COPY ./app/go.sum .
RUN go mod download

# Copy the code into the container
COPY ./app/main.go .

EXPOSE 8080

# Set command for the container
CMD ["go", "run", "./main.go"]