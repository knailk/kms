package cron

import (
	"context"
	"kms/pkg/logger"

	"kms/app/usecase/cron"

	"github.com/go-co-op/gocron"
)

type cronService struct {
	scheduler *gocron.Scheduler
	cronUc    cron.IUseCase
}

func NewCronService(
	scheduler *gocron.Scheduler,
	cronUc cron.IUseCase,
) *cronService {
	return &cronService{
		scheduler: scheduler,
		cronUc:    cronUc,
	}
}

func (s *cronService) Scheduler(ctx context.Context) error {
	logger.Info("Start schedule for class")

	_, err := s.scheduler.Every(3).Day().Do(func() {
		s.cronUc.ScheduleEveryWeek(ctx)
	})
	if err != nil {
		logger.WithError(err).Error("Failed to schedule cron job")
		return err
	}

	_, err = s.scheduler.Every(1).Day().Do(func() {
		s.cronUc.ManageClassStatus(ctx)
	})
	if err != nil {
		logger.WithError(err).Error("Failed to schedule cron job")
		return err
	}

	return nil
}

func (s *cronService) Start() {
	s.scheduler.StartAsync()
}
