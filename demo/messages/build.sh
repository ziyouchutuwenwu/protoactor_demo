#protoc -I=. -I=$GOPATH/src --gogoslick_out=plugins=grpc:. protos.proto
protoc -I=. -I=$HOME/projects/golang/protoactor_demo/vendor --gogoslick_out=plugins=grpc:. protos.proto
go build