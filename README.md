# webdav

WebDAV 传输，支持浏览器打开

## 使用说明

默认监听 8088 端口，浏览器打开需使用 `/web` 路径

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

- [ ] 美化 autoindex 的样式
- [x] 记录日志

## 开源地址

https://github.com/nibazshab/webdav

## 许可证

MIT © ZShab Niba
