package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Debit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		addr := ctx.Param("wallet_id")
		detail := &struct {
			Amount float64 `json:"amount"`
		}{}
		err := ctx.ShouldBindJSON(detail)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())

			return
		}
		balance, code, err := h.Wallet.DebitWallet(detail.Amount, addr)
		if code != 200 {
			ctx.JSON(code, err)
			return
		}
		val := fmt.Sprintf("%v", balance)
		h.Redis.Pub(&val)
		ctx.JSON(code, gin.H{"balance": balance})
	}
}
