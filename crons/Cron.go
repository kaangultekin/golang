package crons

import (
	"github.com/robfig/cron/v3"
)

func Cron() {
	c := cron.New()

	c.AddFunc("*/1 * * * *", UsersToElasticcearchCron)

	c.Start()
}
