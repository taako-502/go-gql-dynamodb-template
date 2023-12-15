package model

type User struct {
	ID   string `json:"id" dynamo:"ID,hash"`
	Name string `json:"name" dynamo:"name"`
}
