# 包管理工具

Go version 1.12

`go mod init mode_name init_mode_name` 项目目录下生成一个go.mod文件，包管理通过此文件<br>

不需要通过`go get`命令下载包到`$GOPATH/src`,直接`go run xxx.go`,go会自动查找代码中的包，下载依赖包，并把具体的依赖关系和版本写入到go.mod和go.sum文件中<br>
依赖包下载到`$GOPATH/pkg/mod`下<br>
go.mod未指定版本，下载最新版本