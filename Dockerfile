# Use the specified base image
ARG BASE_IMAGE=golang:1.21.3
FROM ${BASE_IMAGE} AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the parent directory's go.mod and go.sum files to the container
COPY go.mod .

# Download and install Go dependencies
RUN go mod download

# Copy the rest of your application code to the container
COPY . .

# Build the server
RUN go build -o tierra

# Define default value for SERVER_TYPE
ARG SERVER_TYPE=tierra

# Set up the command to run the server based on SERVER_TYPE
CMD if [ "$SERVER_TYPE" = "tierra" ]; then \
        ./tierra; \
    else \
        echo "Invalid SERVER_TYPE argument."; \
    fi
