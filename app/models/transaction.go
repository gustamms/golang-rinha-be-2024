package models

type Transaction struct {
	Id          int
	ClientId    string
	Type        string
	Description string
	Value       string
	CreatedAt   string
}
