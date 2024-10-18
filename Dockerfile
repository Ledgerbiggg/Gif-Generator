# Use the official Golang images as the builder
FROM golang AS builder
# Maintainer information
MAINTAINER LEDGERBIGGG

# Set the working directory inside the Docker container
WORKDIR /go/src

# Set the Go proxy for module downloading
ENV GOPROXY https://goproxy.cn

# Copy the local project files to the working directory
ADD . /go/src

# Compile the project
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo main.go

# Use the Alpine images for the final, smaller production images
FROM alpine AS prod

# Copy the compiled binary from the builder images
COPY --from=builder /go/src/main /main
# Copy the configuration file from the builder images
COPY --from=builder /go/src/config.yaml ./config.yaml

# Set the command to execute the binary
CMD ["/main"]
