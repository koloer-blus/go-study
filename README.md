# go-study

## 基础

1. 基本命令
```shell
go run main.go # >>> 执行go文件
go build main.go # >>> 生成可执行的二进制文件
go get {url} # >>> 获取网上的代码

```
2. 名词解释

- 包（`package`）:  一个包由位于单个目录下的一个或多个.go源代码文件组成，目录定义包的作用；
  - `main`包：定义一个可以独立执行的函数，而不是一个库，是程序的入口

3. 变量赋值

- 定义变量
  - `var`：定义变量
  - `const`：定义常量
- 初始化隐形赋值
  - `string` >>> ""(空字符串)
  - `int` >>> 0
  ```go
  s := ""   // >>> 简短声明，只能在函数内部声明
  var s string
  var s = ""
  var s string = ""
  ```
- `map`：`Go`的`map`类似于`Java`语言中的`HashMap`

4. 循环
 - `for...range variable {}`

```go
  for a,b := range "hello" {
    fmt.println(a,b)
  }
```

5. `Printf`动词展示

```
%d  十进制整数
%x,%o,%b  十六进制，八进制，二进制
%f,%g,%e  浮点数
%t  bool
%c  char
%s  string
%q  string or char
%v  变量的默认形式
%T  变量的类型
%%  %  
```

## 实战

1. [gee-web](./gee-web)

## 框架

### Iris

1. Get `Iris`

- 使用命令行
```shell
# install iris v12
go get github.com/kataras/iris/v12@lastest
```
- 修改 `go.mod`

```go
module your_project_name

go 1.14

require (
    github.com/kataras/iris/v12 v12.1.8
)
```

2. 入门

- 创建新的服务
  - `iris.New()`：返回一个可配置的`iris.Application`实例
  - `app.Run`：第一个参数为需要启动的服务或者监听的Listener，第二个及之后的参数是可选的`iris.Configurator`配置参数。**在`iris`中每一个核心的模块（视图引擎、websocket、session等）都有一个内部的`iris.Configurator`**
  - `app.Listen`：传入需要监听的端口号
```go
app := iris.New()
// 这行代码会启动一个服务并监听 localhost:8080或者127.0.0.1:8080
app.Listen(":8080")


/*-------------------------------------*/

// 当然，想要获取完整的`http.Server`的实例的使用也是可以的
import "net/http"

// 这里实现的功能和上面的代码没有差异，但是可以确保你使用完整的`http.Server`实例
server := &http.Server{Addr: ":8080"}
app.Run(iris.Server(server))

/*-------------------------------------*/

// 还有更好的用法就是使用自定义的`net.Listener`

listener, err := net.Listen("tcp4", ":8080")
if err != nil {
	panic(err)
}
app.Run(iris.Listener(listener))
```
- 关闭服务并停止默认行为

```go
package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"time"
)

func main() {
	app := iris.New().
		iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		app.Shutdown(ctx)
	})
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Closed</h1>")
    })
	// iris 配置
	confit := iris.WithConfiguration(iris.Configuration{
		DisableStartupLog: true,
		Charset: "UTF-8",
    })
	app.Run(iris.Addr(":8080"), config)
}

```
