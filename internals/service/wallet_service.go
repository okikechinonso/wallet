package service

import (
	"errors"
	"log"
	"net/http"

	"github.com/okikechinonso/internals/domain/entities"
	"github.com/okikechinonso/internals/ports"
	mysqlmemo "github.com/okikechinonso/internals/repository/mysql_memo"
	redismemo "github.com/okikechinonso/internals/repository/redis_memo"
	"github.com/okikechinonso/pkg/database"
)

type WalletService struct {
	Repo  ports.WalletRepository
	Redis ports.RedisRepository
}

func NewWalletService() *WalletService {
	db := database.ConnectDB()

	repo := mysqlmemo.NewWalletRepository(db)
	redis := database.ConnectRedisDB()
	redisrepo := redismemo.NewRedisRepository(redis)
	return &WalletService{
		Repo:  repo,
		Redis: redisrepo,
	}
}

func (w *WalletService) GetWallet(addr string) (*entities.Wallet, int, error) {
	wallet, err := w.Repo.GetWallet(addr)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return wallet, http.StatusOK, nil
}

func (w *WalletService) DebitWallet(amount float64, addr string) (*float64, int, error) {
	if amount < 0 {
		return nil, http.StatusBadRequest, errors.New("enter valid amount")
	}
	wallet, err := w.Repo.GetWallet(addr)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	if wallet.Balance < 0 {
		return nil, http.StatusBadRequest, errors.New("insufficient fund")
	}
	if wallet.Balance < amount {
		return nil, http.StatusBadRequest, errors.New("insufficient fund")
	}
	wallet.Balance -= amount
	err = w.Repo.UpdateBalance(wallet.Address, wallet.Balance)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	return &wallet.Balance, http.StatusOK, err
}

func (w *WalletService) CreditWallet(addr string, amount float64) (*float64, int, error) {
	if amount < 0 {
		return nil, http.StatusBadRequest, errors.New("enter a valid amount")
	}
	wallet, err := w.Repo.GetWallet(addr)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	wallet.Balance += amount
	err = w.Repo.UpdateBalance(wallet.Address, wallet.Balance)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	return &wallet.Balance, http.StatusOK, nil
}

func (w *WalletService) CreateWallet(wallet entities.Wallet) (string, int, error) {
	_, err := w.Repo.GetWallet(wallet.Address)
	if err == nil {
		return "", http.StatusBadRequest, err
	}
	addr, err := w.Repo.Create(wallet)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	return addr, http.StatusOK, nil
}

func (w *WalletService) GetBalance(addr string) (*float64, int, error) {
	val, err := w.Redis.Get("balance")
	if err != nil {
		wallet, code, err := w.GetWallet(addr)
		log.Println(wallet.Balance)
		if err != nil {
			return nil, code, err
		}
		err = w.Redis.Set("balance", wallet.Balance)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
		val = &wallet.Balance
	}
	return val, http.StatusOK, nil
}
