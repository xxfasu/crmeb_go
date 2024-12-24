package user_visit_record_repository

import (
	"context"
)

type Reader interface {
	FindPageViewsByDate(ctx context.Context, start int64, end int64) (data int64, err error)
}

type Writer interface {
}

type Repository interface {
	Reader
	Writer
}
