package auth

import (
	"kms/app/api/handler/base"
	"kms/app/errs"
	"kms/app/usecase/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	authCookie base.AuthCookieHandler

	uc auth.IUseCase
}

func NewHandler(uc auth.IUseCase, base base.AuthCookieHandler) *handler {
	return &handler{uc: uc, authCookie: base}
}

func (h *handler) Login(ctx *gin.Context) {
	const op errs.Op = "handler.auth.Login"

	var req auth.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, "bind json error: "+err.Error()))
		return
	}

	user, err := h.uc.Login(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, err)
		return
	}

	h.authCookie.SetAccessCookie(ctx, user.AccessToken.Token)
	h.authCookie.SetRefreshCookie(ctx, user.RefreshToken.Token)

	ctx.JSON(http.StatusOK, user)
}
