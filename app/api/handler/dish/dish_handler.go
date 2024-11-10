package dish

import (
	"kms/app/api/handler/base"
	"kms/app/errs"
	"kms/app/usecase/dish"
	"kms/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type handler struct {
	authCookie base.AuthCookieHandler
	uc         dish.IUseCase
}

func NewHandler(uc dish.IUseCase, base base.AuthCookieHandler) *handler {
	return &handler{uc: uc, authCookie: base}
}

func (h *handler) CreateDish(ctx *gin.Context) {
	const op errs.Op = "handler.dish.CreateDish"

	var req dish.CreateDishRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, "bind json error: "+err.Error()))
		return
	}

	res, err := h.uc.CreateDish(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) UpdateDish(ctx *gin.Context) {
	const op errs.Op = "handler.dish.UpdateDish"

	var req dish.UpdateDishRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, "bind json error: "+err.Error()))
		return
	}

	req.DishID = uuid.MustParse(ctx.Param("id"))

	res, err := h.uc.UpdateDish(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) DeleteDish(ctx *gin.Context) {
	const op errs.Op = "handler.dish.DeleteDish"

	dishID := uuid.MustParse(ctx.Param("id"))

	err := h.uc.DeleteDish(ctx, &dish.DeleteDishRequest{DishID: dishID})
	if err != nil {
		errs.HTTPErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Dish deleted"})
}

func (h *handler) GetDishesForWeek(ctx *gin.Context) {
	const op errs.Op = "handler.dish.GetDishesForWeek"

	req := &dish.GetDishesForWeekRequest{}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		msg := "bind query error"
		logger.WithError(err).Error(msg)
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, msg))
		return
	}

	res, err := h.uc.GetDishesForWeek(ctx, req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
