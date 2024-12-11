package response

type SystemLoginResp struct {
	ID       int64  `json:"id"`
	Account  string `json:"account"`
	RealName string `json:"real_name"`
	Token    string `json:"token"`
	IsSMS    bool   `json:"is_sms"`
}
