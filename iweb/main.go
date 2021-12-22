package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
)

type BookController struct {
}

func main() {
	app := iris.New()
	booksApi := app.Party("/books")
	{
		booksApi.Get("/", list)
		booksApi.Post("/", create)
	}
	m := mvc.New(booksApi)
	m.Handle(new(BookController))
	err := app.Listen(":8080")
	if err != nil {
		return
	}
}

type Book struct {
	Title string `json:"title"`
}

func (c *BookController) Get() []Book {
	return []Book{
		{"Master Concurrency in Go"},
		{"GO design"},
		{"Black Hat Go"},
	}
}

func list(ctx *context.Context) {
	books := []Book{
		{"Master Concurrency in Go"},
		{"GO design"},
		{"Black Hat Go"},
	}

	json, err := ctx.JSON(books)
	if err != nil {
		return
	}
	println(json)
}
func (c *BookController) Post(b Book) int {
	println("Received Book: " + b.Title)
	return iris.StatusCreated
}
func create(ctx *context.Context) {
	var b Book
	err := ctx.ReadJSON(&b)

	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title("Book Create Fail").DetailErr(err))
	}

	println("Received Book" + b.Title)

	ctx.StatusCode(iris.StatusCreated)
}
