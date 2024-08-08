# Use the official Golang image as the base image
FROM golang:1.22.2-alpine

# Install bash
RUN apk add --no-cache bash

LABEL maintainer="bernadokumu1@gmail.com" \
    version="1.0" \ 
    environment="dev"

# Set the working directory inside the container
WORKDIR /ascii

# Copy the source code to the workspace
COPY . .

# Build the Go app
RUN go build -o ascii-art-web-dockerize .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD [ "./ascii-art-web-dockerize" ]