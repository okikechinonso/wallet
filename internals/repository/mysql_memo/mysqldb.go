package mysqlmemo

import (
	"log"

	"github.com/google/uuid"
	"github.com/okikechinonso/wallet/internals/domain/entities"
	"gorm.io/gorm"
)

type MySqlDb struct {
	MyDB *gorm.DB
}

func (d *MySqlDb) GetWallet(email string) (*entities.Wallet, error) {
	var user entities.Wallet
	err := d.MyDB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (d *MySqlDb) GetBalance(walletAddress string) float64 {
	var balance float64
	err := d.MyDB.Model(&entities.Wallet{}).Where("wallet_address = ?", walletAddress).First(&balance).Error
	if err != nil {
		log.Println(err)
	}
	return balance
}
func (d *MySqlDb) Create(wallet entities.Wallet) (string, error) {
	wallet.Address = uuid.New().String()
	result := d.MyDB.Create(&wallet)
	return wallet.Address, result.Error

}

func (d *MySqlDb) UpdateBalance(WalletAddress string, amount float64) error {
	result := d.MyDB.Model(&entities.Wallet{}).Where("wallet_address = ?", WalletAddress).Update("balance", amount)
	return result.Error
}
