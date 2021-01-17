package models

type Role struct{
	Id int
	Name string
	Permissions map[int64][]string
}

var AllRoles = map[int64]Role{
	1: {
		Id: 1,
		Name: "sample1",
		Permissions: map[int64][]string{1: {"read", "write"}},
	},
}