FROM golang:alpine3.20

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod .
COPY go.sum .

RUN mkdir .logs

# Download dependencies
RUN go mod download

# Copy the entire application to the container
COPY . .

# Build the Go application
RUN  go build -o main cmd/grpc/main.go

# Expose port 9000 for the Golang application
EXPOSE 9000

# Command to run the executable
CMD ["./main"]