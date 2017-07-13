PROTOC_GEN_GO=$$GOPATH/bin/protoc-gen-go
PROTOC_GEN_GRPC_GATEWAY=$$GOPATH/bin/protoc-gen-grpc-gateway
GRPC_OUT=.
GRPC_GATEWAY_OUT=$(GRPC_OUT)
VENDOR=../vendor

INC=-I. -I../third_party/protobuf/src/ \
	-I../third_party/googleapis/

# Directories containing one or more .proto files
PROTODIRS=$(shell find . -name "*.proto" | sed 's/\(.*\)\/.*/\1/' | uniq)

.PHONY: api
api: *.proto
	mkdir -p $(GRPC_OUT) $(GRPC_GATEWAY_OUT)

        # Compile grpc and gateway stubs
	protoc --plugin=protoc-gen-go=$(PROTOC_GEN_GO) \
		--go_out=plugins=grpc:$(GRPC_OUT) \
		--plugin=protoc-gen-grpc-gateway=$(PROTOC_GEN_GRPC_GATEWAY) \
		--grpc-gateway_out=logtostderr=true:$(GRPC_GATEWAY_OUT) \
		$(INC) \
		$?
