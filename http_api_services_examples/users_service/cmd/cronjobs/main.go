package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/robfig/cron/v3"

	config "users-service/configs"
	"users-service/internal/cronjob"
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
	userSvc := service.NewUserService(userRepo)

	cleanup := cronjob.NewCleanupJob(userSvc)

	c := cron.New()

	c.AddFunc("@every 30s", cleanup.Run)
	c.Start()

	log.Println("Cron scheduler started.")
	// Block forever. Bad for real applications, need to handle graceful shutdown.
	select {}
}
