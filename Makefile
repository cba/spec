PROJECTS=auth cms file helloworld order pay product user
PROTO_DIR=./proto
DOC_DIR=./docs

all: proto protodoc-html

.PHONY:proto
proto:
	$(shell for var in $(PROJECTS); do mkdir -p $${var}; protoc --go_out=$${var} --go_opt=paths=source_relative --go-grpc_out=$${var} --go-grpc_opt=paths=source_relative --proto_path=$(PROTO_DIR) $${var}.proto; done)

.PHONY:auth
auth:
	@mkdir -p auth
	@protoc --go_out=auth --go_opt=paths=source_relative --go-grpc_out=auth --go-grpc_opt=paths=source_relative --proto_path=$(PROTO_DIR) auth.proto

product:
	@mkdir -p product
	@protoc --go_out=product --go_opt=paths=source_relative --go-grpc_out=product --go-grpc_opt=paths=source_relative --proto_path=$(PROTO_DIR) product.proto

pay:
	@mkdir -p pay
	@protoc --go_out=pay --go_opt=paths=source_relative --go-grpc_out=pay --go-grpc_opt=paths=source_relative --proto_path=$(PROTO_DIR) pay.proto

helloworld:
	@mkdir -p helloworld
	@protoc --go_out=helloworld --go_opt=paths=source_relative --go-grpc_out=helloworld --go-grpc_opt=paths=source_relative --proto_path=$(PROTO_DIR) helloworld.proto


#生成proto文件的markdown接口定义文档
protodoc-markdown:
	@go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
	@mkdir -p $(DOC_DIR)
	@protoc --doc_out=$(DOC_DIR) --proto_path=$(PROTO_DIR) --doc_opt=markdown,$(DOC_DIR)/proto.md $(PROTO_DIR)/*.proto

#生成proto文件的html接口定义文档
protodoc-html:
	@go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
	@mkdir -p $(DOC_DIR)
	@protoc --doc_out=$(DOC_DIR) --proto_path=$(PROTO_DIR) --doc_opt=html,$(DOC_DIR)/index.html $(PROTO_DIR)/*.proto