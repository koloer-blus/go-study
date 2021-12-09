package model

type UserInfo struct {
	name     string
	age      int
	weight   float32
	hobby    []string
	MoreInfo map[string]interface{}
}
