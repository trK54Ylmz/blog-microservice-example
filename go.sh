#!/bin/bash

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. \
    --go-grpc_opt=require_unimplemented_servers=false,paths=source_relative \
    proto/article/article.proto

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. \
    --go-grpc_opt=require_unimplemented_servers=false,paths=source_relative \
    proto/user/user.proto
