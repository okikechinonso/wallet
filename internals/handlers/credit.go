package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) Credit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		addr := ctx.Param("wallet_id")
		log.Println("here", addr)
		detail := &struct {
			Amount float64 `json:"amount"`
		}{}
		err := ctx.ShouldBindJSON(detail)
		if err != nil {

			ctx.JSON(http.StatusInternalServerError, err.Error())
			//h.ch.RequestLog(ctx)
			return
		}
		balance, code, err := h.Wallet.CreditWallet(addr, detail.Amount)
		if code != 200 {
			ctx.JSON(code, err)
			//h.ch.RequestLog(ctx)
			return
		}
		val := fmt.Sprintf("%v", balance)
		h.Wallet.Redis.Pub(&val)
		//h.ch.RequestLog(ctx)
		ctx.JSON(code, gin.H{"balance": balance})
	}
}
