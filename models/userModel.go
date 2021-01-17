package models


type User struct {
	Id int64
	Name string
	RoleIds []int
}

var AllUsers = map[int64]User{
	1: {
		Id: 1,
		Name: "sarat",
		RoleIds: []int{},
	},
}