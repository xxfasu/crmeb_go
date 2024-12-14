package crmeb_config

type CrmebConfigData struct {
	Version                    string `json:"version"`                        // 当前代码版本
	Domain                     string `json:"domain"`                         // 待部署的域名
	WechatApiUrl               string `json:"wechat_api_url"`                 // 微信接口中专用服务器 URL
	WechatJsApiDebug           bool   `json:"wechat_js_api_debug"`            // 是否开启微信 JS API 调试模式
	WechatJsApiBeta            bool   `json:"wechat_js_api_beta"`             // 微信 JS API 是否是 Beta 版本
	AsyncConfig                bool   `json:"async_config"`                   // 是否同步配置表数据到 Redis
	AsyncWeChatProgramTempList bool   `json:"async_we_chat_program_tempList"` // 是否同步小程序公共模板库
	ImagePath                  string `json:"image_path"`                     // 本地图片路径配置
}
