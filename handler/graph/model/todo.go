package model

type Todo struct {
	ID   string `json:"id" dynamo:"ID,hash"`
	Text string `json:"text" dynamo:"text"`
	Done bool   `json:"done" dynamo:"done"`
	User *User  `json:"user" dynamo:"user"`
}
