package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/okikechinonso/internals/ports"
	redismemo "github.com/okikechinonso/internals/repository/redis_memo"
	"github.com/okikechinonso/internals/service"
	"github.com/okikechinonso/pkg/database"
)

type Handler struct {
	Wallet ports.WalletService
	Redis  ports.RedisRepository
}

func NewHandler() ports.IHandler {
	wallet := service.NewWalletService()
	redis := database.NewRedisDB().ConnectRedisDB()
	redisrepo := redismemo.NewRedisRepository(redis)
	return &Handler{
		Wallet: wallet,
		Redis:  redisrepo,
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
