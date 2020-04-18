# 说明

## 安装 protobuf 需要的库

```bash
sudo apt install golang-goprotobuf-dev
```

## 安装 protoc 和 protoc-gen-gogoslick 命令行工具

### 下载 protoc

- 到`https://github.com/protocolbuffers/protobuf` 下载

### 安装 protoc-gen-gogoslick

- 到`https://github.com/gogo/protobuf` 的 `protoc-gen-gogoslick` 目录

```bash
export GO111MODULE=off; go build
```

## vendor 目录说明

- vendor 目录仅用于生成 proto 文件生成 go 源码用
- 先注掉 message 相关的调用，然后 go mod vendor，然后 message 目录下 build.sh
