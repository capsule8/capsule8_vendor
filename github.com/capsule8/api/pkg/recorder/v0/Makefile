PROTO_ROOT=../../../../../

PROTOC_GEN_GO=${GOPATH}/bin/protoc-gen-go

.PHONY: recoder
recorder: $(PROTO_ROOT)/capsule8/api/pkg/recorder/v0/*.proto

        # Compile grpc and gateway stubs
	protoc --plugin=protoc-gen-go=$(PROTOC_GEN_GO) \
		--go_out=plugins=grpc:$(PROTO_ROOT)../ \
		-I$(PROTO_ROOT) \
		$?
