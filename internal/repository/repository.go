package repository

import (
	"context"
	"crmeb_go/internal/conf"
	"crmeb_go/internal/repository/gen"
	"crmeb_go/internal/repository/user_repository"
	"crmeb_go/pkg/logs"
	"crmeb_go/pkg/safe"
	"crmeb_go/pkg/zapgorm2"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var ProviderSet = wire.NewSet(
	InitDB,
	NewTransaction,
	user_repository.NewUserRepository,
)

type transaction struct {
	DB *gorm.DB
	// rdb    *redis.Client
}

func NewTransaction(
	DB *gorm.DB,
) Transaction {
	return &transaction{
		DB: DB,
	}
}

type Transaction interface {
	Transaction(ctx context.Context, fn func(query *gen.Query) error) error
}

func (r *transaction) Transaction(ctx context.Context, fn func(query *gen.Query) error) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var flag error
		var wg sync.WaitGroup
		done := make(chan struct{})
		wg.Add(1)
		safe.Go(func() {
			defer wg.Done()
			query := gen.Use(tx)
			flag = fn(query)
		})
		safe.Go(func() {
			defer close(done)
			wg.Wait()
		})
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-done:
			return flag
		}
	})
}

func InitDB() (*gorm.DB, func(), error) {
	var (
		db  *gorm.DB
		err error
	)

	logger := zapgorm2.New(logs.Log.Logger)
	// GORM doc: https://gorm.io/docs/connecting_to_the_database.html

	db, err = gorm.Open(mysql.Open(conf.Config.Mysql.Source), &gorm.Config{
		Logger: logger,
	})

	if err != nil {
		return nil, nil, err
	}
	db = db.Debug()
	gen.SetDefault(db)
	// Connection Pool config
	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db, func() {
		sqlDB.Close()
	}, nil
}
