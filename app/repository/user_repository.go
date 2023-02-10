package repository

import (
	"context"
	"customer/domain"
	"customer/pb"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRespitory struct {
	user *mongo.Collection
}

func NewUserRepository(user *mongo.Collection) domain.UserRepository {
	return &UserRespitory{
		user: user,
	}
}

func (u *UserRespitory) Find(ctx context.Context, req *pb.UserFindRequest) (res *pb.User, err error) {
	var result domain.User

	userId := req.UserId
	filter := bson.M{"user_id": userId}
	err = u.user.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	var location *pb.UserLocation

	if result.Detail.Location != nil {
		location = &pb.UserLocation{
			Address:    result.Detail.Location.Address,
			Village:    result.Detail.Location.Village,
			District:   result.Detail.Location.District,
			City:       result.Detail.Location.City,
			Province:   result.Detail.Location.Province,
			PostalCode: result.Detail.Location.PostalCode,
		}
	}

	detail := &pb.UserDetail{
		Name:     result.Detail.Name,
		Email:    result.Detail.Email,
		Phone:    result.Detail.Phone,
		Location: location,
	}

	res = &pb.User{
		UserId:    result.UserId,
		User:      result.User,
		Pass:      result.Pass,
		Detail:    detail,
		Level:     result.Level,
		Status:    result.Status,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	return
}

func (u *UserRespitory) Login(ctx context.Context, req *pb.UserLoginRequest) (result domain.LoginResponsePayload, err error) {
	user := req.User
	pass := req.Pass

	filter := bson.M{"user": user, "pass": pass}
	err = u.user.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Println("error no document")
	}

	return
}
