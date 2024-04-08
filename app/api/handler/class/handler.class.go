package class

import (
	"kms/app/api/handler/base"
	"kms/app/errs"
	"kms/app/usecase/class"
	"kms/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type handler struct {
	authCookie base.AuthCookieHandler

	uc class.IUseCase
}

func NewHandler(uc class.IUseCase, base base.AuthCookieHandler) *handler {
	return &handler{uc: uc, authCookie: base}
}

func (h *handler) GetClass(ctx *gin.Context) {
	const op errs.Op = "handler.class.GetClass"

	var req class.GetClassRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		msg := "bind query error: " + err.Error()
		logger.Error(msg)
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, msg))
		return
	}

	req.ID = uuid.MustParse(ctx.Param("id"))

	class, err := h.uc.GetClass(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, err))
		return
	}

	ctx.JSON(http.StatusOK, class)
}

func (h *handler) ListClasses(ctx *gin.Context) {
	const op errs.Op = "handler.auth.ListClasses"

	var req class.ListClassesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		msg := "bind query error: " + err.Error()
		logger.Error(msg)
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, msg))
		return
	}

	res, err := h.uc.ListClasses(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) CreateClass(ctx *gin.Context) {
	const op errs.Op = "handler.class.CreateClass"

	var req class.CreateClassRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := "bind json error: " + err.Error()
		logger.Error(msg)
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, msg))
		return
	}

	res, err := h.uc.CreateClass(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) UpdateClass(ctx *gin.Context) {
	const op errs.Op = "handler.class.UpdateClass"

	var req class.UpdateClassRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := "bind json error: " + err.Error()
		logger.Error(msg)
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, msg))
		return
	}

	req.ID = uuid.MustParse(ctx.Param("id"))

	res, err := h.uc.UpdateClass(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) DeleteClass(ctx *gin.Context) {
	const op errs.Op = "handler.class.DeleteClass"

	res, err := h.uc.DeleteClass(ctx, &class.DeleteClassRequest{ID: uuid.MustParse(ctx.Param("id"))})
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) AddMembersToClass(ctx *gin.Context) {
	const op errs.Op = "handler.class.AddMembersToClass"

	var req class.AddMembersToClassRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := "bind json error: " + err.Error()
		logger.Error(msg)
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, msg))
		return
	}

	res, err := h.uc.AddMembersToClass(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) RemoveMembersFromClass(ctx *gin.Context) {
	const op errs.Op = "handler.class.RemoveMembersFromClass"

	var req class.RemoveMembersFromClassRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := "bind json error: " + err.Error()
		logger.Error(msg)
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, msg))
		return
	}

	res, err := h.uc.RemoveMembersFromClass(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) ListMembersInClass(ctx *gin.Context) {
	const op errs.Op = "handler.class.ListMembersInClass"

	var req class.ListMembersInClassRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		msg := "bind query error: " + err.Error()
		logger.Error(msg)
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, msg))
		return
	}

	req.ClassID = uuid.MustParse(ctx.Param("id"))

	res, err := h.uc.ListMembersInClass(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) CheckInOut(ctx *gin.Context) {
	const op errs.Op = "handler.class.CheckInOut"

	var req class.CheckInOutRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := "bind json error: " + err.Error()
		logger.Error(msg)
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, msg))
		return
	}

	req.ClassID = uuid.MustParse(ctx.Param("id"))

	res, err := h.uc.CheckInOut(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}
