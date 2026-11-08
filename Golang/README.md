
# Golang 基础

## Go mod 依赖管理

go mod init 初始化一个 Go 模块，创建 go.mod 模块
go mod tidy 分析项目实际使用的依赖，自动补充缺失依赖并清理无用依赖

require 当前 module 依赖的包
exclude 排除第三方包
replace 修改依赖包的路径或版本
retract 

go mod verify  验证依赖是否正确
go mod why 返回对指定模块依赖关系的最短路径

## Golang 编译运行

go build
go build .

编译指定文件  go build ./main.go
编译 main 文件，将会生成可执行文件

Go 中，一个文件夹就是一个包


go run ./main.go   运行


### 交叉编译

需要修改 GOOS、GOARCH、CGO_ENABLED 三个环境变量 (目标平台os、架构、是否启用 CGO)

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/app .

编译完成后检查：file ./ou/app


### 函数与作用域

golang 的全局变量就是一个包的全局变量
大写开头的变量才能在外部访问到