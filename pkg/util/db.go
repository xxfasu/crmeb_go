package util

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

// 获取指定表的所有字段
func getColumns(db *sql.DB, tableName string) ([]string, error) {
	// 使用 SHOW COLUMNS FROM 语句获取字段信息
	query := fmt.Sprintf("SHOW COLUMNS FROM %s", tableName)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []string
	for rows.Next() {
		var field, colType, isNull, key, def, extra sql.NullString
		if err := rows.Scan(&field, &colType, &isNull, &key, &def, &extra); err != nil {
			return nil, err
		}
		columns = append(columns, field.String)
	}

	// 如果需要检查 rows.Next() 之后可能的错误，可以放在这里
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return columns, nil
}

// 格式化输出字段的函数
// formatType = 1: 输出 "id , name , email" 这种格式
// formatType = 2: 输出 "id","name","email" 这种格式
// formatType = 3: 两种格式都输出
func formatColumns(columns []string, formatType int) string {
	switch formatType {
	case 1:
		// "id , name , email"
		return strings.Join(columns, " , ")
	case 2:
		// "id","name","email"
		quoted := make([]string, len(columns))
		for i, col := range columns {
			quoted[i] = fmt.Sprintf(`"%s"`, col)
		}
		return strings.Join(quoted, ",")
	case 3:
		// 同时输出两种格式（这里仅作示例，具体按需求组合）
		format1 := strings.Join(columns, " , ")
		quoted := make([]string, len(columns))
		for i, col := range columns {
			quoted[i] = fmt.Sprintf(`"%s"`, col)
		}
		format2 := strings.Join(quoted, ",")
		return fmt.Sprintf("%s\n%s", format1, format2)
	default:
		return ""
	}
}
