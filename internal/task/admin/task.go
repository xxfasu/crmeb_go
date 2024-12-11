package admin

import (
	"crmeb_go/internal/task/admin/test_task"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(test_task.Start)
