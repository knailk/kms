package cron

import (
	"context"
	"kms/app/domain/entity"
	"kms/app/external/persistence/database/repository"
	"kms/pkg/date"
	"kms/pkg/logger"
	"kms/pkg/schedule_weekly"
	"time"

	"github.com/google/uuid"
)

type useCase struct {
	repo *repository.PostgresRepository
}

func NewUseCase(repo *repository.PostgresRepository) IUseCase {
	return &useCase{
		repo: repo,
	}
}

func (uc *useCase) ScheduleEveryWeek(ctx context.Context) error {
	currentDay := int(time.Now().Weekday())
	if currentDay == 0 {
		currentDay = 7 // Sunday
	}

	nextMonday := date.ToDate(time.Now().AddDate(0, 0, -currentDay+8))
	nextSunday := date.ToDate(time.Now().AddDate(0, 0, -currentDay+14))

	c := uc.repo.Class
	s := uc.repo.Schedule

	classes, err := c.
		Where(c.Status.Eq(int(entity.InProgress))).
		Preload(c.Schedules.On(s.Date.Between(nextMonday-1, nextSunday+1))).
		Find()
	if err != nil {
		logger.WithError(err).Error("list class error")
		return err
	}

	newSchedule := make([]*entity.Schedule, 0)

	for _, class := range classes {
		if len(class.Schedules) > 10 {
			continue
		}

		schedulesNextWeek := schedule_weekly.AdjustScheduleForCurrentWeek(time.Now().AddDate(0, 0, 7))

		for i := range schedulesNextWeek {
			schedulesNextWeek[i].ID = uuid.New()
			schedulesNextWeek[i].ClassID = class.ID
		}

		newSchedule = append(newSchedule, schedulesNextWeek...)
	}

	_ = s.CreateInBatches(newSchedule, 100)

	return nil
}

func (uc *useCase) ManageClassStatus(ctx context.Context) error {
	y, m, d := time.Now().Date()
	now := int64(d + 100*int(m) + 10000*y)

	c := uc.repo.Class

	_, _ = uc.repo.Class.
		Where(
			c.FromDate.Gte(now),
			c.ToDate.Lte(now),
			c.Status.Eq(int(entity.Scheduled)),
		).
		Update(c.Status, entity.InProgress)
	_, _ = uc.repo.Class.
		Where(
			c.ToDate.Gt(now),
			c.Status.Eq(int(entity.InProgress)),
		).
		Update(c.Status, entity.Completed)

	return nil
}
