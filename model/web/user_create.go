package web

type UserCreate struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
