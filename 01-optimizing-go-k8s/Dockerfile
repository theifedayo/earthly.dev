# Use the official Golang base image
FROM golang:1.16

# Set working directory
WORKDIR /app

# Copy and build the Go application
COPY . .
RUN go build -o myapp

# Run the application
CMD ["./myapp"]


#############
#OR
############
FROM golang:1.16-alpine

# Set working directory
WORKDIR /app

# Copy and build the Go application
COPY . .
RUN go build -o myapp

# Run the application
CMD ["./myapp"]