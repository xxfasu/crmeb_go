package service_data

type LoginResp struct {
	AccessToken string `json:"accessToken"`
}

type Condition struct {
	UserID   string
	Nickname string
	Email    string
}
