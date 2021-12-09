package main

import (
	"fmt"
	"go-study/package/action"
	"go-study/package/model"
)

func main() {
	action.GetActionList()

	owner := model.UserInfo{
		MoreInfo: map[string]interface{}{
			"work": "fe",
		},
	}
	fmt.Println(owner)
}
