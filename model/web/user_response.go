package web

type UserResponse struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

type LoginResponse struct {
	User  *UserResponse `json:"user,omitempty"`
	Token string        `json:"token,omitempty"`
}
