package main

import "fmt"

func main() {
	map1 := make(map[string]string)
	map1["hello"] = "go"
	map1["numbers"] = "3232"

	fmt.Println(map1)

	map2 := map[string]int{
		"start": 0,
		"end":   1,
		"false": -1,
	}
	fmt.Println(map2)

	map3 := map[int][]string{}
	slice1 := make([]string, 3)
	slice1[0] = "Hello"
	slice1[1] = "Test"
	map3[0] = slice1
	fmt.Print(map3)

	changeMap(map2)
	fmt.Print(map2)

	slice := []string{}
	for k, v := range map2 {
		fmt.Println(k, v)
		slice = append(slice, k)
	}
	fmt.Print(slice)

	type goObj struct {
		key      string
		val      string
		enabled  int16
		onlyRead bool
	}

	object1 := goObj{
		key:      "name",
		val:      "fe",
		enabled:  3,
		onlyRead: true,
	}

	map4 := make(map[string]goObj)
	map4["start"] = object1
	fmt.Print(map4)

	value, enable := map4["start"]
	if enable {
		fmt.Print(value)
	} else {
		fmt.Print("no value")
	}

	if v, enable := map4["start"]; enable {
		fmt.Println(v)
	}

	var map5 []map[string]interface{}

	map5 = make([]map[string]interface{}, 2)
	map5[0] = make(map[string]interface{}, 2)
	map5[0]["name"] = "baiziyu-fe"
	map5[0]["email"] = "test@gmail.com"

}

func changeMap(map1 map[string]int) {
	map1["start"] = 100
}
