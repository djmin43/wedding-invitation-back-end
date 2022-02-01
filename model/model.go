package model

type Blog struct {
	Id          string `json:"id"`
	User        string `json:"user"`
	Body        string `json:"body"`
	AvatarColor string `json:"avatarColor"`
	CreateDt    string `json:"created"`
}