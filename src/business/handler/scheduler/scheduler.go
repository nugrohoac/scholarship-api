package scheduler

import (
	"github.com/Nusantara-Muda/scholarship-api/src/business"
	"github.com/go-co-op/gocron"
	"github.com/labstack/echo"
)

type sch struct {
	scheduler      *gocron.Scheduler
	log            echo.Logger
	scholarshipSvc business.ScholarshipService
}

type Scheduler interface {
	ScholarshipStatusChecker(d int) error
}

func Init(s *gocron.Scheduler, log echo.Logger, scholarshipSvc business.ScholarshipService) Scheduler {
	return &sch{
		scheduler:      s,
		log:            log,
		scholarshipSvc: scholarshipSvc,
	}
}
