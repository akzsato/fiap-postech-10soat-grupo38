# Start from a Debian-based Golang official image
FROM golang:1.23.7-alpine AS builder

# Set the working directory inside the container  
WORKDIR /app  

# Copy the go mod and sum files  
COPY go.mod go.sum ./  

# Download all dependencies  
RUN go mod download  

# Copy the source code from your host to your image filesystem.  
COPY . . 

# Build the Go app  
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main .  

# Use a minimal alpine image for the final stage  
FROM alpine:latest  

# Set the working directory inside the container  
WORKDIR /root/  

# Copy environment variables definition
COPY --from=builder /app/.env .

# Copy the binary from the builder stage  
COPY --from=builder /app/main .  

# Run the binary  
CMD ["./main"]

# Expose port 8080
EXPOSE 8080