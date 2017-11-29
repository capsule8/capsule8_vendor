# Capsule8 APIs and Common Configuration Definitions

This directory defines component-level APIs and common configuration
formats for the Capsule8 platform. These definitions also serve as the
ground truth of the shared vocabulary used by Capsule8.

These definitions are specified using the
[protobuf](https://github.com/google/protobuf) syntax and follow
Google's [API Design Guide](https://cloud.google.com/apis/design/).
For useful examples, see the public
[Google APIs](https://github.com/googleapis/googleapis).

### Compiling protos

You will need the following dependencies:
[protoc](https://github.com/golang/protobuf)
[grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)

- `git submodule init`
- `git submodule update`
- `make`


### Compiling docs

Just run the following:

- `make docs`
