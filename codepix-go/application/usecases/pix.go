package usecases

import (
	"github.com/bertoldoklinger/codepix/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}


func (p *PixUseCase) RegisterKey(key string,kind string, accountId string) (*model.PixKey, error) {
	// busca a conta e ve se ela existe	
	account, err := p.PixKeyRepository.FindAccount(accountId)

		if err != nil {
			return nil,err
		}
		// cria uma pixKey passando a conta que foi buscada, o tipo e a chave
		pixKey, err := model.NewPixKey(kind,account,key)

		if err != nil {
			return nil, err
		}

		// instancia o pixKeyRepository e usa o método pra registrar a pixKey

		p.PixKeyRepository.RegisterKey(pixKey)

		if pixKey.ID == "" {
			return nil, err
		}

		return pixKey, nil
}


func (p *PixUseCase) FindKey(key string,kind string) (*model.PixKey, error) {
		pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)

		if err != nil {
			return nil, err
		}

		return pixKey, nil
}



