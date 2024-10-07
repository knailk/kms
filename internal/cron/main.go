package cron

import (
	"context"
	"kms/app/cron"
	"kms/app/registry"
	"kms/internal/shutdown"
	"time"

	"github.com/go-co-op/gocron"
)

type CronTask struct {
	scheduler *gocron.Scheduler
}

func StartCron(ctx context.Context, provider *registry.Provider, tasks *shutdown.Tasks) error {
	scheduler := gocron.NewScheduler(time.UTC)

	cronTask := &CronTask{scheduler: scheduler}
	tasks.Add(cronTask)

	cronService := cron.NewCronService(
		scheduler,
		registry.InjectedCronUseCase(ctx, provider),
	)

	cronService.Scheduler(ctx)

	cronService.Start()

	return nil
}

func (ct *CronTask) Shutdown(ctx context.Context) error {
	ct.scheduler.Stop()
	return nil
}

func (ct *CronTask) Name() string {
	return "Cron Scheduler"
}
