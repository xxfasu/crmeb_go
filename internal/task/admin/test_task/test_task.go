package test_task

import (
	"crmeb_go/internal/cron"
	"crmeb_go/pkg/logs"
)

type task struct {
}

func Start(timer cron.Timer) error {
	t := &task{}
	_, err := timer.AddTaskByJobWithSeconds("admin", "*/1 * * * * *", t, "testTask")
	return err
}

func (t *task) Run() {
	logs.Log.Info("testTask")
}
