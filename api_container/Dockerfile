# Build: docker build -t libary-apiserver . 
#         && docker images
# Run: docker run -p 3000:3000 libary-apiserver // docker run -p <host-port>:<container-port> <image-name>
# Test: check OpenAPI Definition

# Use golang-based image for container; golang version 1.12.4
FROM golang:1.12.4-alpine

# Enable go modules for downloading dependencies
ENV GO111MODULE=on
ENV GOPATH=/go/delivery

# Add git executable in container
RUN apk add --no-cache git

# Copy apiserver go file in local computer to container
RUN mkdir -p /go/src/apiserver
COPY . /go/src/apiserver
COPY ./main.go /go/src/app/main.go


# Set working directory in container
WORKDIR /go/delivery
COPY go.mod /go/delivery
COPY go.sum /go/delivery

# Build the application
# RUN go mod init github.com/derchrischkya/libary
#RUN go build -o server /go/src/app/main.go

# Expose listening port for application
EXPOSE 3000

# Run the application
ENTRYPOINT ["/go/bin/main"]