# gRPC

## Protobuffer
1. 安装生成protobuf文件的工具
```bash
# protoc-gen-go
RUN go get -u -v -ldflags '-w -s' google.golang.org/protobuf/cmd/protoc-gen-go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# protoc-gen-go-grpc
RUN go get -u -v -ldflags '-w -s' google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# MacOSX下需安装
brew install protobuf
```

2. 生成pb和grpc.pb文件
```bash
protoc -I=./ --go_out=./ --go-grpc_out=require_unimplemented_servers=false:./ user.proto
```



