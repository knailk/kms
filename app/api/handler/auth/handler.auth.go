package auth

import (
	"kms/app/usecase/auth"

	"github.com/gin-gonic/gin"
)

type handler struct {
	uc auth.IUseCase
}

func NewHandler(uc auth.IUseCase) *handler {
	return &handler{uc: uc}
}

func (h *handler) Login(ctx *gin.Context) {

}