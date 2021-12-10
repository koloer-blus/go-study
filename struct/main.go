package main

import "fmt"

//custom design type

type Designer struct {
	name     string
	position string
	level    int
	products []string
	other    map[string]interface{}
}

func (d *Designer) getDesignerInfo() {

	var jack *Designer = new(Designer)
	jack.name = "JACK"
	fmt.Print(jack)

}

func main() {
	mark := Designer{
		name:     "Mark",
		position: "UE",
		level:    3,
		products: []string{"Athena"},
		other: map[string]interface{}{
			"addr": "China",
		},
	}
	fmt.Print(mark)
	mark.getDesignerInfo()
}
