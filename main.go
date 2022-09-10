package main

import (
	"crowdfund-go/auth"
	"crowdfund-go/config"
	"crowdfund-go/graph"
	"crowdfund-go/graph/generated"
	"log"
	"os"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_authService "crowdfund-go/auth/service"
	_campaignHandler "crowdfund-go/campaign/handler"
	_campaignRepo "crowdfund-go/campaign/repository"
	_campaignService "crowdfund-go/campaign/service"
	_userHandler "crowdfund-go/user/handler"
	_userRepo "crowdfund-go/user/repository"
	_userService "crowdfund-go/user/service"
)

func init() {
	config.SetConfiguration()
}

func main() {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	userRepository := _userRepo.NewRepository(db)
	campaignRepository := _campaignRepo.NewRepository(db)

	userService := _userService.NewService(userRepository)
	authService := _authService.NewService()
	campaignService := _campaignService.NewService(campaignRepository)

	userHandler := _userHandler.NewUserHandler(userService, authService)
	campaignHandler := _campaignHandler.NewCampaignHandler(campaignService)

	router := gin.Default()

	//REST
	router.Static("/images", "./images")
	api := router.Group("api/v1")

	api.POST("/users", userHandler.Register)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", auth.Middleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/campaigns", campaignHandler.Campaigns)
	api.POST("/campaigns", auth.Middleware(authService, userService), campaignHandler.Create)

	//GraphQL
	router.POST("/query", auth.Middleware(authService, userService), graphqlHandler(db))
	router.GET("/", playgroundHandler())

	router.Run()
}

func graphqlHandler(db *gorm.DB) gin.HandlerFunc {
	h := gqlHandler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
