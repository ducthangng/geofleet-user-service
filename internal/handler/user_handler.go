package handler

import (
	"context"
	"log"

	pb "github.com/ducthangng/geofleet-proto/user"
	usecase "github.com/ducthangng/geofleet/user-service/internal/usercase"
	"github.com/ducthangng/geofleet/user-service/internal/usercase/usecase_dto"
	"github.com/ducthangng/geofleet/user-service/service/copier"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	UserUsecase usecase.UserUsecaseService
}

func NewUserRestfulHandler(userUsecase usecase.UserUsecaseService) *UserHandler {
	return &UserHandler{
		UserUsecase: userUsecase,
	}
}

func (u *UserHandler) CreateUserProfile(ctx context.Context, data *pb.UserCreationRequest) (*pb.UserCreationResponse, error) {
	var (
		dto usecase_dto.User
		res *pb.UserCreationResponse
		err error
	)

	copier.MustCopy(&dto, data)
	log.Println("user-service receive: ", dto)

	dto, err = u.UserUsecase.CreateUser(ctx, dto)
	if err != nil {
		log.Println("user-service error: ", err)
		return nil, err
	}

	res = &pb.UserCreationResponse{
		UserId:       dto.ID,
		IsDuplicated: 0,
	}

	return res, nil
}

func (u *UserHandler) Register(ctx *gin.Context) {

}

// func (u *UserHandler) Login(ctx *gin.Context) {
// 	var (
// 		req      presenter.User
// 		dto      usecase_dto.User
// 		res      usecase_dto.User
// 		jwtToken string
// 		err      error
// 	)

// 	if err = ctx.ShouldBindJSON(&req); err != nil {
// 		return
// 	}

// 	if (!req.VerifyUsername()) || (!req.VerifyPassword()) {
// 		err = errors.New("username or password is not correct 3")
// 		return
// 	}

// 	copier.MustCopy(&dto, &req)
// 	// res, err = u.UserUsecase.GetUser(ctx, dto)
// 	// if err != nil {
// 	// 	return
// 	// }

// 	if len(res.ID) == 0 {
// 		err = errors.New("username or password is not correct 4")
// 		return
// 	}

// 	// set credentials
// 	jwtToken, err = jwtService.GenerateToken(res.ID, res.Username, "")
// 	if err != nil {
// 		log.Println("here 1")
// 		return
// 	}

// 	cfg := singleton.GetConfig().Cookie

// 	// set cookie
// 	ctx.SetCookie(cfg.CookieName, jwtToken, cfg.MaxAge, "/", cfg.CookieDomain, cfg.CookieSecure, cfg.CookieHTTPOnly)
// }

// func (u *UserHandler) GetMyself(ctx *gin.Context) {
// 	var (
// 		req presenter.User
// 		dto usecase_dto.User
// 		res usecase_dto.User
// 		err error
// 	)

// 	if err = ctx.ShouldBindJSON(&req); err != nil {
// 		return
// 	}

// 	validate := validator.New()
// 	if err = validate.Struct(req); err != nil {
// 		return
// 	}

// 	copier.MustCopy(&dto, &req)
// 	// res, err = u.UserUsecase.CreateUser(ctx, dto)
// 	// if err != nil {
// 	// 	return
// 	// }

// 	log.Println(res)
// }
