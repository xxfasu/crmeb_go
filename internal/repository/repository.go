package repository

import (
	"crmeb_go/internal/repository/system_admin_repository"
	"crmeb_go/internal/repository/system_config_repository"
	"crmeb_go/internal/repository/system_group_data_repository"
	"crmeb_go/internal/repository/system_menu_repository"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	InitDB,
	NewTransaction,
	system_admin_repository.New,
	system_config_repository.New,
	system_group_data_repository.New,
	system_menu_repository.New,
)
