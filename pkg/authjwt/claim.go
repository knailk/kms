package authjwt

// AuthClaims the claim for authentication
type AuthClaims struct {
	UID    string `json:"uid,omitempty"`
	Role   string `json:"role,omitempty"`
	UserID string `json:"user_id,omitempty"`
	Email  string `json:"email,omitempty"`
}
