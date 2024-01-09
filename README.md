# webdav

WebDAV 传输

残次品，仅支持从网页打开

## 使用说明

默认监听 8088 端口

1. 编译 main.go
2. 附带目录参数运行程序 `./webdav .`

__编译步骤__

```sh
git clone https://github.com/nibazshab/webdav.git
cd webdav
go get -d -v ./...
CGO_ENABLED=0 go build -ldflags="-s -w"
```

## 计划

- [ ] 更改前端样式
- [ ] 记录日志
- [ ] 支持连接工具

## 开源地址

https://github.com/nibazshab/webdav

## 许可证

MIT © ZShab Niba
