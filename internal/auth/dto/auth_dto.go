package dto

type LoginErrorDto struct {
	Error bool
}

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
