package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKey struct {
		Base  `valid:"required"`
		Kind string	`json:"kind" valid:"notnull"`
		Key string	`json:"key" valid:"notnull"`
		AccountID string `json:"account_id" valid:"notnull"`
		Account *Account	`valid:"-"`
		Status string  `json:"status" valid:"notnull"`
}

type IPixKeyRepository interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccout(id string) (*Account, error)
}


func(pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)

	if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
		return errors.New("invalid type of key")
	}

	if pixKey.Status != "active" && pixKey.Kind != "inactive" {
		return errors.New("invalid status")
	}


	if err != nil {
		return err
	}
	
	return nil
}


func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {
	pixKey := PixKey{
		Kind: kind,
		Key: 			key, 			
		Account: account,
		Status: "active",
	}
	pixKey.ID= uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()
	
	err := pixKey.isValid()

	if err != nil {
		return nil, err
	}

	return &pixKey, nil
}
