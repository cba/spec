PROJECTS=base auth cms file order pay product sms user
PROTO_DIR=./proto
DOC_DIR=./docs

all: build-go build-php protodoc-html

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

.PHONY:build-php
build-php:
	$(shell for var in $(PROJECTS); do mkdir -p $${var}; protoc --proto_path=$(PROTO_DIR) --proto_path=proto --php_out=php $${var}.proto; done)

.PHONY:build-go
build-go:
	$(shell for var in $(PROJECTS); do mkdir -p ./go/$${var}; protoc --go_out=./go/$${var} --go_opt=paths=source_relative --go-grpc_out=./go/$${var} --go-grpc_opt=paths=source_relative --proto_path=$(PROTO_DIR) $${var}.proto; done)

.PHONY:build-dart
build-dart:
	$(shell for var in $(PROJECTS); do mkdir -p ./dart/$${var}; protoc --proto_path=$(PROTO_DIR) --go_out=./dart/$${var} --dart_out=./dart/$${var} -Iprotos $${var}.proto; done)
