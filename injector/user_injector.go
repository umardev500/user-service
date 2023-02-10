package injector

import (
	"customer/app/delivery"
	"customer/app/repository"
	"customer/app/usecase"
	"customer/config"
)

func NewUserInjector() *delivery.UserDelivery {
	client := config.Conn()
	db := client.Database("db_user")
	user := db.Collection("users")

	ur := repository.NewUserRepository(user)
	uc := usecase.NewUserUsecase(ur)
	return delivery.NewUserDelivery(uc)
}
