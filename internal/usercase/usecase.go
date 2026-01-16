package usecase

import (
	"context"

	"github.com/ducthangng/geofleet/user-service/internal/usercase/usecase_dto"
)

type UserUsecaseService interface {
	CreateUser(ctx context.Context, dto usecase_dto.User) (
		createdUser usecase_dto.User, err error)
	Login(ctx context.Context, dto usecase_dto.User) (
		result usecase_dto.User, err error)
}
