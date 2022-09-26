## 环境变量
安装完毕后，命令行输入 `go env` 查看go环境变量
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
- `go get` ：获取远程包
- `go run` : 直接运行程序
- `go build` ：测试编译，检查是否有编译错误
- `go fmt` ：格式化源码(部分IDE在保存时会自动调用)
- `go install` ：编译包文件并编译整个程序
- `go test` ：运行测试文件
- `go doc` ：查看文档

## 基础
GO程序的一般结构：basic_structure.go

- Go程序是通过 `package` 来组织的（与Python相似）
- 只有 `package` 名称为 `main` s的包可以包含 main 函数
- 一个可执行程序 **有且仅有** 一个 main 包  
   
- 通过 `import` 关键字来导入其他非 main 包
- 通过 `const` 关键字来进行常量的定义
- 通过在函数体外使用 `var` 关键字来进行全局变量的声明与赋值
- 通过 `func` 关键字来进行函数的声明

### 可见性规则
- GO 语言中，使用 `大小写` 来决定该 变量、常量、类型、接口、结构 或 函数 是否可以被外部包所调用
- 函数名首字母 `小写` 即为 `Private` ，函数名首字母 `大写` 即为 `public`
