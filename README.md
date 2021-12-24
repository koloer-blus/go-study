# go-study

## golang

### 基础

1. 基本命令
```shell
go run main.go # >>> 执行go文件
go build main.go # >>> 生成可执行的二进制文件
go get {url} # >>> 获取网上的代码

```
2. 名词解释

- 包（`package`）:  一个包由位于单个目录下的一个或多个.go源代码文件组成，目录定义包的作用；
  - `main`包：定义一个可以独立执行的函数，而不是一个库，是程序的入口
  - 每一个包都有对应独立的命名空间
  - 通过名字大小写控制变量的私有性（变量名小写字母开头代表小写）
  - 包的第一句注释应当是报的功能概要说明
  - 包的初始化可以使用`func init() {}`来执行，该函数同时不支持被调用或引用，在程序开始执行时按照声明顺序依次调用

3. 程序结构

- 声明(<var|const> name [types] [= value])
  - `var`：定义变量
  - `const`：定义常量
  - `type`：类型
  - `func`：函数实体
- 简短声明
  -`:=`：`temp := 2`
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
- 指针
  - 一个指针的值是另一个变量的地址。一个指针对应变量在内存中的存储位置。并不是每一个值都会有一个内存地址，但是对于每一个变量必然有对应的内存地址。
  - 对于聚合类型每个成员——比如结构体的每个字段、或者是数组的每个元素——也都是对应一个变量，因此可以被取地址。 
  - 任何类型的指针零值都是`nil`
```go
x := 1
p := &x
fmt.println(*p)
```
- `new`函数
  - 表达式`new(T)`将创建一个T类型的匿名变量，初始化为`T`类型的零值，然后**返回变量地址**，**返回的指针类型为`*T`**。
  - new函数类似是一种语法糖，而不是一个新的基础概念
  - 每次调用`new`函数都会返回一个新的变量地址，如果两个类型都是空的，也就是说类型的大小是0，例如`struct{}`和`[0]int`，有可能有相同的地址（依赖具体的语言实现）,请谨慎使用大小为0的类型，因为如果类型的大小为0的话，可能导致Go语言的自动垃圾回收器有不同的行为
```go
p := new(int)
fmt.println(*p)

p := new(int)
q := new(int)
fmt.println(p == q) //false
```

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

6. Go语言的自动垃圾收集器
> 从每个包级的变量和每个当前运行函数的每一个局部变量开始，通过指针或引用的访问路径遍历，是否可以找到该变量。如果不存在这样的访问路径，那么说明该变量是不可达的，也就是说它是否存在并不会影响程序后续的计算结果。




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

3. Router

- Handler Type(请求处理器)
  - 处理过程：响应Http请求 -> 写入响应头和数据到`Context.ResponseWriter()` -> 返回信号 -> 请求处理完成
  - 注意事项：
    - 提前读取`Context.Request().Body`中的数据，因为**在写入`Context.ResponseWriter()`后无法访问`Context.Request().Body`**
    - Handler不因改变传入的Context
    - 服务器出现`panic(异常)`，服务器会认为当前的panic的影响与运行的请求无关。会重启当前的panic，并且记录栈追踪日志到服务器错误日志同时关闭连接。