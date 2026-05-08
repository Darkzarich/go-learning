package main

import (
	"database/sql"
	"log"
	"net/http"

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

	mux := http.NewServeMux()
	userHandler.RegisterRoutes(mux)

	log.Printf("API server listening on :%s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
