package encapsulation

import (
	"fmt"
	"go-study/package/model"
)

type Product struct {
	name string
}

func PrintAuthor() {
	author := model.NewUserInfo(
		"baiziyu-fe",
		22,
		80.2,
		[]string{"Cook"},
		nil)
	fmt.Println(author)
}

func (p *Product) SetName(name string) {
	p.name = name
}
