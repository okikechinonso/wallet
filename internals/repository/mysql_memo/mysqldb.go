package mysqlmemo

import (
	"github.com/google/uuid"
	"github.com/okikechinonso/internals/domain/entities"
	"github.com/okikechinonso/internals/ports"
	"gorm.io/gorm"
	"log"
)

type sqLRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) ports.WalletRepository {
	return &sqLRepository{
		db: db,
	}
}

func (d *sqLRepository) GetWallet(addr string) (*entities.Wallet, error) {
	var user entities.Wallet
	err := d.db.Where("address = ?", addr).First(&user).Error
	return &user, err
}

func (d *sqLRepository) GetBalance(walletAddress string) float64 {
	var balance float64
	err := d.db.Model(&entities.Wallet{}).Where("wallet_address = ?", walletAddress).First(&balance).Error
	if err != nil {
		log.Println(err)
	}
	return balance
}
func (d *sqLRepository) Create(wallet entities.Wallet) (string, error) {
	wallet.Address = uuid.New().String()
	result := d.db.Create(&wallet)
	return wallet.Address, result.Error

}

func (d *sqLRepository) UpdateBalance(WalletAddress string, amount float64) error {
	result := d.db.Model(&entities.Wallet{}).Where("address = ?", WalletAddress).Update("balance", amount)
	return result.Error
}
