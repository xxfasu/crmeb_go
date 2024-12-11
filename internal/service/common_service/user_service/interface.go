package user_service

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/validation"
)

type Service interface {
	Register(ctx context.Context, req *validation.Register) error
	Login(ctx context.Context, req *validation.Login) (string, error)
	FindUser(ctx context.Context, req *validation.FindUser) (*model.User, error)
}
