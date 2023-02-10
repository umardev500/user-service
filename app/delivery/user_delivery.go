package delivery

import (
	"context"
	"customer/domain"
	"customer/pb"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserDelivery struct {
	usecase domain.UserUsecase
	pb.UnimplementedUserServiceServer
}

func NewUserDelivery(usecase domain.UserUsecase) *UserDelivery {
	return &UserDelivery{
		usecase: usecase,
	}
}

func (u *UserDelivery) UpdateCreds(ctx context.Context, req *pb.UserUpdateCredsRequest) (res *pb.OperationResponse, err error) {
	res, err = u.usecase.UpdateCreds(ctx, req)
	return
}

func (u *UserDelivery) Find(ctx context.Context, req *pb.UserFindRequest) (res *pb.User, err error) {
	res, err = u.usecase.Find(ctx, req)
	return
}

func (u *UserDelivery) Login(ctx context.Context, req *pb.UserLoginRequest) (res *pb.UserLoginResponse, err error) {
	result, err := u.usecase.Login(ctx, req)

	if err == mongo.ErrNoDocuments {
		return &pb.UserLoginResponse{}, nil
	}

	res = &pb.UserLoginResponse{
		UserId: result.UserId,
		User:   result.User,
		Level:  result.Level,
	}

	return
}
