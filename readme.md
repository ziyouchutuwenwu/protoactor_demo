# 添加一些须知

## 生成protobuf对应的go文件
### 库需要安装到GOPATH
```bash
sudo apt install golang-goprotobuf-dev
```

### 下载protoc
- 到`https://github.com/protocolbuffers/protobuf` 下载 

### 安装protoc-gen-gogoslick
- 到`https://github.com/gogo/protobuf` 的 `protoc-gen-gogoslick` 目录
```bash
export export GO111MODULE=off; go build
```

## vendor目录说明
- vendor目录仅用于生成proto文件生成go源码用
- 先注掉message相关的调用，然后go mod vendor，然后message目录下build.sh