package author

import (
	"kms/app/errs"
	"kms/pkg/authjwt"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	role string
}

func NewAuthMiddleware(role string) gin.HandlerFunc {
	return (&AuthMiddleware{
		role: role,
	}).Handle
}

func (m *AuthMiddleware) Handle(ctx *gin.Context) {
	const op errs.Op = "middleware.author.Handle"

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.Unauthorized, "missing authorization header"))
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.Unauthorized, "invalid header format"))
		return
	}

	if headerParts[0] != "Bearer" {
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.Unauthorized, "token must content bearer"))
		return
	}

	user, err := authjwt.VerifyToken(headerParts[1])
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.Unauthorized, "invalid token"))
		return
	}

	if user.Role != m.role && user.Role != "admin" {
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.Unauthorized, "you don't have access for this action"))
		return
	}

	ctx.Set("CtxAuthenticatedUserKey", user)
	ctx.Next()
}
