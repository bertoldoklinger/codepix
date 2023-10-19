package usecases

import (
	"errors"
	"log"

	"github.com/bertoldoklinger/codepix/domain/model"
)

type TransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
	PixRepository 				model.PixKeyRepositoryInterface
}

func (t *TransactionUseCase) Register(accountId string, amount float64, pixKeyKindTo string, pixKeyTo string, description string) (*model.Transaction, error) {
	// verifica se a conta existe
	account, err:= t.PixRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	//verifica se o pixKey existe
	pixKey, err:= t.PixRepository.FindKeyByKind(pixKeyTo, pixKeyKindTo)

	if err != nil {
		return nil, err
	}

	//valida a nova transação no model/domain com as regras de negocio
	transaction, err:= model.NewTransaction(account,amount,pixKey,description)

	if err != nil {
		return nil, err
	}

	t.TransactionRepository.Save(transaction)

	if transaction.ID != "" {
		return transaction, nil
	}

	return nil, errors.New("unable to process this transaction")
}

func (t *TransactionUseCase) Confirm(transactionId string) (*model.Transaction, error) {
	// verifica se a transação existe
	transaction, err:= t.TransactionRepository.Find(transactionId)

	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionConfirmed
	err = t.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionUseCase) Complete(transactionId string) (*model.Transaction, error) {
	// verifica se a transação existe
	transaction, err:= t.TransactionRepository.Find(transactionId)

	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}
 // muda status da transação
	transaction.Status = model.TransactionCompleted
	err = t.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionUseCase) Error(transactionId string, reason string) (*model.Transaction, error) {
	// verifica se a transação existe
	transaction, err:= t.TransactionRepository.Find(transactionId)

	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}
 // muda status da transação
	transaction.Status = model.TransactionError
	transaction.CancelDescription = reason
	
	err = t.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}