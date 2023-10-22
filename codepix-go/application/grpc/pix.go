package grpc

import (
	"context"

	"github.com/bertoldoklinger/codepix/application/grpc/pb"
	"github.com/bertoldoklinger/codepix/application/usecases"
)

type PixGrpcService struct {
	PixUseCase usecases.PixUseCase
	pb.UnimplementedPixServiceServer
}

func (p *PixGrpcService) RegisterPixKey(ctx context.Context, in *pb.PixKeyRegistration) (*pb.PixKeyCreatedResult, error) {
	key, err := p.PixUseCase.RegisterKey(in.Key, in.Kind, in.AccountId)

	if err != nil {
		return &pb.PixKeyCreatedResult{
			Status: "not created",
			Error: err.Error(),
	}, err
	}

	return &pb.PixKeyCreatedResult{
		Id: key.ID,
		Status: "created",
}, nil
}

func (p *PixGrpcService) Find(ctx context.Context, in *pb.PixKey) (*pb.PixKeyInfo, error) {
	pixKey, err := p.PixUseCase.FindKey(in.Key, in.Kind)

	if err != nil {
		return &pb.PixKeyInfo{}, err
	}

	return &pb.PixKeyInfo{
		Id: pixKey.ID,
		Kind: pixKey.Kind,
		Key: pixKey.Key,
		Account: &pb.Account{
			AccountId: pixKey.AccountID,
			AccountNumber: pixKey.Account.Number,
			BankId: pixKey.Account.BankID,
			OwnerName: pixKey.Account.OwnerName,
			CreatedAt: pixKey.Account.CreatedAt.String(),
		},
		CreatedAt: pixKey.CreatedAt.String(),
	}, nil

}

func NewPixGrpcService(usecases usecases.PixUseCase) *PixGrpcService {
	return &PixGrpcService{
		PixUseCase: usecases,
	}
}