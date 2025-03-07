# Use the official Golang image from the Docker Hub
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.* ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./main"]
