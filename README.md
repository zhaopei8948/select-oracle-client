# 先克隆项目到本地
```
git clone https://github.com/zhaopei8948/select-oracle-client
```
## 生成grpc code用下面命令
```
protoc -I . protos/select.proto --go_out=plugins=grpc:.
```

## 运行客户端，先运行服务端，服务端地址： https://github.com/zhaopei8948/select-oracle
```
go run main.go
```

