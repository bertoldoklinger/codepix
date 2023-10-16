package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)


type Bank struct {
	Base `valid:"required"`
	Code string `code:"id" valid:"notnull"`
	Name string `name:"id" valid:"notnull"`
	Accounts []*Account `valid:"-"`
}

func(bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)

	if err != nil {
		return err
	}
	
	return nil
}


// Constructor que recebe codigo e nome, e retorna ou um Banco, ou um erro
func NewBank(code string, name string) (*Bank, error) {
		bank:=Bank{
			Code: code,
			Name: name,
		}
		bank.ID= uuid.NewV4().String()
		bank.CreatedAt = time.Now()
		
		err := bank.isValid()

		if err != nil {
			return nil, err
		}

		return &bank, nil
}