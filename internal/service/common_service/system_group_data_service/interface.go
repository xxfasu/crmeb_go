package system_group_data_service

type Service interface {
	GetListByGid(gid int) ([]any, error)
}
