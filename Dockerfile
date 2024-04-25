# syntax=docker/dockerfile:1

# Use the official Golang image as base
FROM golang:1.22.0-alpine

# Define metadata labels
LABEL version="1.0" \
      name="Ascii Art Web Project" \
      description="Application for generating and processing ascii art text" \
      maintainers="Gary Norman & Kamil Ornal"

# Add current directory to the Go workspace
ADD . /go/src/app

# Set working directory
WORKDIR /go/src/app

# Download and install Go dependencies
COPY go.mod go.sum ./
#RUN go get ./
#RUN go install
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Build the application
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /myapp
RUN go build  -o /ascii-art-web
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can (optionally) document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD [ "/ascii-art-web" ]

# Builder stage
#FROM golang:alpine as builder
#
#ENV GO111MODULE=on \
#    CGO_ENABLED=0 \
#    GOOS=linux \
#    PORT=8080 \
#    GOARCH=amd64
#
#EXPOSE 8080
#
#WORKDIR /build
#COPY go.mod .
#COPY go.sum .
#RUN go mod download
#COPY . .
#RUN go build -o asciiartweb ./
#
## Final stage
#FROM alpine:latest
#LABEL maintainers="Kamil Ornal <maintainer.one@example.com>, \
#                   Gary Norman <maintainer.three@example.com>"
#COPY --from=builder /build/asciiartweb /app/asciiartweb
#WORKDIR /app
#CMD ["./asciiartweb"]