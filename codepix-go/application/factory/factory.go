package factory

import (
	"github.com/bertoldoklinger/codepix/application/usecases"
	"github.com/bertoldoklinger/codepix/infra/repository"
	"github.com/jinzhu/gorm"
)

func TransactionUseCaseFactory(database *gorm.DB) usecases.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	transactionRepository := repository.TransactionRepositoryDb{Db: database}

	transactionUseCase := usecases.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixRepository: pixRepository,
	}

	return transactionUseCase
}