package base

import (
	"kms/app/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// authCookieHandler hold auth Cookie Handler props
type authCookieHandler struct {
	accessKey              string
	refreshKey             string
	cookiePath             string
	cookieRefreshTokenPath string

	cookieHTTPOnly bool
	cookieSameSite http.SameSite
	cookieSecure   bool
	cookieMaxAge   int
}

// AuthCookieHandler hold auth Cookie Handler interface
type AuthCookieHandler interface {
	// GetAccessToken(ctx *gin.Context) (token string, err error)
	// GetRefreshToken(ctx *gin.Context) (token string, err error)
	SetAccessCookie(ctx *gin.Context, token string)
	SetRefreshCookie(ctx *gin.Context, token string)
	ExpireJWTCookie(ctx *gin.Context)
}

// NewAuthCookieHandler return new auth cookie handler instance
func NewAuthCookieHandler(
	ENV string,
	accessKey string, refreshKey string,
	cookiePath string, cookieRefreshTokenPath string,
	cookieHTTPOnly bool, cookieMaxAge int,
) AuthCookieHandler {
	var (
		cookieSameSite = http.SameSiteLaxMode
		cookieSecure   = true
	)

	if ENV == config.EnvDevelopment || ENV == config.EnvStaging {
		cookieSameSite = http.SameSiteNoneMode
	}
	if ENV == config.EnvLocal {
		cookieSecure = false
	}
	return &authCookieHandler{
		accessKey:              accessKey,
		refreshKey:             refreshKey,
		cookiePath:             cookiePath,
		cookieRefreshTokenPath: cookieRefreshTokenPath,
		cookieHTTPOnly:         cookieHTTPOnly,
		cookieSameSite:         cookieSameSite,
		cookieSecure:           cookieSecure,
		cookieMaxAge:           cookieMaxAge,
	}
}

// // GetAccessToken gets the access token.
// func (h *authCookieHandler) GetAccessToken(ctx *gin.Context) (token string, err error) {
// 	if token, err = ctx.Cookie(h.accessKey); err != nil {
// 		return "", apperrors.ErrUnauthorized
// 	}
// 	return token, nil
// }

// // GetRefreshToken gets the refresh token.
// func (h *authCookieHandler) GetRefreshToken(ctx *gin.Context) (token string, err error) {
// 	if token, err = ctx.Cookie(h.refreshKey); err != nil {
// 		return "", apperrors.ErrUnauthorized
// 	}
// 	return token, nil
// }

// SetAccessCookie sets the access cookie.
func (h *authCookieHandler) SetAccessCookie(ctx *gin.Context, token string) {
	h.setTokenWithAge(ctx, h.accessKey, h.cookiePath, token, h.cookieMaxAge)
}

// SetRefreshCookie sets the refresh cookie.
func (h *authCookieHandler) SetRefreshCookie(ctx *gin.Context, token string) {
	h.setTokenWithAge(ctx, h.refreshKey, h.cookieRefreshTokenPath, token, h.cookieMaxAge)
}

// ExpireJWTCookie expires the JWT cookie.
func (h *authCookieHandler) ExpireJWTCookie(ctx *gin.Context) {
	h.setTokenWithAge(ctx, h.accessKey, h.cookiePath, "", 0)
	h.setTokenWithAge(ctx, h.refreshKey, h.cookieRefreshTokenPath, "", 0)
}

func (h *authCookieHandler) setTokenWithAge(ctx *gin.Context, key, path, token string, age int) {
	ctx.SetSameSite(h.cookieSameSite)
	ctx.SetCookie(key, token, age, path, "", h.cookieSecure, h.cookieHTTPOnly)
}
