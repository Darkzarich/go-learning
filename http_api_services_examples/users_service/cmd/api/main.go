package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	config "users-service/configs"
	"users-service/internal/handler"
	"users-service/internal/repository"
	"users-service/internal/service"
)

func main() {
	cfg := config.Load()

	db, err := sql.Open("sqlite3", cfg.DBPath)
	if err != nil {
		log.Fatalf("cannot open db: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	if err := userRepo.InitDatabase(); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	r := gin.Default()

	g := r.Group("/api")

	userHandler.RegisterRoutes(g)

	if err := r.Run(":3000"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
