package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sunthree74/shopping_test/config"
	"github.com/sunthree74/shopping_test/handler"
	"github.com/sunthree74/shopping_test/repository"
	"github.com/sunthree74/shopping_test/usecase"
	"log"
	"net/http"
	"time"
)

func main() {
	db := config.Connect()
	gin.SetMode("debug")
	db = db.Debug()
	gin.DisableConsoleColor()
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Methods",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Authorization",
			"Cookie",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(gin.Recovery())

	userRepo := repository.InitializeUser(db, http.DefaultClient)
	productCategoryRepo := repository.InitializeProductCategory(db, http.DefaultClient)

	userUsecase := usecase.InitializeUser(userRepo)
	productCategoryUsecase := usecase.InitializeProductCategory(productCategoryRepo)

	productCategoryHandler := handler.HandleProductCategory(productCategoryUsecase, userUsecase)
	//authHandler := handler.HandleAuth(userUsecase)
	middleware, err := handler.InitializeMiddleware(userUsecase)
	if err != nil {
		log.Fatalln(err)
	}
	userAuth := middleware.UserAuth()

	router.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte("simple shopping service"))
		return
	})

	router.POST("/auth/:method", userAuth.LoginHandler)

	productCategoryRoute := router.Group("/category")
	{
		productCategoryRoute.Use(userAuth.MiddlewareFunc())
		{
			productCategoryRoute.GET("/list", productCategoryHandler.GetList())
			productCategoryRoute.POST("/create", productCategoryHandler.Create())
			productCategoryRoute.GET("/find/:id", productCategoryHandler.GetById())
		}
	}

	log.Fatalln(router.Run(":8008"))
}
