#!/bin/bash

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. \
    --go-grpc_opt=require_unimplemented_servers=false,paths=source_relative \
    proto/article/article.proto

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. \
    --go-grpc_opt=require_unimplemented_servers=false,paths=source_relative \
    proto/user/user.proto

python3 -m grpc_tools.protoc --proto_path=./proto --python_out=./front/proto --grpc_python_out=./front/proto proto/article/article.proto

python3 -m grpc_tools.protoc --proto_path=./proto --python_out=./front/proto --grpc_python_out=./front/proto proto/user/user.proto
