package validation

type Register struct {
	Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type FindUser struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	UserID   string `json:"user_id"`
}
