package auth

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetInfoRequest struct {
	Username string `json:"-"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}
