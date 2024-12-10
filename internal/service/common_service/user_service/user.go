package user_service

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/gen"
	"crmeb_go/internal/repository/user_repository"
	"crmeb_go/internal/validation"
	"crmeb_go/pkg/cache"
	"crmeb_go/pkg/jwt"
	"crmeb_go/pkg/logs"
	"github.com/go-redsync/redsync/v4"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"golang.org/x/crypto/bcrypt"
	"time"
)

func NewUserService(
	cache *cache.Cache,
	rLock *redsync.Redsync,
	tm repository.Transaction,
	userRepo user_repository.UserRepository,
	jwt *jwt.JWT,
) UserService {
	return &userService{
		userRepo: userRepo,
		tx:       tm,
		jwt:      jwt,
	}
}

type userService struct {
	userRepo user_repository.UserRepository
	tx       repository.Transaction
	jwt      *jwt.JWT
}

func (s *userService) Register(ctx context.Context, req *validation.Register) error {
	// check username
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		logs.Log.Error("get user by email error:", zap.Error(err))
		return err
	}
	if err == nil && user != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Generate user ID
	userId := uuid.NewString()
	user = &model.User{
		UserID:   userId,
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	// Transaction demo
	err = s.tx.Transaction(ctx, func(query *gen.Query) error {
		// Create a user
		if err = s.userRepo.CreateTx(ctx, query, user); err != nil {
			return err
		}
		// TODO: other repo
		return nil
	})
	return err
}

func (s *userService) Login(ctx context.Context, req *validation.Login) (string, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil || user == nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}
	token, err := s.jwt.GenToken(user.UserID, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) FindUser(ctx context.Context, req *validation.FindUser) (*model.User, error) {
	user, err := s.userRepo.GetUserByCondition(ctx, user_data.Condition{
		Email:    req.Email,
		Nickname: req.Nickname,
		UserID:   req.UserID,
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
