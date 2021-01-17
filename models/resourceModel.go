package models

type Resource struct {

	Id   int64  `json:"id"`
	Data string `json:"data"`
}

var AllResources = map[int64]Resource{
	1: {
		Id:   1,
		Data: "abc",
	},
}
