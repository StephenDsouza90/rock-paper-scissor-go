# Base image
FROM golang:1.15.0-alpine

# Create directory
RUN mkdir /app

# Add files from root directory into /app
ADD . /app

# Set /app as working directory
WORKDIR /app

# Compile the binary
RUN go build -o main .

# Expose application port
EXPOSE 8080

# Command to run app
CMD ["/app/main"]