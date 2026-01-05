package external

// func Routing() *gin.Engine {
// 	gin.SetMode(gin.ReleaseMode)

// 	r := gin.New()

// 	r.Use(gin.Recovery(), gin.Logger())

// 	// middleware

// 	r.Use(middleware.CorsMiddleware())
// 	r.Use(middleware.JSONWriterMiddleware)

// 	// userHandler, err := registry.InitializeUserService(context.Background())
// 	userHandler := handler.NewUserRestfulHandler()

// 	// go check for health
// 	r.GET("/api/healthz")

// 	r.POST("/api_us/register", userHandler.CreateUserProfile)
// 	// retrieve the id of the user, require tokens.

// 	// login for all users
// 	// set cookie with JWT token.
// 	r.POST("/api_us/login", userHandler.Login)
// 	// create new account (does not duplicated phone)
// 	r.GET("/api_us/user/me")

// 	return r
// }
