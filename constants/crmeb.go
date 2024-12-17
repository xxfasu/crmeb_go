package constants

// 数值常量
const (
	// TOKEN_EXPRESS_MINUTES 3小时
	TokenExpressMinutes = 1440 // 60 * 24 分钟

	// HTTPSTATUS_CODE_SUCCESS
	HTTPStatusCodeSuccess = 200

	// NUM_ZERO
	NumZero = 0

	// NUM_ONE
	NumOne = 1

	// NUM_TWO
	NumTwo = 2

	// NUM_THREE
	NumThree = 3

	// NUM_FIVE
	NumFive = 5

	// NUM_SEVEN
	NumSeven = 7

	// NUM_TEN
	NumTen = 10

	// NUM_ONE_HUNDRED
	NumOneHundred = 100

	// HEADER_AUTHORIZATION_KEY
	HeaderAuthorizationKey = "Authori-zation"

	// VALIDATE_REDIS_KEY_PREFIX
	ValidateRedisKeyPrefix = "validate_code_"

	// DATE_FORMAT
	DateFormat = "yyyy-MM-dd HH:mm:ss"

	// DATE_FORMAT_UTC
	DateFormatUTC = "yyyy-MM-dd'T'HH:mm:ss.SSS'Z'"

	// DATE_FORMAT_DATE
	DateFormatDate = "yyyy-MM-dd"

	// DATE_FORMAT_YEAR
	DateFormatYear = "yyyy"

	// DATE_FORMAT_MONTH_DATE
	DateFormatMonthDate = "MM-dd"

	// DATE_FORMAT_MONTH
	DateFormatMonth = "yyyy-MM"

	// DATE_TIME_FORMAT_NUM
	DateTimeFormatNum = "yyyyMMddHHmmss"

	// DATE_FORMAT_NUM
	DateFormatNum = "yyyyMMdd"

	// DATE_FORMAT_START
	DateFormatStart = "yyyy-MM-dd 00:00:00"

	// DATE_FORMAT_END
	DateFormatEnd = "yyyy-MM-dd 23:59:59"

	// DATE_FORMAT_MONTH_START
	DateFormatMonthStart = "yyyy-MM-01 00:00:00"

	// DATE_FORMAT_YEAR_START
	DateFormatYearStart = "yyyy-01-01 00:00:00"

	// DATE_FORMAT_YEAR_END
	DateFormatYearEnd = "yyyy-12-31 23:59:59"

	// DATE_FORMAT_HHMMSS
	DateFormatHHMMSS = "HH:mm:ss"

	// DATE_FORMAT_START_PEREND
	DateFormatStartPerEnd = "00:00:00"

	// DATE_FORMAT_END_PEREND
	DateFormatEndPerEnd = "23:59:59"

	// DATE_FORMAT_HHMM
	DateFormatHHMM = "yyyy-MM-dd HH:mm"

	// USER_BILL_OPERATE_LOG_TITLE
	UserBillOperateLogTitle = "{$title}{$operate}了{$value}{$founds}"

	// USER_LEVEL_OPERATE_LOG_MARK
	UserLevelOperateLogMark = "尊敬的用户 【{$userName}】, 在{$date}赠送会员等级成为{$levelName}会员"

	// USER_LEVEL_UP_LOG_MARK
	UserLevelUpLogMark = "尊敬的用户 【{$userName}】, 在{$date}您升级为为{$levelName}会员"

	// USER_LOGIN_PASSWORD_MD5_KEYWORDS
	UserLoginPasswordMD5Keywords = "crmeb"

	// USER_TOKEN_REDIS_KEY_PREFIX
	UserTokenRedisKeyPrefix = "TOKEN_USER:"

	// USER_LOGIN_TYPE_H5
	UserLoginTypeH5 = "h5"

	// USER_LOGIN_TYPE_PUBLIC
	UserLoginTypePublic = "wechat"

	// USER_LOGIN_TYPE_PROGRAM
	UserLoginTypeProgram = "routine"

	// USER_LOGIN_TYPE_IOS_WX
	UserLoginTypeIOSWX = "iosWx"

	// USER_LOGIN_TYPE_ANDROID_WX
	UserLoginTypeAndroidWX = "androidWx"

	// USER_LOGIN_TYPE_IOS
	UserLoginTypeIOS = "ios"

	// USER_DEFAULT_AVATAR_CONFIG_KEY
	UserDefaultAvatarConfigKey = "h5_avatar"

	// USER_DEFAULT_SPREAD_ID
	UserDefaultSpreadID = 0

	// DEFAULT_PAGE
	DefaultPage = 1

	// DEFAULT_LIMIT
	DefaultLimit = 20

	// SORT_ASC
	SortAsc = "asc"

	// SORT_DESC
	SortDesc = "desc"

	// EXPORT_MAX_LIMIT
	ExportMaxLimit = 99999

	// PRODUCT_SELECT_CATEGORY_NUM_MAX
	ProductSelectCategoryNumMax = 10

	// UPLOAD_TYPE_IMAGE
	UploadTypeImage = "crmebimage"

	// UPLOAD_TYPE_FILE
	UploadTypeFile = "file"

	// UPLOAD_MODEL_PATH_EXCEL
	UploadModelPathExcel = "excel"

	// UPLOAD_IMAGE_EXT_STR_CONFIG_KEY
	UploadImageExtStrConfigKey = "image_ext_str"

	// UPLOAD_IMAGE_MAX_SIZE_CONFIG_KEY
	UploadImageMaxSizeConfigKey = "image_max_size"

	// UPLOAD_FILE_EXT_STR_CONFIG_KEY
	UploadFileExtStrConfigKey = "file_ext_str"

	// UPLOAD_FILE_MAX_SIZE_CONFIG_KEY
	UploadFileMaxSizeConfigKey = "file_max_size"

	// ARTICLE_BANNER_LIMIT
	ArticleBannerLimit = "news_slides_limit"

	// CITY_LIST
	CityList = "city_list"

	// CITY_LIST_TREE
	CityListTree = "city_list_tree"

	// CITY_LIST_LEVEL_1
	CityListLevel1 = "city_list_level_1"

	// PRODUCT_STOCK_UPDATE
	ProductStockUpdate = "product_stock_update"

	// PRODUCT_SECKILL_STOCK_UPDATE
	ProductSeckillStockUpdate = "product_seckill_stock_update"

	// PRODUCT_BARGAIN_STOCK_UPDATE
	ProductBargainStockUpdate = "product_bargain_stock_update"

	// PRODUCT_COMBINATION_STOCK_UPDATE
	ProductCombinationStockUpdate = "product_combination_stock_update"

	// PRODUCT_STOCK_LIST
	ProductStockList = "product_stock_list"

	// WE_CHAT_MESSAGE_KEY_PUBLIC
	WeChatMessageKeyPublic = "we_chat_public_message_list"

	// WE_CHAT_MESSAGE_KEY_PROGRAM
	WeChatMessageKeyProgram = "we_chat_program_message_list"

	// WE_CHAT_MESSAGE_INDUSTRY_KEY
	WeChatMessageIndustryKey = "we_chat_message_industry"

	// CONFIG_LIST
	ConfigList = "config_list"

	// LOGISTICS_KEY
	LogisticsKey = "logistics_"

	// CONFIG_KEY_SITE_URL
	ConfigKeySiteURL = "site_url"

	// CONFIG_KEY_API_URL
	ConfigKeyAPIURL = "api_url"

	// CONFIG_KEY_FRONT_API_URL
	ConfigKeyFrontAPIURL = "front_api_url"

	// CONFIG_KEY_SITE_LOGO
	ConfigKeySiteLogo = "mobile_top_logo"

	// CONFIG_KEY_MOBILE_LOGIN_LOGO
	ConfigKeyMobileLoginLogo = "mobile_login_logo"

	// CONFIG_KEY_SITE_NAME
	ConfigKeySiteName = "site_name"

	// CONFIG_BANK_LIST
	ConfigBankList = "user_extract_bank"

	// CONFIG_RECHARGE_ATTENTION
	ConfigRechargeAttention = "recharge_attention"

	// CONFIG_KEY_PAY_WE_CHAT_APP_ID
	ConfigKeyPayWeChatAppID = "pay_weixin_appid"

	// CONFIG_KEY_PAY_WE_CHAT_MCH_ID
	ConfigKeyPayWeChatMCHID = "pay_weixin_mchid"

	// CONFIG_KEY_PAY_WE_CHAT_APP_SECRET
	ConfigKeyPayWeChatAppSecret = "pay_weixin_appsecret"

	// CONFIG_KEY_PAY_WE_CHAT_APP_KEY
	ConfigKeyPayWeChatAppKey = "pay_weixin_key"

	// CONFIG_KEY_PAY_ROUTINE_APP_ID
	ConfigKeyPayRoutineAppID = "pay_routine_appid"

	// CONFIG_KEY_PAY_ROUTINE_MCH_ID
	ConfigKeyPayRoutineMCHID = "pay_routine_mchid"

	// CONFIG_KEY_PAY_ROUTINE_APP_SECRET
	ConfigKeyPayRoutineAppSecret = "pay_routine_appsecret"

	// CONFIG_KEY_PAY_ROUTINE_APP_KEY
	ConfigKeyPayRoutineAppKey = "pay_routine_key"

	// CONFIG_KEY_PAY_WE_CHAT_APP_APP_ID
	ConfigKeyPayWeChatAppAppID = "pay_weixin_app_appid"

	// CONFIG_KEY_PAY_WE_CHAT_APP_MCH_ID
	ConfigKeyPayWeChatAppMCHID = "pay_weixin_app_mchid"

	// CONFIG_KEY_PAY_WE_CHAT_APP_APP_KEY
	ConfigKeyPayWeChatAppAppKey = "pay_weixin_app_key"

	// CONFIG_KEY_RECHARGE_MIN_AMOUNT
	ConfigKeyRechargeMinAmount = "store_user_min_recharge"

	// CONFIG_KEY_LOGISTICS_APP_CODE
	ConfigKeyLogisticsAppCode = "system_express_app_code"

	// CONFIG_KEY_YZF_H5_URL
	ConfigKeyYZFH5URL = "yzf_h5_url"

	// CONFIG_KEY_CONSUMER_HOTLINE
	ConfigKeyConsumerHotline = "consumer_hotline"

	// CONFIG_KEY_TELEPHONE_SERVICE_SWITCH
	ConfigKeyTelephoneServiceSwitch = "telephone_service_switch"

	// CONFIG_CATEGORY_CONFIG
	ConfigCategoryConfig = "category_page_config"

	// CONFIG_IS_SHOW_CATEGORY
	ConfigIsShowCategory = "is_show_category"

	// CONFIG_IS_PRODUCT_LIST_STYLE
	ConfigIsProductListStyle = "homePageSaleListStyle"

	// CONFIG_APP_VERSION
	ConfigAppVersion = "app_version"

	// CONFIG_APP_ANDROID_ADDRESS
	ConfigAppAndroidAddress = "android_address"

	// CONFIG_APP_IOS_ADDRESS
	ConfigAppIOSAddress = "ios_address"

	// CONFIG_APP_OPEN_UPGRADE
	ConfigAppOpenUpgrade = "open_upgrade"

	// CONFIG_KEY_STORE_BROKERAGE_LEVEL
	ConfigKeyStoreBrokerageLevel = "store_brokerage_rate_num"

	// CONFIG_KEY_STORE_BROKERAGE_RATE_ONE
	ConfigKeyStoreBrokerageRateOne = "store_brokerage_ratio"

	// CONFIG_KEY_STORE_BROKERAGE_RATE_TWO
	ConfigKeyStoreBrokerageRateTwo = "store_brokerage_two"

	// CONFIG_KEY_STORE_BROKERAGE_USER_EXTRACT_MIN_PRICE
	ConfigKeyStoreBrokerageUserExtractMinPrice = "user_extract_min_price"

	// CONFIG_KEY_STORE_BROKERAGE_MODEL
	ConfigKeyStoreBrokerageModel = "store_brokerage_status"

	// CONFIG_KEY_STORE_BROKERAGE_USER_EXTRACT_BANK
	ConfigKeyStoreBrokerageUserExtractBank = "user_extract_bank"

	// CONFIG_KEY_STORE_BROKERAGE_EXTRACT_TIME
	ConfigKeyStoreBrokerageExtractTime = "extract_time"

	// CONFIG_KEY_STORE_INTEGRAL_EXTRACT_TIME
	ConfigKeyStoreIntegralExtractTime = "freeze_integral_day"

	// CONFIG_KEY_STORE_BROKERAGE_PERSON_PRICE
	ConfigKeyStoreBrokeragePersonPrice = "store_brokerage_price"

	// CONFIG_KEY_STORE_BROKERAGE_IS_OPEN
	ConfigKeyStoreBrokerageIsOpen = "brokerage_func_status"

	// CONFIG_KEY_STORE_BROKERAGE_BIND_TYPE
	ConfigKeyStoreBrokerageBindType = "brokerageBindind"

	// CONFIG_KEY_DISTRIBUTION_TYPE
	ConfigKeyDistributionType = "brokerage_bindind"

	// CONFIG_KEY_SMS_CODE_EXPIRE
	ConfigKeySMSCodeExpire = "sms_code_expire"

	// CONFIG_FORM_ID_INDEX
	ConfigFormIDIndex = 133 // 首页配置

	// CONFIG_FORM_ID_PUBLIC
	ConfigFormIDPublic = 65 // 公众号配置

	// THIRD_LOGIN_TOKEN_TYPE_PUBLIC
	ThirdLoginTokenTypePublic = 1 // 公众号

	// THIRD_LOGIN_TOKEN_TYPE_PROGRAM
	ThirdLoginTokenTypeProgram = 2 // 小程序

	// THIRD_LOGIN_TOKEN_TYPE_UNION_ID
	ThirdLoginTokenTypeUnionID = 3 // UnionID

	// THIRD_ADMIN_LOGIN_TOKEN_TYPE_PUBLIC
	ThirdAdminLoginTokenTypePublic = 4 // 后台登录公众号

	// THIRD_LOGIN_TOKEN_TYPE_IOS_WX
	ThirdLoginTokenTypeIOSWX = 5 // iOS 微信

	// THIRD_LOGIN_TOKEN_TYPE_ANDROID_WX
	ThirdLoginTokenTypeAndroidWX = 6 // Android 微信

	// THIRD_LOGIN_TOKEN_TYPE_IOS
	ThirdLoginTokenTypeIOS = 7 // iOS

	// PRODUCT_TYPE_NORMAL
	ProductTypeNormal = 0

	// PRODUCT_TYPE_NORMAL_STR
	ProductTypeNormalStr = "默认"

	// PRODUCT_TYPE_SECKILL
	ProductTypeSeckill = 1

	// PRODUCT_TYPE_SECKILL_STR
	ProductTypeSeckillStr = "秒杀"

	// PRODUCT_TYPE_BARGAIN
	ProductTypeBargain = 2

	// PRODUCT_TYPE_BARGAIN_STR
	ProductTypeBargainStr = "砍价"

	// PRODUCT_TYPE_PINGTUAN
	ProductTypePintuan = 3

	// PRODUCT_TYPE_PINGTUAN_STR
	ProductTypePintuanStr = "拼团"

	// PRODUCT_TYPE_COMPONENT
	ProductTypeComponent = 4

	// PRODUCT_TYPE_COMPONENT_STR
	ProductTypeComponentStr = "组件"

	// PRODUCT_TYPE_GROUP
	ProductTypeGroup = 0

	// GROUP_DATA_ID_INDEX_BEST_BANNER
	GroupDataIDIndexBestBanner = 37 // 中部推荐banner图

	// GROUP_DATA_ID_INDEX_BANNER
	GroupDataIDIndexBanner = 48 // 首页banner滚动图

	// GROUP_DATA_ID_INDEX_RECOMMEND_BANNER
	GroupDataIDIndexRecommendBanner = 52 // 首页精品推荐Banner图片

	// GROUP_DATA_ID_ORDER_STATUS_PIC
	GroupDataIDOrderStatusPic = 53 // 订单详情状态图

	// GROUP_DATA_ID_USER_CENTER_MENU
	GroupDataIDUserCenterMenu = 54 // 个人中心菜单

	// GROUP_DATA_ID_SIGN
	GroupDataIDSign = 55 // 签到配置

	// GROUP_DATA_ID_HOT_SEARCH
	GroupDataIDHotSearch = 56 // 热门搜索

	// GROUP_DATA_ID_INDEX_HOT_BANNER
	GroupDataIDIndexHotBanner = 57 // 热门榜单推荐Banner图片

	// GROUP_DATA_ID_INDEX_NEW_BANNER
	GroupDataIDIndexNewBanner = 58 // 首发新品推荐Banner图片

	// GROUP_DATA_ID_INDEX_BENEFIT_BANNER
	GroupDataIDIndexBenefitBanner = 59 // 首页促销单品推荐Banner图片

	// GROUP_DATA_ID_SPREAD_BANNER_LIST
	GroupDataIDSpreadBannerList = 60 // 推广海报图

	// GROUP_DATA_ID_RECHARGE_LIST
	GroupDataIDRechargeList = 62 // 充值金额设置

	// GROUP_DATA_ID_USER_CENTER_BANNER
	GroupDataIDUserCenterBanner = 65 // 个人中心轮播图

	// GROUP_DATA_ID_INDEX_MENU
	GroupDataIDIndexMenu = 67 // 导航模块

	// GROUP_DATA_ID_INDEX_NEWS_BANNER
	GroupDataIDIndexNewsBanner = 68 // 首页滚动新闻

	// GROUP_DATA_ID_INDEX_ACTIVITY_BANNER
	GroupDataIDIndexActivityBanner = 69 // 首页活动区域图片

	// GROUP_DATA_ID_INDEX_EX_BANNER
	GroupDataIDIndexExBanner = 70 // 首页超值爆款

	// GROUP_DATA_ID_INDEX_KEYWORDS
	GroupDataIDIndexKeywords = 71 // 热门搜索

	// GROUP_DATA_ID_ADMIN_LOGIN_BANNER_IMAGE_LIST
	GroupDataIDAdminLoginBannerImageList = 72 // 后台登录页面轮播图

	// GROUP_DATA_ID_COMBINATION_LIST_BANNNER
	GroupDataIDCombinationListBanner = 73 // 拼团列表banner

	// SIGN_TYPE_INTEGRAL
	SignTypeIntegral = 1 // 积分

	// SIGN_TYPE_EXPERIENCE
	SignTypeExperience = 2 // 经验

	// SIGN_TYPE_INTEGRAL_TITLE
	SignTypeIntegralTitle = "签到积分奖励" // 积分

	// SIGN_TYPE_EXPERIENCE_TITLE
	SignTypeExperienceTitle = "签到经验奖励" // 经验

	// SEARCH_DATE_DAY
	SearchDateDay = "today" // 今天

	// SEARCH_DATE_YESTERDAY
	SearchDateYesterday = "yesterday" // 昨天

	// SEARCH_DATE_LATELY_7
	SearchDateLately7 = "lately7" // 最近7天

	// SEARCH_DATE_LATELY_30
	SearchDateLately30 = "lately30" // 最近30天

	// SEARCH_DATE_WEEK
	SearchDateWeek = "week" // 本周

	// SEARCH_DATE_PRE_WEEK
	SearchDatePreWeek = "preWeek" // 上周

	// SEARCH_DATE_MONTH
	SearchDateMonth = "month" // 本月

	// SEARCH_DATE_PRE_MONTH
	SearchDatePreMonth = "preMonth" // 上月

	// SEARCH_DATE_YEAR
	SearchDateYear = "year" // 年

	// SEARCH_DATE_PRE_YEAR
	SearchDatePreYear = "preYear" // 上一年

	// CATEGORY_TYPE_PRODUCT
	CategoryTypeProduct = 1 // 产品分类

	// CATEGORY_TYPE_ATTACHMENT
	CategoryTypeAttachment = 2 // 附件分类

	// CATEGORY_TYPE_ARTICLE
	CategoryTypeArticle = 3 // 文章分类

	// CATEGORY_TYPE_SET
	CategoryTypeSet = 4 // 设置分类

	// CATEGORY_TYPE_MENU
	CategoryTypeMenu = 5 // 菜单分类

	// CATEGORY_TYPE_CONFIG
	CategoryTypeConfig = 6 // 配置分类

	// CATEGORY_TYPE_SKILL
	CategoryTypeSkill = 7 // 秒杀配置

	// INDEX_RECOMMEND_BANNER
	IndexRecommendBanner = 1 // 首页精品推荐Banner图片

	// INDEX_HOT_BANNER
	IndexHotBanner = 2 // 热门榜单推荐Banner图片

	// INDEX_NEW_BANNER
	IndexNewBanner = 3 // 首页首发新品推荐Banner图片

	// INDEX_BENEFIT_BANNER
	IndexBenefitBanner = 4 // 首页促销单品推荐Banner图片

	// INDEX_LIMIT_DEFAULT
	IndexLimitDefault = 3 // 首页默认list分页条数

	// INDEX_GOOD_BANNER
	IndexGoodBanner = 5 // 优选推荐

	// INDEX_BAST_LIMIT
	IndexBastLimit = "bastNumber" // 精品推荐个数

	// INDEX_FIRST_LIMIT
	IndexFirstLimit = "firstNumber" // 首发新品个数

	// INDEX_SALES_LIMIT
	IndexSalesLimit = "promotionNumber" // 促销单品个数

	// INDEX_HOT_LIMIT
	IndexHotLimit = "hotNumber" // 热门推荐个数

	// USER_BILL_CATEGORY_MONEY
	UserBillCategoryMoney = "now_money" // 用户余额

	// USER_BILL_CATEGORY_INTEGRAL
	UserBillCategoryIntegral = "integral" // 积分

	// USER_BILL_CATEGORY_SHARE
	UserBillCategoryShare = "share" // 分享

	// USER_BILL_CATEGORY_EXPERIENCE
	UserBillCategoryExperience = "experience" // 经验

	// USER_BILL_CATEGORY_BROKERAGE_PRICE
	UserBillCategoryBrokeragePrice = "brokerage_price" // 佣金金额

	// USER_BILL_CATEGORY_SIGN_NUM
	UserBillCategorySignNum = "sign_num" // 签到天数

	// USER_BILL_TYPE_BROKERAGE
	UserBillTypeBrokerage = "brokerage" // 推广佣金

	// USER_BILL_TYPE_DEDUCTION
	UserBillTypeDeduction = "deduction" // 抵扣

	// USER_BILL_TYPE_EXTRACT
	UserBillTypeExtract = "extract" // 提现

	// USER_BILL_TYPE_TRANSFER_IN
	UserBillTypeTransferIn = "transferIn" // 佣金转入余额

	// USER_BILL_TYPE_GAIN
	UserBillTypeGain = "gain" // 购买商品赠送

	// USER_BILL_TYPE_PAY_MONEY
	UserBillTypePayMoney = "pay_money" // 购买

	// USER_BILL_TYPE_PAY_PRODUCT
	UserBillTypePayProduct = "pay_product" // 购买商品

	// USER_BILL_TYPE_PAY_PRODUCT_INTEGRAL_BACK
	UserBillTypePayProductIntegralBack = "pay_product_integral_back" // 商品退积分

	// USER_BILL_TYPE_PAY_PRODUCT_REFUND
	UserBillTypePayProductRefund = "pay_product_refund" // 商品退款

	// USER_BILL_TYPE_RECHARGE
	UserBillTypeRecharge = "recharge" // 佣金转入

	// USER_BILL_TYPE_PAY_RECHARGE
	UserBillTypePayRecharge = "pay_recharge" // 充值

	// USER_BILL_TYPE_SHARE
	UserBillTypeShare = "share" // 用户分享记录

	// USER_BILL_TYPE_SIGN
	UserBillTypeSign = "sign" // 签到

	// USER_BILL_TYPE_ORDER
	UserBillTypeOrder = "order" // 订单

	// USER_BILL_TYPE_PAY_ORDER
	UserBillTypePayOrder = "pay_order" // 订单支付

	// USER_BILL_TYPE_SYSTEM_ADD
	UserBillTypeSystemAdd = "system_add" // 系统增加

	// USER_BILL_TYPE_SYSTEM_SUB
	UserBillTypeSystemSub = "system_sub" // 系统减少

	// USER_BILL_TYPE_PAY_MEMBER
	UserBillTypePayMember = "pay_member" // 会员支付

	// USER_BILL_TYPE_OFFLINE_SCAN
	UserBillTypeOfflineScan = "offline_scan" // 线下支付

	// USER_BILL_TYPE_USER_RECHARGE_REFUND
	UserBillTypeUserRechargeRefund = "user_recharge_refund" // 用户充值退款

	// ORDER_STATUS_ALL
	OrderStatusAll = "all" // 所有

	// ORDER_STATUS_UNPAID
	OrderStatusUnpaid = "unPaid" // 未支付

	// ORDER_STATUS_NOT_SHIPPED
	OrderStatusNotShipped = "notShipped" // 未发货

	// ORDER_STATUS_SPIKE
	OrderStatusSpike = "spike" // 待收货

	// ORDER_STATUS_BARGAIN
	OrderStatusBargain = "bargain" // 已收货待评价

	// ORDER_STATUS_COMPLETE
	OrderStatusComplete = "complete" // 交易完成

	// ORDER_STATUS_TOBE_WRITTEN_OFF
	OrderStatusToBeWrittenOff = "toBeWrittenOff" // 待核销

	// ORDER_STATUS_APPLY_REFUNDING
	OrderStatusApplyRefunding = "applyRefund" // 申请退款

	// ORDER_STATUS_REFUNDING
	OrderStatusRefunding = "refunding" // 退款中

	// ORDER_STATUS_REFUNDED
	OrderStatusRefunded = "refunded" // 已退款

	// ORDER_STATUS_DELETED
	OrderStatusDeleted = "deleted" // 已删除

	// ORDER_STATUS_STR_UNPAID
	OrderStatusStrUnpaid = "未支付" // 未支付

	// ORDER_STATUS_STR_NOT_SHIPPED
	OrderStatusStrNotShipped = "未发货" // 未发货

	// ORDER_STATUS_STR_SPIKE
	OrderStatusStrSpike = "待收货" // 待收货

	// ORDER_STATUS_STR_BARGAIN
	OrderStatusStrBargain = "待评价" // 已收货待评价

	// ORDER_STATUS_STR_TAKE
	OrderStatusStrTake = "用户已收货" // 用户已收货

	// ORDER_STATUS_STR_COMPLETE
	OrderStatusStrComplete = "交易完成" // 交易完成

	// ORDER_STATUS_STR_TOBE_WRITTEN_OFF
	OrderStatusStrToBeWrittenOff = "待核销" // 待核销

	// ORDER_STATUS_STR_APPLY_REFUNDING
	OrderStatusStrApplyRefunding = "申请退款" // 申请退款

	// ORDER_STATUS_STR_REFUNDING
	OrderStatusStrRefunding = "退款中" // 退款中

	// ORDER_STATUS_STR_REFUNDED
	OrderStatusStrRefunded = "已退款" // 已退款

	// ORDER_STATUS_STR_DELETED
	OrderStatusStrDeleted = "已删除" // 已删除

	// ORDER_STATUS_H5_UNPAID
	OrderStatusH5Unpaid = 0 // 未支付

	// ORDER_STATUS_H5_NOT_SHIPPED
	OrderStatusH5NotShipped = 1 // 待发货

	// ORDER_STATUS_H5_SPIKE
	OrderStatusH5Spike = 2 // 待收货

	// ORDER_STATUS_H5_JUDGE
	OrderStatusH5Judge = 3 // 待评价

	// ORDER_STATUS_H5_COMPLETE
	OrderStatusH5Complete = 4 // 已完成

	// ORDER_STATUS_H5_VERIFICATION
	OrderStatusH5Verification = 5 // 待核销

	// ORDER_STATUS_H5_REFUNDING
	OrderStatusH5Refunding = -1 // 退款中

	// ORDER_STATUS_H5_REFUNDED
	OrderStatusH5Refunded = -2 // 已退款

	// ORDER_STATUS_H5_REFUND
	OrderStatusH5Refund = -3 // 退款

	// ORDER_STATUS_INT_PAID
	OrderStatusIntPaid = 0 // 已支付

	// ORDER_STATUS_INT_SPIKE
	OrderStatusIntSpike = 1 // 待收货

	// ORDER_STATUS_INT_BARGAIN
	OrderStatusIntBargain = 2 // 已收货，待评价

	// ORDER_STATUS_INT_COMPLETE
	OrderStatusIntComplete = 3 // 已完成

	// ORDER_STATUS_STR_SPIKE_KEY
	OrderStatusStrSpikeKey = "send" // 待收货 KEY

	// ORDER_LOG_REFUND_PRICE
	OrderLogRefundPrice = "refund_price" // 退款

	// ORDER_LOG_EXPRESS
	OrderLogExpress = "express" // 快递

	// ORDER_LOG_DELIVERY
	OrderLogDelivery = "delivery" // 送货

	// ORDER_LOG_DELIVERY_GOODS
	OrderLogDeliveryGoods = "delivery_goods" // 送货

	// ORDER_LOG_REFUND_REFUSE
	OrderLogRefundRefuse = "refund_refuse" // 不退款

	// ORDER_LOG_REFUND_APPLY
	OrderLogRefundApply = "apply_refund" //

	// ORDER_LOG_PAY_SUCCESS
	OrderLogPaySuccess = "pay_success" // 支付成功

	// ORDER_LOG_DELIVERY_VI
	OrderLogDeliveryVI = "delivery_fictitious" // 虚拟发货

	// ORDER_LOG_EDIT
	OrderLogEdit = "order_edit" // 编辑订单

	// ORDER_LOG_PAY_OFFLINE
	OrderLogPayOffline = "offline" // 线下付款订单

	// ORDER_CASH_CONFIRM
	OrderCashConfirm = 60

	// ORDER_CACHE_PER
	OrderCachePer = "ORDER_CACHE:" // redis缓存订单前缀

	// ORDER_LOG_MESSAGE_REFUND_PRICE
	OrderLogMessageRefundPrice = "退款给用户{amount}元" // 退款

	// ORDER_LOG_MESSAGE_EXPRESS
	OrderLogMessageExpress = "已发货 快递公司：{deliveryName}, 快递单号：{deliveryCode}" // 快递

	// ORDER_LOG_MESSAGE_DELIVERY
	OrderLogMessageDelivery = "已配送 发货人：{deliveryName}, 发货人电话：{deliveryCode}" // 送货

	// ORDER_LOG_MESSAGE_DELIVERY_FICTITIOUS
	OrderLogMessageDeliveryFictitious = "已虚拟发货" // 已虚拟发货

	// ORDER_LOG_MESSAGE_REFUND_REFUSE
	OrderLogMessageRefundRefuse = "不退款款因：{reason}" // 不退款款因

	// ORDER_LOG_MESSAGE_PAY_SUCCESS
	OrderLogMessagePaySuccess = "用户付款成功" // 用户付款成功

	// ORDER_NO_PREFIX_WE_CHAT
	OrderNoPrefixWeChat = "wx" // 微信平台下单订单号前缀

	// ORDER_NO_PREFIX_H5
	OrderNoPrefixH5 = "h5" // 微信平台下单订单号前缀

	// ORDER_PAY_CHANNEL_PUBLIC
	OrderPayChannelPublic = 0 // 公众号

	// ORDER_PAY_CHANNEL_PROGRAM
	OrderPayChannelProgram = 1 // 小程序

	// ORDER_PAY_CHANNEL_H5
	OrderPayChannelH5 = 2 // H5

	// ORDER_PAY_CHANNEL_YUE
	OrderPayChannelYue = 3 // 余额

	// ORDER_PAY_CHANNEL_APP_IOS
	OrderPayChannelAppIOS = 4 // app-ios

	// ORDER_PAY_CHANNEL_APP_ANDROID
	OrderPayChannelAppAndroid = 5 // app-android

	// WE_CHAT_TEMP_KEY_FIRST
	WeChatTempKeyFirst = "first"

	// WE_CHAT_TEMP_KEY_END
	WeChatTempKeyEnd = "remark"

	// WE_CHAT_TEMP_KEY_COMBINATION_SUCCESS
	WeChatTempKeyCombinationSuccess = "OPENTM407456411" // 拼团成功

	// WE_CHAT_TEMP_KEY_BARGAIN_SUCCESS
	WeChatTempKeyBargainSuccess = "OPENTM410292733" // 砍价成功

	// WE_CHAT_TEMP_KEY_EXPRESS
	WeChatTempKeyExpress = "OPENTM200565259" // 订单发货提醒

	// WE_CHAT_TEMP_KEY_DELIVERY
	WeChatTempKeyDelivery = "OPENTM207707249" // 订单配送通知

	// WE_CHAT_TEMP_KEY_ORDER_PAY
	WeChatTempKeyOrderPay = "OPENTM207791277" // 订单支付成功通知

	// WE_CHAT_TEMP_KEY_ORDER_RECEIVING
	WeChatTempKeyOrderReceiving = "OPENTM413386489" // 订单收货通知

	// WE_CHAT_TEMP_KEY_ORDER_REFUND
	WeChatTempKeyOrderRefund = "OPENTM207791277" // 退款进度通知

	// WE_CHAT_PUBLIC_TEMP_KEY_RECHARGE
	WeChatPublicTempKeyRecharge = "OPENTM200565260" // 充值成功

	// WE_CHAT_PROGRAM_TEMP_KEY_COMBINATION_SUCCESS
	WeChatProgramTempKeyCombinationSuccess = "5164" // 拼团成功

	// WE_CHAT_PROGRAM_TEMP_KEY_BARGAIN_SUCCESS
	WeChatProgramTempKeyBargainSuccess = "2920" // 砍价成功

	// WE_CHAT_PROGRAM_TEMP_KEY_EXPRESS
	WeChatProgramTempKeyExpress = "467" // 订单发货提醒

	// WE_CHAT_PROGRAM_TEMP_KEY_DELIVERY
	WeChatProgramTempKeyDelivery = "14198" // 订单配送通知

	// WE_CHAT_PROGRAM_TEMP_KEY_ORDER_PAY
	WeChatProgramTempKeyOrderPay = "516" // 订单支付成功通知

	// WE_CHAT_PROGRAM_TEMP_KEY_ORDER_RECEIVING
	WeChatProgramTempKeyOrderReceiving = "9283" // 订单收货通知

	// WE_CHAT_PROGRAM_TEMP_KEY_RECHARGE
	WeChatProgramTempKeyRecharge = "OPENTM200565260"

	// COMMON_SWITCH_CLOSE
	CommonSwitchClose = "0" // 公共开关：0关闭

	// COMMON_SWITCH_OPEN
	CommonSwitchOpen = "1" // 公共开关：1开启

	// JS_CONFIG_CRMEB_CHAT_TONGJI
	JSConfigCrmebChatTongji = "crmeb_tongji_js" // CRMEB chat 统计
)

