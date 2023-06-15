package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/okikechinonso/internals/service"
)

type Handler struct {
	Wallet *service.WalletService
}

func NewHandler() *Handler {
	wallet := service.NewWalletService()
	return &Handler{
		Wallet: wallet,
	}
}

func (h *Handler) GetBalance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		addr := ctx.Param("wallet_id")
		balance, code, err := h.Wallet.GetBalance(addr)
		if code != 200 {
			ctx.JSON(code, err)

			return
		}
		ctx.JSON(code, gin.H{"balance": balance})
	}
}
