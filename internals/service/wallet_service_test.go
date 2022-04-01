package service

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/okikechinonso/internals/domain/entities"
	"github.com/okikechinonso/mock"
	"testing"
)

func TestWalletService_DebitWallet(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock.NewMockWalletRepository(ctrl)

	wallet := &entities.Wallet{
		Address: "1",
		Balance: 10,
	}

	t.Run("Successful", func(t *testing.T) {
		mockDB.EXPECT().UpdateBalance(wallet.Address, gomock.Any()).Return(nil)
		mockDB.EXPECT().GetWallet(wallet.Address).Return(wallet, nil)
		service := new(WalletService)
		service.Repo = mockDB
		balance, code, err := service.DebitWallet(10, wallet.Address)
		if err != nil {
			fmt.Errorf("%v", err)
		}
		assert.Equal(t, 200, code)
		assert.Equal(t, wallet.Balance, *balance)
	})

	t.Run("unsuccessful", func(t *testing.T) {
		wallet.Balance = 0
		mockDB.EXPECT().GetWallet(wallet.Address).Return(wallet, nil)
		service := new(WalletService)
		service.Repo = mockDB
		_, code, err := service.DebitWallet(10, wallet.Address)
		if err != nil {
			fmt.Errorf("%v", err)
		}
		assert.Equal(t, 400, code)
	})
}

func TestWalletService_CreditWallet(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock.NewMockWalletRepository(ctrl)

	wallet := &entities.Wallet{
		Address: "1",
		Balance: 10,
	}

	t.Run("Successful", func(t *testing.T) {
		mockDB.EXPECT().UpdateBalance(wallet.Address, gomock.Any()).Return(nil)
		mockDB.EXPECT().GetWallet(wallet.Address).Return(wallet, nil)
		service := new(WalletService)
		service.Repo = mockDB
		balance, code, err := service.CreditWallet(wallet.Address, 10)
		if err != nil {
			fmt.Errorf("%v", err)
		}
		assert.Equal(t, 200, code)
		assert.Equal(t, wallet.Balance, *balance)
	})

}
