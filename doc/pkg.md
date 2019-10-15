# 包管理工具

go modules
-----------
Go version 1.12

`go mod init mode_name` 项目目录下生成一个go.mod文件，通过此文件进行包管理<br>
`go mod download` 下载git.mod声明依赖<br>

不需要通过`go get`命令下载包到`$GOPATH/src`,直接`go run xxx.go`,go会自动查找代码中的包，下载依赖包，并把具体的依赖关系和版本写入到go.mod和go.sum文件中<br>
依赖包下载到`$GOPATH/pkg/mod`下<br>
go.mod未指定版本，下载最新版本<br>

#####依赖包地址变更

方式一：<br>
replace替换包位置(要指明version):<br>
```
replace (
    golang.org/x/text => github.com/golang/text latest
)
```

方式二：下载仓库，编辑go.mod<br>
`go get github.com/golang/image@master` 获取最新提交的版本<br>

`go mod edit -replace=golang.org/x/image/font=github.com/golang/image/font@v0.0.0-20190622003408-7e034cad6442` 命令编辑<br>
方式三：传统方式<br>


通过go.mod中module来引入自身package<br>


vgo
----------

Goland配置代理：`https://goproxy.io`（安装vgo:`go get -u golang.org/x/vgo`）<br>

gopm
-----
`go get `


govendor
-------
`go get -u github.com/kardianos/govendor`

problem:`Error: Package "xxx" not a go package or not in GOPATH.`

godep 
----
解决包依赖问题，依赖vendor<br>
`go get github.com/tools/godep`