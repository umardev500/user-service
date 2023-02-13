package usecase

import (
	"context"
	"customer/domain"
	"customer/pb"
	"time"
)

type UserUsecase struct {
	repository domain.UserRepository
}

func NewUserUsecase(repository domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}

func (u *UserUsecase) UpdateDetail(ctx context.Context, req *pb.UserUpdateDetailRequest) (res *pb.OperationResponse, err error) {
	return u.repository.UpdateDetail(ctx, req)
}

func (u *UserUsecase) UpdateCreds(ctx context.Context, req *pb.UserUpdateCredsRequest) (res *pb.OperationResponse, err error) {
	updatedTime := time.Now().UTC().Unix()
	res, err = u.repository.UpdateCreds(ctx, req, updatedTime)
	return
}

func (u *UserUsecase) Find(ctx context.Context, req *pb.UserFindRequest) (res *pb.User, err error) {
	res, err = u.repository.Find(ctx, req)
	return
}

func (u *UserUsecase) Login(ctx context.Context, req *pb.UserLoginRequest) (res domain.LoginResponsePayload, err error) {
	res, err = u.repository.Login(ctx, req)
	return
}
