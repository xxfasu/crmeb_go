package constants

// 登录相关配置键
const (
	ConfigKeyAdminLoginLogoLeftTop     = "site_logo_lefttop"  // 当前代码版本
	ConfigKeyAdminLoginLogoLogin       = "site_logo_login"    // 登录页LOGO
	ConfigKeyAdminLoginBackgroundImage = "admin_login_bg_pic" // 登录页背景图
)

// 微信分享相关配置键
const (
	ConfigKeyAdminWeChatShareImage    = "wechat_share_img"      // 微信分享图片（公众号）
	ConfigKeyAdminWeChatShareTitle    = "wechat_share_title"    // 微信分享标题（公众号）
	ConfigKeyAdminWeChatShareSynopsis = "wechat_share_synopsis" // 微信分享简介（公众号）
)

// 分销相关配置键
const (
	ConfigKeyBrokerageFuncStatus    = "brokerage_func_status"     // 是否启用分销
	ConfigKeyStoreBrokerageStatus   = "store_brokerage_status"    // 分销模式 :1-指定分销，2-人人分销
	StoreBrokerageStatusAppoint     = "1"                         // 分销模式 :1-指定分销
	StoreBrokerageStatusPeople      = "2"                         // 分销模式 :2-人人分销
	ConfigKeyStoreBrokerageRatio    = "store_brokerage_ratio"     // 一级返佣比例
	ConfigKeyStoreBrokerageTwo      = "store_brokerage_two"       // 二级返佣比例
	ConfigKeyStoreBrokerageIsBubble = "store_brokerage_is_bubble" // 判断是否开启气泡
	ConfigKeyStoreBrokerageQuota    = "store_brokerage_quota"     // 判断是否分销消费门槛
)

// 会员和充值功能配置键
const (
	ConfigKeyVipOpen          = "vip_open"           // 是否开启会员功能
	ConfigKeyRechargeSwitch   = "recharge_switch"    // 是否开启充值功能
	ConfigKeyStoreSelfMention = "store_self_mention" // 是否开启门店自提
)

// 腾讯地图配置键
const (
	ConfigSiteTengXunMapKey = "tengxun_map_key" // 腾讯地图key
)

// 退款和提现配置键
const (
	ConfigKeyStorReason       = "stor_reason"            // 退款理由
	ConfigExtractMinPrice     = "user_extract_min_price" // 提现最低金额
	ConfigExtractFreezingTime = "extract_time"           // 提现冻结时间
)

// 全场包邮设置
const (
	StoreFeePostageSwitch = "store_free_postage_switch" // 全场满额包邮开关
	StoreFeePostage       = "store_free_postage"        // 全场满额包邮金额
)

// 积分设置配置键
const (
	ConfigKeyIntegralRate          = "integral_ratio"      // 积分抵用比例(1积分抵多少金额)
	ConfigKeyIntegralRateOrderGive = "order_give_integral" // 下单支付金额按比例赠送积分（实际支付1元赠送多少积分)
)

// 支付设置配置键
const (
	ConfigPayWeixinOpen = "pay_weixin_open" // 微信支付开关
	ConfigYuePayStatus  = "yue_pay_status"  // 余额支付状态
	ConfigAliPayStatus  = "ali_pay_status"  // 支付宝支付状态
)
