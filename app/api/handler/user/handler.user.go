package user

import (
	"kms/app/api/handler/base"
	"kms/app/domain/entity"
	"kms/app/errs"
	"kms/app/usecase/user"
	"kms/pkg/authjwt"
	"kms/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	authCookie base.AuthCookieHandler

	uc user.IUseCase
}

func NewHandler(uc user.IUseCase, base base.AuthCookieHandler) *handler {
	return &handler{uc: uc, authCookie: base}
}

func (h *handler) GetProfile(ctx *gin.Context) {
	const op errs.Op = "handler.user.GetProfile"

	userClaims := ctx.MustGet(entity.CtxAuthenticatedUserKey).(*authjwt.AuthClaims)

	user, err := h.uc.GetUser(ctx, &user.GetUserRequest{Username: userClaims.Username})
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *handler) UpdateUser(ctx *gin.Context) {
	const op errs.Op = "handler.user.UpdateUser"

	userClaims := ctx.MustGet(entity.CtxAuthenticatedUserKey).(*authjwt.AuthClaims)

	var req user.UpdateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := "bind json error: " + err.Error()
		logger.Error(msg)
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, msg))
		return
	}

	req.Username = userClaims.Username

	res, err := h.uc.UpdateUser(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) SearchUser(ctx *gin.Context) {
	const op errs.Op = "handler.user.SearchUser"

	var req user.SearchUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		msg := "bind query error: " + err.Error()
		logger.Error(msg)
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, msg))
		return
	}

	res, err := h.uc.SearchUser(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) ListTeachersAvailable(ctx *gin.Context) {
	const op errs.Op = "handler.user.ListTeachersAvailable"

	res, err := h.uc.ListTeachersAvailable(ctx, &user.ListTeachersAvailableRequest{})
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) ListDriversAvailable(ctx *gin.Context) {
	const op errs.Op = "handler.user.ListDriversAvailable"

	res, err := h.uc.ListDriversAvailable(ctx, &user.ListDriversAvailableRequest{})
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}
