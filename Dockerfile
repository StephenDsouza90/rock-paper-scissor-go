# Base image
FROM golang:alpine AS Builder

# Create and add files into /app directory
RUN mkdir /app
WORKDIR /app
ADD . /app

# Compile binary
RUN go build -o app .

# New base image
FROM alpine:latest

# Add artifacts from previous image to new image
COPY --from=Builder /app .

# Expose application port
EXPOSE 8080

# Command to run app
CMD ["/app"]