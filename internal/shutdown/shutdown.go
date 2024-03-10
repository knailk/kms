package shutdown

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/gommon/log"
)

var (
	ErrAbortedAsGotStopSignal = errors.New("aborted as got stop signal")
)

var defaultStopSigs = []os.Signal{syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM}

type Tasks struct {
	sigChan chan os.Signal
	tasks   []Task
}

type Task interface {
	Shutdown(ctx context.Context) error
	Name() string
}

func NewShutdownTasks() *Tasks {
	return &Tasks{
		tasks:   make([]Task, 0),
		sigChan: make(chan os.Signal, 1),
	}
}

func (t *Tasks) Add(tasks ...Task) {
	t.tasks = append(t.tasks, tasks...)
}

func (t *Tasks) ExecuteAll(ctx context.Context) {
	for i := len(t.tasks) - 1; i >= 0; i-- {
		task := t.tasks[i]
		if task == nil {
			continue
		}

		log.Infof("Shutting down %s...", task.Name())
		if err := task.Shutdown(ctx); err != nil {
			log.Error(err, "Failed to shutdown %s", task.Name())
		}
		t.tasks[i] = nil
	}
}

func (t *Tasks) WaitForServerStop(ctx context.Context) {
	signal.Notify(t.sigChan, defaultStopSigs...)
	sig := <-t.sigChan

	log.Infof("got stop sig: %s", sig.String())
}

func (t *Tasks) GetSigChan() chan os.Signal {
	return t.sigChan
}

func (t *Tasks) HasStopSignal() bool {
	return len(t.sigChan) > 0
}
