package repository

import (
	"context"
	"crmeb_go/internal/repository/gen"
	"crmeb_go/internal/repository/system_admin_repository"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

func setupRepository() (system_admin_repository.Repository, sqlmock.Sqlmock, error) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("failed to create sqlmock: %v", err))
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("failed to open gorm connection: %v", err))
	}

	// rdb, _ := redismock.NewClientMock()
	db = db.Debug()
	gen.SetDefault(db)
	userRepo := system_admin_repository.New(db)

	return userRepo, mock, nil
}

func TestSystemAdminRepository_GetUser(t *testing.T) {
	Convey("SystemAdminRepository GetUser方法测试", t, func() {
		userRepo, mock, err := setupRepository()
		So(err, ShouldBeNil)
		So(userRepo, ShouldNotBeNil)
		ctx := context.Background()

		// 设置 mock 行为
		mock.ExpectQuery(regexp.QuoteMeta(
			"SELECT * FROM `eb_system_admin` WHERE `eb_system_admin`.`account` = ? AND `eb_system_admin`.`deleted_at` = ? ORDER BY `eb_system_admin`.`id` LIMIT ?",
		)).
			WithArgs("admin", 0, 1).
			WillReturnRows(sqlmock.NewRows([]string{
				"id", "account", "pwd", "real_name", "roles", "last_ip",
				"login_count", "level", "status", "phone", "is_sms",
				"created_at", "updated_at", "deleted_at",
			}).
				AddRow(1, "admin", "hashed_password", "管理员", "1", "127.0.0.1",
					10, 1, 1, "12345678901", 1,
					1617181723, 1617181723, 0))
		systemAdmin, err := userRepo.GetUser(ctx, "admin")

		So(err, ShouldBeNil)
		So(mock.ExpectationsWereMet(), ShouldBeNil)
		So(systemAdmin, ShouldNotBeNil)
	})
}
