#
# Hackery so that import "capsule8/api/v0/file.proto" works
#
PROTO_ROOT=../../../

PROTOC_GEN_GO=${GOPATH}/bin/protoc-gen-go
PROTOC_GEN_GRPC_GATEWAY=${GOPATH}/bin/protoc-gen-grpc-gateway

GRPC_OUT=$(PROTO_ROOT)../
GRPC_GATEWAY_OUT=$(PROTO_ROOT)

IMPORT_PATH=github.com/capsule8/api/v0

INC=-I$(PROTO_ROOT) -I../third_party/googleapis -I../third_party/protobuf/src

.PHONY: api
api: $(PROTO_ROOT)/capsule8/api/v0/*.proto
	mkdir -p $(GRPC_OUT) $(GRPC_GATEWAY_OUT)

        # Compile grpc and gateway stubs
	protoc --plugin=protoc-gen-go=$(PROTOC_GEN_GO) \
		--go_out=plugins=grpc:$(GRPC_OUT) \
		--plugin=protoc-gen-grpc-gateway=$(PROTOC_GEN_GRPC_GATEWAY) \
		--grpc-gateway_out=logtostderr=true:$(GRPC_GATEWAY_OUT) \
		--descriptor_set_out=./descriptor.protoset \
		$(INC) \
		$?
