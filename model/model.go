package model

type User struct {
	Username string `json:"username"`
	Follower int    `json:"followers"`
}
