package author

import (
	"kms/app/domain/entity"
	"kms/app/errs"
	"kms/pkg/authjwt"

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

	token, err := ctx.Cookie(entity.AccessKey)
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.Unauthorized, "missing authorization token"))
		return
	}

	user, err := authjwt.VerifyToken(token)
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.Unauthorized, "invalid token"))
		return
	}

	if user.Role != m.role && user.Role != "admin" {
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.Unauthorized, "you don't have access for this action"))
		return
	}

	ctx.Set(entity.CtxAuthenticatedUserKey, user)
	ctx.Next()
}
