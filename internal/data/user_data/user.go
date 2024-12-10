package user_data

import "crmeb_go/internal/model"

type LoginResp struct {
	AccessToken string `json:"accessToken"`
}

type LoginUserData struct {
	Token         string            `json:"token"`          // 用户唯一标识
	LoginTime     int64             `json:"login_time"`     // 登陆时间
	ExpireTime    int64             `json:"expire_time"`    // 过期时间
	Ipaddr        string            `json:"ipaddr"`         // 登录IP地址
	LoginLocation string            `json:"login_location"` // 登录地点
	Browser       string            `json:"browser"`        // 浏览器类型
	OS            string            `json:"os"`             // 操作系统
	User          model.SystemAdmin `json:"user"`           // 用户信息
}

type Condition struct {
	UserID   string
	Nickname string
	Email    string
}
