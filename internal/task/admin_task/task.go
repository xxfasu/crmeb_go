package admin_task

import (
	"crmeb_go/internal/task/admin_task/test_task"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(test_task.Start)
