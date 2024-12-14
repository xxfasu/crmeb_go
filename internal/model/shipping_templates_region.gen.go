// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/plugin/soft_delete"
)

const TableNameShippingTemplatesRegion = "eb_shipping_templates_region"

// ShippingTemplatesRegion 运费模板指定区域费用
type ShippingTemplatesRegion struct {
	ID           int64                 `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:编号" json:"id"`                           // 编号
	TempID       int64                 `gorm:"column:temp_id;type:int;not null;comment:模板ID" json:"temp_id"`                                    // 模板ID
	CityID       int64                 `gorm:"column:city_id;type:int;not null;comment:城市ID" json:"city_id"`                                    // 城市ID
	Title        string                `gorm:"column:title;type:text;comment:描述" json:"title"`                                                  // 描述
	First        decimal.Decimal       `gorm:"column:first;type:decimal(10,2);not null;default:0.00;comment:首件" json:"first"`                   // 首件
	FirstPrice   decimal.Decimal       `gorm:"column:first_price;type:decimal(10,2);not null;default:0.00;comment:首件运费" json:"first_price"`     // 首件运费
	Renewal      decimal.Decimal       `gorm:"column:renewal;type:decimal(10,2);not null;default:0.00;comment:续件" json:"renewal"`               // 续件
	RenewalPrice decimal.Decimal       `gorm:"column:renewal_price;type:decimal(10,2);not null;default:0.00;comment:续件运费" json:"renewal_price"` // 续件运费
	Type         bool                  `gorm:"column:type;type:tinyint(1);not null;default:1;comment:计费方式 1按件数 2按重量 3按体积" json:"type"`          // 计费方式 1按件数 2按重量 3按体积
	Uniqid       string                `gorm:"column:uniqid;type:varchar(32);not null;comment:分组唯一值" json:"uniqid"`                             // 分组唯一值
	Status       bool                  `gorm:"column:status;type:tinyint(1);comment:是否无效" json:"status"`                                        // 是否无效
	CreatedAt    int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt    int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt    soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName ShippingTemplatesRegion's table name
func (*ShippingTemplatesRegion) TableName() string {
	return TableNameShippingTemplatesRegion
}
