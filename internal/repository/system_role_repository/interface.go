package system_role_repository

type Reader interface {
}

type Writer interface {
}

type Repository interface {
	Reader
	Writer
}
