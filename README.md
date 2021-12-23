# go-study

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
  - `app.Run`：传入需要启动的服务或者监听的Listener
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
	
	app.Run(iris.Addr(":8080"), iris.WithoutInterruptHandler)
}

```
