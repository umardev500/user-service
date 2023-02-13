package domain

import (
	"context"
	"customer/pb"
)

type LoginResponsePayload struct {
	UserId string `bson:"user_id"`
	User   string `bson:"user"`
	Pass   string `bson:"pass"`
	Level  string `bson:"level"`
}

type UserLocation struct {
	Address    string `bson:"address,omitempty"`
	Village    string `bson:"village,omitempty"`
	District   string `bon:"district,omitempty"`
	City       string `bson:"city,omitempty"`
	Province   string `bson:"province,omitempty"`
	PostalCode string `bson:"postal_code,omitempty"`
}

type UserDetail struct {
	Name     string        `bson:"name,omitempty"`
	Email    string        `bson:"email,omitempty"`
	Phone    string        `bson:"phone,omitempty"`
	Avatar   string        `bson:"avatar,omitempty"`
	Location *UserLocation `bson:"location,omitempty"`
}

type User struct {
	UserId    string      `bson:"user_id,omitempty"`
	User      string      `bson:"user,omitempty"`
	Pass      string      `bson:"pass,omitempty"`
	Detail    *UserDetail `bson:"detail,omitempty"`
	Level     string      `bson:"level,omitempty"`
	Status    string      `bson:"status,omitempty"`
	CreatedAt int64       `bson:"created_at,omitempty"`
	UpdatedAt int64       `bson:"updated_at,omitempty"`
	DeletedAt int64       `bson:"deleted_at,omitempty"`
}

type UserRepository interface {
	Login(ctx context.Context, req *pb.UserLoginRequest) (res LoginResponsePayload, err error)
	Find(ctx context.Context, req *pb.UserFindRequest) (res *pb.User, err error)
	UpdateCreds(ctx context.Context, req *pb.UserUpdateCredsRequest, updatedTime int64) (res *pb.OperationResponse, err error)
	UpdateDetail(ctx context.Context, req *pb.UserUpdateDetailRequest) (res *pb.OperationResponse, err error)
}

type UserUsecase interface {
	Login(ctx context.Context, req *pb.UserLoginRequest) (res LoginResponsePayload, err error)
	Find(ctx context.Context, req *pb.UserFindRequest) (res *pb.User, err error)
	UpdateCreds(ctx context.Context, req *pb.UserUpdateCredsRequest) (res *pb.OperationResponse, err error)
	UpdateDetail(ctx context.Context, req *pb.UserUpdateDetailRequest) (res *pb.OperationResponse, err error)
}
