.PHONY: pre-push generate-proto
pre-push: generate-proto go-generate-after

.PHONY: protoc-upgrade protoc-update generate
protoc-upgrade: protoc-update generate-proto go-generate-proto

protoc-update:
	go get -u google.golang.org/protobuf@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

generate-proto:
	@echo "Compiling protobuf files..."
	protoc \
		-I=. \
		--go_out=. \
		--go_opt=module="blog" \
		--go-grpc_out=. \
		--go-grpc_opt=module="blog" \
		--proto_path=models/grpc models/grpc/*.proto
	@echo "Compiling completed."

.PHONY: go-generate-proto generate-proto go-generate-after
go-generate-proto: generate-proto go-generate-after

go-generate-after:
	$ go install github.com/favadi/protoc-go-inject-tag@latest && \
	protoc-go-inject-tag -input=models/grpc/proto/*.pb.go
