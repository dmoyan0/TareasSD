# Use the specified base image
ARG BASE_IMAGE=golang:1.21.3
FROM ${BASE_IMAGE} AS builder

ARG EARTH_PORT= 50051

ARG SERVER_TYPE

# Set the working directory inside the container
WORKDIR /app

# Copy the parent directory's go.mod and go.sum files to the container
COPY go.mod .

# Download and install Go dependencies
RUN go mod download

# Copy the rest of your application code to the container
COPY . .

CMD if [ "$SERVER_TYPE" = "tierra" ]; then \
        cd /app; \
        go build -o tierra; \
        ./tierra-server; \
    else \
        echo "Invalid SERVER_TYPE argument."; \
    fi