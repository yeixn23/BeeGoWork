package models

type User struct {
	User     string `json:"user"`
	Birthday string `json:"birthday"`
	Address  string `json:"address"`
	Nick     string `json:"nick"`
	Password string `json:"password"`
}
