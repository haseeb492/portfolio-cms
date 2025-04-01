# Use an official lightweight Go image
FROM golang:1.24-alpine
# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o myapp .

# Expose port 8080 for the application
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]
