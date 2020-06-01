curl -L -o protoc3.zip https://github.com/google/protobuf/releases/download/v3.12.2/protoc-3.12.2-linux-x86_64.zip
unzip protoc3.zip  -d protoc3
mv protoc3/bin/* /usr/local/bin/
mv protoc3/include/* /usr/local/include/

go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/golang/protobuf/proto

mkdir -p ${GOPATH}/src/github.com/nageshborate/grpc-go-course
cd ${GOPATH}/src/github.com/nageshborate/grpc-go-course
ls

tail -f /dev/null