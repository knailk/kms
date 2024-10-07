package cron

import "context"

type IUseCase interface {
	ScheduleEveryWeek(ctx context.Context) error
	ManageClassStatus(ctx context.Context) error
}
