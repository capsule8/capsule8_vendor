# Dockerfile for build container used by buildkite
FROM golang:1.8.4

# Install unzip
RUN apt-get update && apt-get install -y unzip

# Install protoc version 3.3.0
RUN wget https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip
RUN unzip protoc-3.3.0-linux-x86_64.zip

# Install protoc-gen-go
RUN go get -u github.com/golang/protobuf/protoc-gen-go

# Install protoc-gen-doc
RUN go get -u github.com/capsule8/protoc-gen-doc/cmd/protoc-gen-doc

LABEL com.capsule8.name="builder" \
    com.capsule8.url="https://www.capsule8.com" \
    com.capsule8.vcs-url="https://github.com/capsule8/api" \
    com.capsule8.version="0.0.1" \
    com.capsule8.schema-version="0.0"

ENTRYPOINT ["/go/bin/protoc"]
