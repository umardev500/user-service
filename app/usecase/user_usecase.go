package usecase

import (
	"context"
	"customer/domain"
	"customer/pb"
)

type UserUsecase struct {
	repository domain.UserRepository
}

func NewUserUsecase(repository domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}

func (u *UserUsecase) Find(ctx context.Context, req *pb.UserFindRequest) (res *pb.User, err error) {
	res, err = u.repository.Find(ctx, req)
	return
}

func (u *UserUsecase) Login(ctx context.Context, req *pb.UserLoginRequest) (res domain.LoginResponsePayload, err error) {
	res, err = u.repository.Login(ctx, req)
	return
}