// 可变变量
var (
	// CND_URL CND URL测试用
	CNDURL = "https://wuht-1300909283.cos.ap-chengdu.myqcloud.com"

	// LOGISTICS_API_URL 物流信息缓存
	LogisticsAPIURL = "https://wuliu.market.alicloudapi.com/kdi"

	// RESULT_ORDER_NOTFOUND 订单基本操作字样
	ResultOrderNotFound = "订单号 ${orderCode} 未找到"

	// RESULT_ORDER_NOTFOUND_IN_ID 订单id未找到
	ResultOrderNotFoundInID = "订单id ${orderId} 未找到"

	// RESULT_ORDER_PAYED 订单已支付
	ResultOrderPayed = "订单号 ${orderCode} 已支付"

	// RESULT_ORDER_EDIT_PRICE_SAME 修改价格不能和支付价格相同
	ResultOrderEditPriceSame = "修改价格不能和支付价格相同 原价 ${oldPrice} 修改价 ${editPrice}"

	// RESULT_ORDER_EDIT_PRICE_SUCCESS 修改价格成功
	ResultOrderEditPriceSuccess = "订单号 ${orderNo} 修改价格 ${price} 成功"

	// RESULT_ORDER_EDIT_PRICE_LOGS 修改价格日志
	ResultOrderEditPriceLogs = "订单价格 ${orderPrice} 修改实际支付金额为 ${price} 元"

	// RESULT_ORDER_PAY_OFFLINE 订单线下付款成功
	ResultOrderPayOffline = "订单号 ${orderNo} 现在付款 ${price} 成功"

	// RESULT_VERIFICATION_ORDER_NOT_FUND 核销码的订单未找到
	ResultVerificationOrderNotFound = "核销码 ${vCode} 的订单未找到"

	// RESULT_VERIFICATION_ORDER_VED 核销码的订单已核销
	ResultVerificationOrderVED = "核销码 ${vCode} 的订单已核销"

	// RESULT_VERIFICATION_NOTAUTH 没有核销权限
	ResultVerificationNotAuth = "没有核销权限"

	// RESULT_VERIFICATION_USER_EXIST 当前用户已经是核销员
	ResultVerificationUserExist = "当前用户已经是核销员"

	// RESULT_QRCODE_PRAMERROR QRcode参数错误
	ResultQRCodeParamError = "生成二维码参数不合法"

	// BARGAIN_TATIO_DOWN 砍价计算比例下行
	BargainTatioDown = "0.2"

	// BARGAIN_TATIO_UP 砍价计算比例上行
	BargainTatioUp = "0.8"
)

// 其他常量
const (
	// STORE_REPLY_TYPE_PRODUCT 商品评论类型——普通商品
	StoreReplyTypeProduct = "product"

	// STORE_REPLY_TYPE_SECKILL 商品评论类型——秒杀
	StoreReplyTypeSeckill = "seckill"

	// STORE_REPLY_TYPE_PINTUAN 商品评论类型——拼团
	StoreReplyTypePintuan = "pintuan"

	// STORE_REPLY_TYPE_BARGAIN 商品评论类型——砍价
	StoreReplyTypeBargain = "bargain"

	// PRODUCT_LOG_KEY 商品记录Key（pv、uv）用
	ProductLogKey = "visit_log_key"

	// FAIL
	Fail = "FAIL"

	// SUCCESS
	Success = "SUCCESS"

	// ORDER_AUTO_CANCEL_KEY 订单取消Key
	OrderAutoCancelKey = "order_auto_cancel_key"
)
