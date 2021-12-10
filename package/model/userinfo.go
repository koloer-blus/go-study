package model

type UserInfo struct {
	name     string
	age      int
	weight   float32
	hobby    []string
	MoreInfo map[string]interface{}
}

func NewUserInfo(
	name string,
	age int,
	weight float32,
	hobby []string,
	MoreInfo map[string]interface{},
) *UserInfo {
	return &UserInfo{
		name:     name,
		age:      age,
		weight:   weight,
		hobby:    hobby,
		MoreInfo: MoreInfo,
	}
}
