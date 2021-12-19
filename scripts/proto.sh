#!/bin/bash

protoc \
  --proto_path=api/protobuf "api/protobuf/questioner.proto" \
  "--go_out=internal/common/genproto/questioner" --go_opt=paths=source_relative \
  --go-grpc_opt=require_unimplemented_servers=false \
  "--go-grpc_out=internal/common/genproto/questioner" --go-grpc_opt=paths=source_relative