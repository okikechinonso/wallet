package ports

import "github.com/okikechinonso/internals/domain/entities"

type WalletRepository interface {
	GetWallet(addr string) (*entities.Wallet, error)
	GetBalance(walletAddress string) float64
	Create(wallet entities.Wallet) (string, error)
	UpdateBalance(WalletAddress string, amount float64) error
}
