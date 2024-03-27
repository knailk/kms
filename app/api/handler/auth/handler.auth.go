package auth

import (
	"kms/app/api/handler/base"
	"kms/app/domain/entity"
	"kms/app/errs"
	"kms/app/usecase/auth"
	"kms/pkg/authjwt"
	"kms/pkg/logger"
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

func (h *handler) Refresh(ctx *gin.Context) {
	const op errs.Op = "handler.auth.Refresh"

	token, err := ctx.Cookie(entity.AccessKey)
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.Unauthorized, "missing authorization refresh token"))
		return
	}

	tokenPair, err := h.uc.Refresh(ctx, &auth.RefreshRequest{RefreshToken: token})
	if err != nil {
		errs.HTTPErrorResponse(ctx, err)
		return
	}

	h.authCookie.SetAccessCookie(ctx, tokenPair.AccessToken.Token)
	h.authCookie.SetRefreshCookie(ctx, tokenPair.RefreshToken.Token)

	ctx.JSON(http.StatusOK, "OK")
}

// Logout log user out
func (h *handler) Logout(ctx *gin.Context) {
	h.authCookie.ExpireJWTCookie(ctx)
	ctx.JSON(http.StatusOK, "OK")
}

func (h *handler) GetProfile(ctx *gin.Context) {
	const op errs.Op = "handler.auth.GetInfo"

	userClaims := ctx.MustGet(entity.CtxAuthenticatedUserKey).(*authjwt.AuthClaims)

	user, err := h.uc.GetProfile(ctx, &auth.GetInfoRequest{Username: userClaims.Username})
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *handler) UpdateProfile(ctx *gin.Context) {
	const op errs.Op = "handler.auth.UpdateProfile"

	userClaims := ctx.MustGet(entity.CtxAuthenticatedUserKey).(*authjwt.AuthClaims)

	var req auth.UpdateProfileRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := "bind json error: " + err.Error()
		logger.Error(msg)
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, msg))
		return
	}

	req.Username = userClaims.Username

	res, err := h.uc.UpdateProfile(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}
