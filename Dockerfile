# Use Go base image
FROM golang:1.22.3

# Set working directory inside container
WORKDIR /app

# Copy Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o food

# Expose the required port
EXPOSE 8081

# Command to run the application
CMD ["./food"]
