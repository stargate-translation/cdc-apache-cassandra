FROM ubuntu:20.04

# Copy your Go application source code into the container
COPY ./proxy.go .

RUN ["mkdir", "/golang"]
ENV GOROOT /golang/.go
ENV GOPATH /golang/go

USER root
# Install any required dependencies for your Go server
RUN apt update && apt upgrade && apt install curl docker.io -y

RUN curl -sSL https://git.io/g-install | sh -s -- bash -y

RUN ["chmod", "+x", "/golang"]

# Build the Go application
RUN ["/golang/go/bin/go", "mod", "init", "proxy"]
RUN ["/golang/go/bin/go", "build", "-o", "proxy"]

# Start your Go server
CMD ["nohup", "./proxy"]
