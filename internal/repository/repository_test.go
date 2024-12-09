package repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository/gen"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"os"
	"sync"
	"testing"
	"time"
)

var DB *gorm.DB

func TestMain(m *testing.M) {
	const MySQLDSN = "root:password@tcp(127.0.0.1:3306)/template?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	DB = db
	gen.SetDefault(db)
	code := m.Run()
	fmt.Println("test end")
	os.Exit(code)
}

func TestUserRepository_Create(t *testing.T) {
	user := &model.User{
		Email:    "test@test.com",
		Nickname: "test",
		Password: "test",
		UserID:   "test",
	}
	err := gen.User.WithContext(context.Background()).Create(user)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_Delete(t *testing.T) {

	_, err := gen.User.WithContext(context.Background()).Where(gen.User.UserID.Eq("test")).Delete()
	if err != nil {
		t.Error(err)
	}
}

func AA() []field.Expr {
	return []field.Expr{
		gen.User.UserID,
		gen.User.Nickname,
		gen.User.Email,
		gen.User.Password,
	}
}

func TestUserRepository_Select(t *testing.T) {

	userList, err := gen.User.WithContext(context.Background()).Where(gen.User.Nickname.Eq("test")).Select(AA()...).Find()
	if err != nil {
		t.Error(err)
	}
	t.Log(userList)
}

func TestUserRepository_Update(t *testing.T) {
	user := &model.User{
		ID:       8,
		Email:    "test@test.com",
		Nickname: "test8",
		Password: "test8",
		UserID:   "test8",
	}
	err := gen.User.WithContext(context.Background()).Save(user)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_Transaction(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	err := TransactionTest(ctx, func(query *gen.Query) error {
		query.User.WithContext(ctx).Create(&model.User{
			Email:    "test@test.com",
			Nickname: "test",
			Password: "test",
			UserID:   "test1",
		})

		query.User.WithContext(ctx).Create(&model.User{
			Email:    "test@test.com",
			Nickname: "test",
			Password: "test",
			UserID:   "test2",
		})
		time.Sleep(5 * time.Second)
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

func TransactionTest(ctx context.Context, fn func(query *gen.Query) error) error {
	return DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var flag error
		var wg sync.WaitGroup
		done := make(chan struct{})
		wg.Add(1)
		go func() {
			defer wg.Done()
			query := gen.Use(tx)
			flag = fn(query)
		}()
		go func() {
			defer close(done)
			wg.Wait()
		}()
		select {
		case <-ctx.Done():
			fmt.Println("Transaction Rollback due to timeout:", ctx.Err())
			return ctx.Err()
		case <-done:
			return flag
		}
	})
}
