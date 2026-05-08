package cronjob

import (
	"log"

	"users-service/internal/service"
)

type CleanupJob struct {
	svc *service.UserService
}

func NewCleanupJob(svc *service.UserService) *CleanupJob {
	return &CleanupJob{svc: svc}
}

func (j *CleanupJob) Run() {
	deleted, err := j.svc.CleanupInactive(30) // 30 days
	if err != nil {
		log.Printf("CleanupJob error: %v\n", err)
		return
	}

	log.Printf("CleanupJob: removed %d inactive users\n", deleted)
}
