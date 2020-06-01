wget https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip
unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip -d protoc
mv protoc/bin/protoc /usr/bin/protoc

mkdir -p ${GOPATH}/src/github.com/golang
cd ${GOPATH}/src/github.com/golang
git clone https://github.com/golang/protobuf
cd protobuf
go get ./protoc-gen-go

mkdir -p ${GOPATH}/src/github.com/nageshborate/grpc-go-course
cd ${GOPATH}/src/github.com/nageshborate/grpc-go-course
ls
