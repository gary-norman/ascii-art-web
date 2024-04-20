# syntax=docker/dockerfile:1

# Use the official Golang image as base
FROM golang:1.22.0

# Define metadata labels
LABEL version="1.0" \
       name="Ascii Art Web Project"
      description="Application for generating and processing ascii art text" \
      maintainers="Gary Norman & Kamil Ornal"

# Add current directory to the Go workspace
ADD . /go/src/myapp

# Set working directory
WORKDIR /go/src/myapp

# Download Go modules
COPY go.mod go.sum ./
#RUN go mod download
RUN #go get myapp
RUN go install

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /myapp

# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can (optionally) document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD "/myapp"