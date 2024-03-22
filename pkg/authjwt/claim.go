package authjwt

import "github.com/google/uuid"

// AuthClaims the claim for authentication
type AuthClaims struct {
	UID      uuid.UUID `json:"uid,omitempty"`
	Role     string    `json:"role,omitempty"`
	Username string    `json:"user_id,omitempty"`
	Email    string    `json:"email,omitempty"`
}
