package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/ducthangng/geofleet/user-service/internal/interface/postgresql"
	"github.com/ducthangng/geofleet/user-service/internal/usercase/usecase_dto"
	"github.com/ducthangng/geofleet/user-service/service/copier"
	"github.com/ducthangng/geofleet/user-service/service/encoder"
	"github.com/jackc/pgx/v5"
)

type UserUsecaseInteractor struct {
	UserDataService postgresql.Querier
}

func NewUserUsecaseInteractor(dataService postgresql.Querier) *UserUsecaseInteractor {
	return &UserUsecaseInteractor{
		UserDataService: dataService,
	}
}

func (ui *UserUsecaseInteractor) CreateUser(ctx context.Context, dto usecase_dto.User) (
	createdUser usecase_dto.User, err error) {

	// check if user's username is duplicated, if yes then return
	duplicates, err := ui.UserDataService.GetUserByUsername(ctx, dto.Username)
	if err != pgx.ErrNoRows && err != nil {
		return createdUser, err
	}

	if len(duplicates.ID.String()) > 0 {
		return createdUser, errors.New("username already exists")
	}

	// create user
	var entityUser postgresql.CreateUserParams
	copier.MustCopy(&entityUser, &dto)

	entityUser.Password, err = encoder.HashPassword(dto.Password)
	if err != nil {
		return createdUser, err
	}

	postgreUser, err := ui.UserDataService.CreateUser(ctx, entityUser)
	if err != nil {
		return createdUser, err
	}

	// return id
	copier.MustCopy(&createdUser, &postgreUser)
	createdUser.Password = "" // empty password before sending out

	return
}

func (ui *UserUsecaseInteractor) GetUser(ctx context.Context, dto usecase_dto.User) (
	result usecase_dto.User, err error) {

	// check if user's username is duplicated, if yes then return
	currUser, err := ui.UserDataService.GetUserByUsername(ctx, dto.Username)
	if err == pgx.ErrNoRows {
		return result, errors.New("username or password is not correct 1")
	}

	if err != nil {
		return result, err
	}

	check := encoder.CheckPasswordHash(dto.Password, currUser.Password)
	log.Println("password: ", dto.Password)
	log.Println("hashed: ", currUser.Password)
	if !check {
		return result, errors.New("username or password is not correct 2")
	}

	copier.MustCopy(&result, &currUser)
	result.Password = "" // empty password before sending out

	return result, nil
}
