package handler

import (
	"context"
	"log"

	identity_v1 "github.com/ducthangng/geofleet-proto/gen/go/identity/v1"
	usecase "github.com/ducthangng/geofleet/user-service/internal/usercase"
	"github.com/ducthangng/geofleet/user-service/internal/usercase/usecase_dto"
)

type UserHandler struct {
	identity_v1.UnimplementedUserServiceServer
	UserUsecase usecase.UserUsecaseService
}

func NewUserRestfulHandler(userUsecase usecase.UserUsecaseService) *UserHandler {
	return &UserHandler{
		UserUsecase: userUsecase,
	}
}

func (u *UserHandler) CreateUserProfile(ctx context.Context, data *identity_v1.CreateUserProfileRequest) (*identity_v1.CreateUserProfileResponse, error) {
	var (
		dto usecase_dto.User
		res *identity_v1.CreateUserProfileResponse
		err error
	)

	dto = usecase_dto.User{
		Fullname: data.Fullname,
		Password: data.Password.Value,
		Email:    data.Email,
		Phone:    data.Phone,
		Address:  data.Address,
	}

	log.Println("user-service receive: ", dto)
	dto, err = u.UserUsecase.CreateUser(ctx, dto)
	if err != nil {
		log.Println("user-service error: ", err)
		return nil, err
	}

	res = &identity_v1.CreateUserProfileResponse{
		UserId: dto.ID,
	}

	return res, nil
}

func (u *UserHandler) CheckDuplicatedPhone(context.Context, *identity_v1.CheckDuplicatedPhoneRequest) (*identity_v1.CheckDuplicatedPhoneResponse, error) {

	return nil, nil
}

func (u *UserHandler) GetUserProfile(context.Context, *identity_v1.GetUserProfileRequest) (*identity_v1.GetUserProfileResponse, error) {
	return nil, nil
}

func (u *UserHandler) Login(context.Context, *identity_v1.LoginRequest) (*identity_v1.LoginResponse, error) {
	return nil, nil
}
