package auth

import (
	"kms/app/api/handler/base"
	"kms/app/domain/entity"
	"kms/app/errs"
	"kms/app/usecase/auth"
	"kms/app/usecase/class"
	"kms/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	authCookie base.AuthCookieHandler

	uc      auth.IUseCase
	classUc class.IUseCase
}

func NewHandler(uc auth.IUseCase, classUc class.IUseCase, base base.AuthCookieHandler) *handler {
	return &handler{uc: uc, classUc: classUc, authCookie: base}
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

func (h *handler) Register(ctx *gin.Context) {
	const op errs.Op = "handler.auth.Register"

	var req auth.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, "bind json error: "+err.Error()))
		return
	}

	rep, err := h.uc.Register(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, rep)
}

func (h *handler) RegisterConfirm(ctx *gin.Context) {
	const op errs.Op = "handler.auth.RegisterConfirm"

	var req auth.RegisterConfirmRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, "bind json error: "+err.Error()))
		return
	}

	confirmRep, err := h.uc.RegisterConfirm(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, err)
		return
	}

	if confirmRep.Status == entity.UserRequestedStatusApproved {
		_, err = h.classUc.AddMembersToClass(ctx, &class.AddMembersToClassRequest{
			ClassID:   confirmRep.ClassID,
			Usernames: []string{confirmRep.Username},
		})
		if err != nil {
			errs.HTTPErrorResponse(ctx, err)
			return
		}
	}

	ctx.JSON(http.StatusOK, confirmRep)
}

func (h *handler) RegisterRequestList(ctx *gin.Context) {
	const op errs.Op = "handler.auth.RegisterRequestList"

	var req auth.RegisterListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		msg := "bind query error: " + err.Error()
		logger.Error(msg)
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, msg))
		return
	}
	rep, err := h.uc.RegisterList(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, rep)
}
