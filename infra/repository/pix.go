package repository

import (
	"fmt"

	"github.com/bertoldoklinger/codepix/domain/model"
	"gorm.io/gorm"
)

type PixKeyRepositoryDb struct {
	Db * gorm.DB
}

func (r PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
		err := r.Db.Create(bank).Error

		if err != nil {
			return nil
		}
		return nil
}

func (r PixKeyRepositoryDb) AddAccount(account *model.Account) error {
	err := r.Db.Create(account).Error

	if err != nil {
		return nil
	}
	return nil
}

func (r PixKeyRepositoryDb) RegisterKey(pixKey *model.PixKey) error {
	err := r.Db.Create(pixKey).Error

	if err != nil {
		return nil
	}
	return nil
}

func (r PixKeyRepositoryDb) FindKeyById(key string, kind string) (*model.PixKey, error) {
			var pixKey model.PixKey
			// preload tras todo o encadeamento (resultado da conta, e todo o resultado do banco que ta atrelado a pixKey)
			r.Db.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

			if pixKey.ID == "" {
				return nil, fmt.Errorf("no key was found")
			}

			return &pixKey, nil
}

func (r PixKeyRepositoryDb) FindAccount(id string) (*model.Account, error) {
	var account model.Account
	// preload tras todo o encadeamento (resultado da conta, e todo o resultado do banco que ta atrelado a pixKey)
	r.Db.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("no account  found")
	}

	return &account, nil
}

func (r PixKeyRepositoryDb) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank
	// preload tras todo o encadeamento (resultado da conta, e todo o resultado do banco que ta atrelado a pixKey)
	r.Db.Preload("Bank").First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("no account  found")
	}

	return &bank, nil
}

