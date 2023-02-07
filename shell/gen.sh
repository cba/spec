#!/bin/bash

PROJECTS=(base auth cms file order pay product sms user)
PROTO_DIR=proto
DOC_DIR=docs

mkdir -p docs/openapiv2;
mkdir -p php;

protoc --doc_out=$DOC_DIR --proto_path=$PROTO_DIR --doc_opt=html,$DOC_DIR/index.html $PROTO_DIR/*.proto

# shellcheck disable=SC2068
for var in ${PROJECTS[@]}; do
    mkdir -p go/$var;

    # 生成pb
    protoc --proto_path=$PROTO_DIR -I . \
        --go_out go/$var \
        --go_opt paths=source_relative \
        --go-grpc_out go/$var \
        --go-grpc_opt paths=source_relative \
        $var.proto;

    # 生成 grpc-gateway
    protoc --proto_path $PROTO_DIR -I . \
     		--grpc-gateway_out go/$var \
    		--grpc-gateway_opt logtostderr=true \
    		--grpc-gateway_opt paths=source_relative \
    		--grpc-gateway_opt generate_unbound_methods=true \
    		$var.proto

    # 生成 OpenApi V2
    protoc --proto_path $PROTO_DIR -I . \
 		    --openapiv2_out ./docs/openapiv2 \
        --openapiv2_opt logtostderr=true \
        $var.proto

    # 生成php
    protoc --proto_path=$PROTO_DIR --php_out=php $var.proto
done;