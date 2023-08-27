package models

type Chat struct {
	Id     string
	Name   string
	UserId string
}

type Message struct {
	Id      string
	ChatId  string
	Content string
	Role    string
}
