## 环境变量
安装完毕后，命令行输入` go env `查看go环境变量
```
set GOCACHE=C:\Users\40483\AppData\Local\go-build
set GOENV=C:\Users\40483\AppData\Roaming\go\env
set GOEXE=.exe

set GOMODCACHE=C:\Users\40483\go\pkg\mod

## 工作目录、存放sdk以外的第三方类库
set GOPATH=C:\Users\40483\go

 ## 安装目录
set GOROOT=H:\Program Files\Go

## 工具目录
set GOTOOLDIR=H:\Program Files\Go\pkg\tool\windows_amd64

## 版本
set GOVERSION=go1.19.1
```

GOPATH下需要建立3个目录：
- bin - 存放编译后生成的可执行文件
- pkg - 存放编译后生成的包文件
- src - 存放项目源码 

## 常用命令
- go get ：获取远程包
- go run : 直接运行程序
- go build ：测试编译，检查是否有编译错误
- go fmt ：格式化源码(部分IDE在保存时会自动调用)
- go install ：编译包文件并编译整个程序
- go test ：运行测试文件
- go doc ：查看文档

