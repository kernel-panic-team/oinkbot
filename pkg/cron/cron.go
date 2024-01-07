package cron

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/kernel-panic-team/oinkbot/internal/config"
)

type Cron struct {
	cfg       *config.Config
	scheduler *gocron.Scheduler
}

func New(cfg *config.Config) *Cron {
	s := gocron.NewScheduler(time.UTC)
	s = s.Every(time.Duration(cfg.CronInterval))
	return &Cron{scheduler: s}
}

func (c *Cron) Start(f func()) error {
	_, err := c.scheduler.Do(f)
	if err != nil {
		return err
	}
	c.scheduler.StartAsync()
	return nil
}

func (c *Cron) Stop() {
	c.scheduler.Stop()
}
