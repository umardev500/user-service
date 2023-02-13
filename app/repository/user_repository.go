package repository

import (
	"context"
	"customer/domain"
	"customer/pb"
	"customer/variable"
	"fmt"
	"reflect"
	"time"

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

func (u *UserRespitory) Update(ctx context.Context, payload bson.M, userId string) (res *pb.OperationResponse, err error) {
	affected := false
	now := time.Now().UTC().Unix()

	filter := bson.M{"user_id": userId}
	payload["updated_at"] = now
	set := bson.M{"$set": payload}

	resp, err := u.user.UpdateOne(ctx, filter, set)
	if err != nil {
		return
	}

	if resp.ModifiedCount > 0 {
		affected = true
	}

	res = &pb.OperationResponse{IsAffected: affected}
	return
}

func (u *UserRespitory) UpdateDetail(ctx context.Context, req *pb.UserUpdateDetailRequest) (res *pb.OperationResponse, err error) {
	userId := req.UserId
	detail := req.Detail
	location := detail.Location
	payload := bson.M{}

	var locationValue []variable.KeyValueStruct
	if location != nil {
		locationValue = []variable.KeyValueStruct{
			{Key: "address", Value: location.Address},
			{Key: "village", Value: location.Village},
			{Key: "district", Value: location.District},
			{Key: "city", Value: location.City},
			{Key: "province", Value: location.Province},
			{Key: "postal_code", Value: location.PostalCode},
		}
	}

	var fields []variable.KeyValueStruct = []variable.KeyValueStruct{
		{Key: "detail.name", Value: detail.Name},
		{Key: "detail.email", Value: detail.Email},
		{Key: "detail.phone", Value: detail.Phone},
		{Key: "detail.avatar", Value: detail.Avatar},
		{Key: "detail.location", Value: locationValue},
	}

	for _, field := range fields {
		if !reflect.ValueOf(field.Value).IsZero() {

			if field.Key != "detail.location" {
				payload[field.Key] = field.Value
			}
			if field.Key == "detail.location" {
				for _, locationField := range field.Value.([]variable.KeyValueStruct) {
					// check for zero
					if !reflect.ValueOf(locationField.Value).IsZero() {
						payload[fmt.Sprintf("detail.location.%s", locationField.Key)] = locationField.Value
					}
				}
			}
		}
	}

	return u.Update(ctx, payload, userId)
}

func (u *UserRespitory) UpdateCreds(ctx context.Context, req *pb.UserUpdateCredsRequest, updatedTime int64) (res *pb.OperationResponse, err error) {
	userId := req.UserId
	payload := bson.M{"user": req.User, "pass": req.Pass}

	return u.Update(ctx, payload, userId)
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
		Avatar:   result.Detail.Avatar,
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

	return
}
