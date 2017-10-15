# protoc -I=. -I=$GOPATH/src --gogoslick_out=plugins=grpc:. protos.proto
protoc -I=. -I=/Users/mmc/dev/golang/my_apps/src/my_app/vendor/ --gogoslick_out=plugins=grpc:. protos.proto
go build
