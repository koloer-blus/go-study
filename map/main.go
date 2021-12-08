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

}

func changeMap(map1 map[string]int) {
	map1["start"] = 100
}
