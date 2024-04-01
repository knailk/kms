package chat

import (
	"kms/app/api/handler/base"
	"kms/app/domain/entity"
	"kms/app/errs"
	"kms/app/usecase/chat"
	"kms/pkg/authjwt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type handler struct {
	authCookie base.AuthCookieHandler

	uc chat.IUseCase
}

func NewHandler(uc chat.IUseCase, base base.AuthCookieHandler) *handler {
	return &handler{uc: uc, authCookie: base}
}

func (h *handler) CreateChat(ctx *gin.Context) {
	const op errs.Op = "handler.chat.CreateChat"

	var req chat.CreateChatRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, "bind json error: "+err.Error()))
		return
	}

	userClaims := ctx.MustGet(entity.CtxAuthenticatedUserKey).(*authjwt.AuthClaims)

	req.Owner = userClaims.Username

	res, err := h.uc.CreateChat(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) AddMember(ctx *gin.Context) {
	const op errs.Op = "handler.chat.AddMember"

	var req chat.AddMemberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, "bind json error: "+err.Error()))
		return
	}

	userClaims := ctx.MustGet(entity.CtxAuthenticatedUserKey).(*authjwt.AuthClaims)

	req.Adder = userClaims.Username

	res, err := h.uc.AddMember(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) ListChats(ctx *gin.Context) {
	const op errs.Op = "handler.chat.ListChats"

	userClaims := ctx.MustGet(entity.CtxAuthenticatedUserKey).(*authjwt.AuthClaims)

	res, err := h.uc.ListChats(ctx, &chat.ListChatsRequest{
		UserRequester: userClaims.Username,
	})
	if err != nil {
		errs.HTTPErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) GetChat(ctx *gin.Context) {
	const op errs.Op = "handler.chat.GetChat"

	userClaims := ctx.MustGet(entity.CtxAuthenticatedUserKey).(*authjwt.AuthClaims)

	res, err := h.uc.GetChat(ctx, &chat.GetChatRequest{
		UserRequester: userClaims.Username,
		ChatSessionID: uuid.MustParse(ctx.Param("id")),
	})
	if err != nil {
		errs.HTTPErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) UpdateChat(ctx *gin.Context) {
	const op errs.Op = "handler.chat.UpdateChat"

	var req chat.UpdateChatRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, "bind json error: "+err.Error()))
		return
	}

	req.ChatSessionID = uuid.MustParse(ctx.Param("id"))

	res, err := h.uc.UpdateChat(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) DeleteChat(ctx *gin.Context) {
	const op errs.Op = "handler.chat.DeleteChat"

	res, err := h.uc.DeleteChat(ctx, &chat.DeleteChatRequest{
		ChatSessionID: uuid.MustParse(ctx.Param("id")),
	})
	if err != nil {
		errs.HTTPErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) CreateMessage(ctx *gin.Context) {
	const op errs.Op = "handler.chat.CreateMessage"

	var req chat.CreateMessageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errs.HTTPErrorResponse(ctx, errs.E(op, errs.InvalidRequest, "bind json error: "+err.Error()))
		return
	}

	userClaims := ctx.MustGet(entity.CtxAuthenticatedUserKey).(*authjwt.AuthClaims)

	req.ChatSessionID = uuid.MustParse(ctx.Param("id"))
	req.Sender = userClaims.Username

	res, err := h.uc.CreateMessage(ctx, &req)
	if err != nil {
		errs.HTTPErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
