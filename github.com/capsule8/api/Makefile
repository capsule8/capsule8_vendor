PWD=$(shell pwd)

all:
	$(MAKE) -C v0
	$(MAKE) -C pkg/recorder/v0

builder:
	docker build . -t api_builder

.PHONY: docs
docs: builder
	docker run --rm \
		-v $(PWD)/docs:/out \
		-v $(PWD)/third_party/googleapis/google:/protos/google \
		-v $(PWD)/v0:/protos/capsule8/api/v0 \
		api_builder --doc_out=/out --doc_opt=html,API.html:google/* -I /protos /protos/capsule8/api/v0/telemetry_service.proto
	docker run --rm \
		-v $(PWD)/docs:/out \
		-v $(PWD)/third_party/googleapis/google:/protos/google \
		-v $(PWD)/v0:/protos/capsule8/api/v0 \
		api_builder --doc_out=/out --doc_opt=markdown,API.md:google/* -I /protos /protos/capsule8/api/v0/telemetry_service.proto
