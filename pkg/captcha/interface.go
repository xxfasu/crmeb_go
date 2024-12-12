package captcha

type Captcha interface {
	Gen() (string, string, error)
	Verify(id, answer string) bool
}
