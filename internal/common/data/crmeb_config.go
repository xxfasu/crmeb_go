package data

type CrmebConfig struct {
	Version                    string `json:"version"`                    // 当前代码版本
	Domain                     string `json:"domain"`                     // 待部署的域名
	WechatApiUrl               string `json:"wechatApiUrl"`               // 微信接口中专用服务器 URL
	WechatJsApiDebug           bool   `json:"wechatJsApiDebug"`           // 是否开启微信 JS API 调试模式
	WechatJsApiBeta            bool   `json:"wechatJsApiBeta"`            // 微信 JS API 是否是 Beta 版本
	AsyncConfig                bool   `json:"asyncConfig"`                // 是否同步配置表数据到 Redis
	AsyncWeChatProgramTempList bool   `json:"asyncWeChatProgramTempList"` // 是否同步小程序公共模板库
	ImagePath                  string `json:"imagePath"`                  // 本地图片路径配置
}
