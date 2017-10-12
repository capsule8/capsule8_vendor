PROTO_ROOT=../../../../

PROTOC_GEN_GO=${GOPATH}/bin/protoc-gen-go

.PHONY: config
config: $(PROTO_ROOT)/capsule8/api/pkg/config/*.proto

        # Compile grpc and gateway stubs
	protoc --plugin=protoc-gen-go=$(PROTOC_GEN_GO) \
		--go_out=plugins=grpc:$(PROTO_ROOT)../ \
		--descriptor_set_out=./descriptor.protoset \
		-I$(PROTO_ROOT) \
		$?
