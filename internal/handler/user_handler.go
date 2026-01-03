package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/ducthangng/geofleet/user-service/internal/handler/presenter"
	usecase "github.com/ducthangng/geofleet/user-service/internal/usercase"
	"github.com/ducthangng/geofleet/user-service/internal/usercase/usecase_dto"
	"github.com/ducthangng/geofleet/user-service/service/copier"
	jwtService "github.com/ducthangng/geofleet/user-service/service/jwt"
	"github.com/ducthangng/geofleet/user-service/singleton"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserRestfulHandler struct {
	BaseHandler
	UserUsecase usecase.UserUsecaseService
}

func NewUserRestfulHandler(usecase usecase.UserUsecaseService) *UserRestfulHandler {
	return &UserRestfulHandler{
		UserUsecase: usecase,
	}
}

func (u *UserRestfulHandler) Register(ctx *gin.Context) {
	var (
		req presenter.User
		dto usecase_dto.User
		res usecase_dto.User
		err error
	)

	defer func() {
		if err != nil {
			u.SetError(ctx, err)
		}
	}()

	if err = ctx.ShouldBindJSON(&req); err != nil {
		return
	}

	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		return
	}

	copier.MustCopy(&dto, &req)
	res, err = u.UserUsecase.CreateUser(ctx, dto)
	if err != nil {
		return
	}

	u.SetData(ctx, res)
	u.SetMeta(ctx, presenter.MetaResponse{
		Code: http.StatusCreated,
	})
}

func (u *UserRestfulHandler) Login(ctx *gin.Context) {
	var (
		req      presenter.User
		dto      usecase_dto.User
		res      usecase_dto.User
		jwtToken string
		err      error
	)

	defer func() {
		if err != nil {
			u.SetError(ctx, err)
		}
	}()

	if err = ctx.ShouldBindJSON(&req); err != nil {
		return
	}

	if (!req.VerifyUsername()) || (!req.VerifyPassword()) {
		err = errors.New("username or password is not correct 3")
		return
	}

	copier.MustCopy(&dto, &req)
	res, err = u.UserUsecase.GetUser(ctx, dto)
	if err != nil {
		return
	}

	if len(res.ID) == 0 {
		err = errors.New("username or password is not correct 4")
		return
	}

	// set credentials
	jwtToken, err = jwtService.GenerateToken(res.ID, res.Username, "")
	if err != nil {
		log.Println("here 1")
		return
	}

	cfg := singleton.GetConfig().Cookie

	// set cookie
	ctx.SetCookie(cfg.CookieName, jwtToken, cfg.MaxAge, "/", cfg.CookieDomain, cfg.CookieSecure, cfg.CookieHTTPOnly)

	u.SetData(ctx, res)
	u.SetMeta(ctx, presenter.MetaResponse{
		Code: http.StatusCreated,
	})
}

func (u *UserRestfulHandler) GetMyself(ctx *gin.Context) {
	var (
		req presenter.User
		dto usecase_dto.User
		res usecase_dto.User
		err error
	)

	defer func() {
		if err != nil {
			u.SetError(ctx, err)
		}
	}()

	if err = ctx.ShouldBindJSON(&req); err != nil {
		return
	}

	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		return
	}

	copier.MustCopy(&dto, &req)
	res, err = u.UserUsecase.CreateUser(ctx, dto)
	if err != nil {
		return
	}

	u.SetData(ctx, res)
	u.SetMeta(ctx, presenter.MetaResponse{
		Code: http.StatusCreated,
	})
}
