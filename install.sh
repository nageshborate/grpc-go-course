cd $GOPATH
curl -L -o protoc3.zip https://github.com/google/protobuf/releases/download/v3.12.2/protoc-3.12.2-linux-x86_64.zip
unzip protoc3.zip  -d protoc3
mv protoc3/bin/* /usr/local/bin/
mv protoc3/include/* /usr/local/include/

go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/golang/protobuf/proto
go get -u google.golang.org/grpc

tail -f /dev/null