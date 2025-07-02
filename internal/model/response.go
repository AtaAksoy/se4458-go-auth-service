package model

type AuthResponse struct {
	Status string      `json:"status"`
	Token  string      `json:"token,omitempty"`
	User   *UserPublic `json:"user,omitempty"`
	Error  string      `json:"error,omitempty"`
}

type UserPublic struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}
