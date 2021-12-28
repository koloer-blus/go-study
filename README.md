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
    - 对于每一个类型T，都有一个对应的类型转换操作T(x)，用于将x转为T类型（译注：如果T是指针类型，可能会需要用小括弧包装T，比如`(*int)(0)`）。只有当两个类型的底层基础类型相同时，才允许这种转型操作，或者是两者都是指向相同底层结构的指针类型，这些转换只改变类型而不会影响值本身。
    - 数值类型之间的转型也是允许的，并且在字符串和一些特定类型的slice之间也是可以转换的，在下一章我们会看到这样的例子。这类转换可能改变值的表现。例如，将一个浮点数转为整数将丢弃小数部分，将一个字符串转为[]byte类型的slice将拷贝一个字符串数据的副本。在任何情况下，运行时不会发生转换失败的错误（错误只会发生在编译阶段）。
  - 
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
- 行为
  - `iris`默认接受和注册形如`/api/login`的路径路由，并且尾部不带斜杠。
  - 如果尝试访问`/api/login/`，将会自动永久重定向到`/api/login`
- API
  - 参数：(HTTP方法，请求的路径，多个`iris.Handler`)
```go
app := iris.New()

app.Handle("GET", "/contact", func(ctx iris.Context) {
	ctx.HTML("<h1>Hello World</h1>")
})

app.Get("/", func(ctx iris.Context) {
	ctx.HTML("<h1>hello</h1>")
})
```
- 路由组（`party`）
  - 通过对路由的路径前缀进行分组，共享相同的中间件和模板。
  - 写法1：
```go
app := iris.New()
users := app.Party("/user", handler)
users.Get("/{id:uint64}/info", handler1)
user.Get("/login", handler2)
```
  - 写法2：
```go
app.PartyFunc("/user", func(user, iris.Party) {
	user.Use(AuthMiddleware)
	user.Get("/lgoin", handler2)
})
```
- 路径参数
  - `/user/{id: string}`：`user/*`
  - `/user/{name: path}` ：`/user/**/*`
- 中间件（执行过程类似与`nodejs Express`框架）
  - 中间件仅是一个 Handler 格式的函数 `func(ctx iris.Context)`，当前一个中间件调用 `ctx.Next()` 方法时，此中间件被执行，这可以用作身份验证，即如果请求验证通过，就调用 `ctx.Next()` 来执行该请求剩下链上的处理器，否则触发一个错误响应。
- 处理`http`错误
  - Iris 内建支持 HTTP APIs 的错误详情。 
  - Context.Problem 编写一个 JSON 或者 XML 问题响应，行为完全类似 Context.JSON，但是默认 ProblemOptions.JSON 的缩进是 " "，响应的 Content-type 为 application/problem+json。 
  - 使用 options.RenderXML 和 XML 字段来改变他的行为，用 application/problem+xml 的文本类型替代。

4. MVC

![](https://www.topgoer.com/static/Iris/mvc.png)

控制器结构体内部的模型(在方法函数中设置，并通过视图渲染)。你可以从一个控制器的方法中返回模型，或者在请求的声明周期中设置一个字段，在同一个请求的生命周期中的另一个方法中返回这个字段。

就像你以前使用的流程一样，MVC 程序有自己的 Router，这是 iris/router.Party 类型的，标准的 iris api Controllers 可以被注册到任何 Party 中，包括子域名，Party 的开始和完成处理器与预期的一样工作。

可选的 BeginRequest（ctx） 函数，用于在方法执行之前执行任何初始化，这对调用中间件或许多方法使用相同的数据收集很有用。

可选的 EndRequest（ctx）函数， 可在执行任何方法之后执行任何终结处理。

递归继承，例如 我们的mvc会话控制器示例具有 Session * sessions.Session 作为字段，由会话管理器的 Start 填充为MVC应用程序的动态依赖项：mvcApp.Register(sessions.New(sessions.Config{Cookie:"iris_session_id"}).Start）

通过控制器方法的输入参数访问动态路径参数，不需要绑定。当你使用 Iris 的默认语法从一个控制器中解析处理器，你需要定义方法的后缀为 By，大写字母是新的子路径。