all:
	$(MAKE) -C v0
	$(MAKE) -C pkg/recorder/v0

.PHONY: docs
docs:
	docker run --rm \
		-v $(shell pwd)/docs:/out \
		-v $(shell pwd)/third_party/googleapis/google:/protos/google \
		-v $(shell pwd)/v0:/protos/capsule8/api/v0 \
		pseudomuto/protoc-gen-doc --doc_opt=markdown,API.md -I /protos /protos/capsule8/api/v0/telemetry_service.proto
	docker run --rm \
		-v $(shell pwd)/docs:/out \
		-v $(shell pwd)/third_party/googleapis/google:/protos/google \
		-v $(shell pwd)/v0:/protos/capsule8/api/v0 \
		pseudomuto/protoc-gen-doc --doc_opt=html,API.html -I /protos /protos/capsule8/api/v0/telemetry_service.proto
