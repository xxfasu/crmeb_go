// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/plugin/soft_delete"
)

const TableNameStoreBargain = "eb_store_bargain"

// StoreBargain 砍价表
type StoreBargain struct {
	ID              int64                 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true;comment:砍价商品ID" json:"id"`                              // 砍价商品ID
	ProductID       int64                 `gorm:"column:product_id;type:int unsigned;not null;comment:关联商品ID" json:"product_id"`                                   // 关联商品ID
	Title           string                `gorm:"column:title;type:varchar(255);not null;comment:砍价活动名称" json:"title"`                                             // 砍价活动名称
	Image           string                `gorm:"column:image;type:varchar(150);not null;comment:砍价活动图片" json:"image"`                                             // 砍价活动图片
	UnitName        string                `gorm:"column:unit_name;type:varchar(16);comment:单位名称" json:"unit_name"`                                                 // 单位名称
	Stock           int64                 `gorm:"column:stock;type:int unsigned;comment:库存" json:"stock"`                                                          // 库存
	Sales           int64                 `gorm:"column:sales;type:int unsigned;comment:销量" json:"sales"`                                                          // 销量
	Images          string                `gorm:"column:images;type:varchar(2000);not null;comment:砍价商品轮播图" json:"images"`                                         // 砍价商品轮播图
	StartTime       int64                 `gorm:"column:start_time;type:bigint unsigned;not null;comment:砍价开启时间" json:"start_time"`                                // 砍价开启时间
	StopTime        int64                 `gorm:"column:stop_time;type:bigint unsigned;not null;comment:砍价结束时间" json:"stop_time"`                                  // 砍价结束时间
	StoreName       string                `gorm:"column:store_name;type:varchar(255);comment:砍价商品名称" json:"store_name"`                                            // 砍价商品名称
	Price           decimal.Decimal       `gorm:"column:price;type:decimal(8,2) unsigned;comment:砍价金额" json:"price"`                                               // 砍价金额
	MinPrice        decimal.Decimal       `gorm:"column:min_price;type:decimal(8,2) unsigned;comment:砍价商品最低价" json:"min_price"`                                    // 砍价商品最低价
	Num             int64                 `gorm:"column:num;type:int unsigned;comment:购买数量限制——单个活动每个用户发起砍价次数限制" json:"num"`                                        // 购买数量限制——单个活动每个用户发起砍价次数限制
	BargainMaxPrice decimal.Decimal       `gorm:"column:bargain_max_price;type:decimal(8,2) unsigned;comment:用户每次砍价的最大金额" json:"bargain_max_price"`                // 用户每次砍价的最大金额
	BargainMinPrice decimal.Decimal       `gorm:"column:bargain_min_price;type:decimal(8,2) unsigned;comment:用户每次砍价的最小金额" json:"bargain_min_price"`                // 用户每次砍价的最小金额
	BargainNum      int64                 `gorm:"column:bargain_num;type:int unsigned;not null;default:1;comment:帮砍次数——单个商品用户可以帮砍的次数" json:"bargain_num"`          // 帮砍次数——单个商品用户可以帮砍的次数
	Status          int32                 `gorm:"column:status;type:tinyint unsigned;not null;default:1;comment:砍价状态 0(到砍价时间不自动开启)  1(到砍价时间自动开启时间)" json:"status"` // 砍价状态 0(到砍价时间不自动开启)  1(到砍价时间自动开启时间)
	GiveIntegral    int64                 `gorm:"column:give_integral;type:int;comment:反多少积分" json:"give_integral"`                                                // 反多少积分
	Info            string                `gorm:"column:info;type:varchar(255);comment:砍价活动简介" json:"info"`                                                        // 砍价活动简介
	Cost            decimal.Decimal       `gorm:"column:cost;type:decimal(8,2) unsigned;comment:成本价" json:"cost"`                                                  // 成本价
	Sort            int64                 `gorm:"column:sort;type:int unsigned;not null;comment:排序" json:"sort"`                                                   // 排序
	IsHot           int32                 `gorm:"column:is_hot;type:tinyint unsigned;not null;comment:是否推荐0不推荐1推荐" json:"is_hot"`                                  // 是否推荐0不推荐1推荐
	IsDel           int32                 `gorm:"column:is_del;type:tinyint unsigned;not null;comment:是否删除 0未删除 1删除" json:"is_del"`                                // 是否删除 0未删除 1删除
	AddTime         int64                 `gorm:"column:add_time;type:bigint unsigned;comment:添加时间" json:"add_time"`                                               // 添加时间
	IsPostage       int32                 `gorm:"column:is_postage;type:tinyint unsigned;not null;default:1;comment:是否包邮 0不包邮 1包邮" json:"is_postage"`              // 是否包邮 0不包邮 1包邮
	Postage         decimal.Decimal       `gorm:"column:postage;type:decimal(10,2) unsigned;comment:邮费" json:"postage"`                                            // 邮费
	Rule            string                `gorm:"column:rule;type:text;comment:砍价规则" json:"rule"`                                                                  // 砍价规则
	Look            int64                 `gorm:"column:look;type:int unsigned;comment:砍价商品浏览量" json:"look"`                                                       // 砍价商品浏览量
	Share           int64                 `gorm:"column:share;type:int unsigned;comment:砍价商品分享量" json:"share"`                                                     // 砍价商品分享量
	TempID          int64                 `gorm:"column:temp_id;type:int;comment:运费模板ID" json:"temp_id"`                                                           // 运费模板ID
	Weight          decimal.Decimal       `gorm:"column:weight;type:decimal(8,2);default:0.00;comment:重量" json:"weight"`                                           // 重量
	Volume          decimal.Decimal       `gorm:"column:volume;type:decimal(8,2);default:0.00;comment:体积" json:"volume"`                                           // 体积
	Quota           int64                 `gorm:"column:quota;type:int;not null;comment:限购总数" json:"quota"`                                                        // 限购总数
	QuotaShow       int64                 `gorm:"column:quota_show;type:int;not null;comment:限量总数显示" json:"quota_show"`                                            // 限量总数显示
	PeopleNum       int64                 `gorm:"column:people_num;type:int;comment:砍价人数——需要多少人砍价成功" json:"people_num"`                                            // 砍价人数——需要多少人砍价成功
	CreatedAt       int64                 `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt       int64                 `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	DeletedAt       soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint" json:"-"`
}

// TableName StoreBargain's table name
func (*StoreBargain) TableName() string {
	return TableNameStoreBargain
}
