#!/bin/bash

python3 -m grpc_tools.protoc --proto_path=./proto --python_out=./front/proto --grpc_python_out=./front/proto proto/article/article.proto

python3 -m grpc_tools.protoc --proto_path=./proto --python_out=./front/proto --grpc_python_out=./front/proto proto/user/user.proto

if [[ $(uname -s) == Darwin* ]];
then
    sed -i '' 's/from article import/from . import/1' front/proto/article/article_pb2_grpc.py
    sed -i '' 's/from user import/from . import/1' front/proto/user/user_pb2_grpc.py
else
    sed -i 's/from article import/from . import/1' front/proto/article/article_pb2_grpc.py
    sed -i 's/from user import/from . import/1' front/proto/user/user_pb2_grpc.py
fi
