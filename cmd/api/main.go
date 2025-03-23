package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "jamlink-backend/docs"
	"jamlink-backend/internal/adapter/http"
	"jamlink-backend/internal/infra/db"
	animalRepository "jamlink-backend/internal/modules/animal/repository"
	animalUsecase "jamlink-backend/internal/modules/animal/usecase"
	userRepository "jamlink-backend/internal/modules/user/repository"
	userUsecase "jamlink-backend/internal/modules/user/usecase"
	"jamlink-backend/internal/shared/security"
)

// @title Jamlink API
// @version 1.0
// @description This is an API with Swagger and Gin.
// @host localhost:8080
// @BasePath /

func main() {
	_ = godotenv.Load()
	database := db.ConnectDB()
	db.MigrateDB(database)

	// Repositories
	animalRepo := animalRepository.NewPostgresAnimalRepository(database)
	userRepo := userRepository.NewPostgresUserRepository(database)

	// Services
	securityService := security.NewSecurityService()

	// Use Cases
	createAnimalUseCase := animalUsecase.NewCreateAnimalUseCase(animalRepo)
	getAnimalListUseCase := animalUsecase.NewGetAnimalListUseCase(animalRepo)
	getAnimalByIdUseCase := animalUsecase.NewGetAnimalByIdUseCase(animalRepo)

	createUserUseCase := userUsecase.NewCreateUserUseCase(userRepo, securityService)
	loginUserUseCase := userUsecase.NewLoginUserUseCase(userRepo, securityService)
	refreshTokenUseCase := userUsecase.NewRefreshTokenUseCase(securityService)

	// Setup router
	r := gin.Default()

	http.NewAnimalHandler(r, createAnimalUseCase, getAnimalListUseCase, getAnimalByIdUseCase, securityService)
	http.NewAuthHandler(r, createUserUseCase, loginUserUseCase, refreshTokenUseCase)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	// Run server
	if err := r.Run(":8080"); err != nil {
		return
	}
}
