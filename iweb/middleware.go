package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {
	app := iris.New()
	app.Get("/", before, middleware, after)
	app.OnErrorCode(iris.StatusNotFound, notFound)
}

func notFound(ctx *context.Context) {
	_, _ = ctx.WriteString("Not found")
}

func before(ctx *context.Context) {
	share := "This is before"

	requestPath := ctx.Path()
	println("Before the mainHandler: " + requestPath)
	ctx.Values().Set("info", share)
	ctx.Next()
}

func after(ctx *context.Context) {
	println("after")
}

func middleware(ctx *context.Context) {
	println("Inside the mainHandler")
	info := ctx.Values().GetString("info")
	_, _ = ctx.HTML("<h1>Res</h1>")
	_, _ = ctx.HTML("<br /> Info :" + info)
	ctx.Next()
}
