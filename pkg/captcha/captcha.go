package captcha

import (
	"context"
	"crmeb_go/pkg/cache"
	"github.com/mojocn/base64Captcha"
	"strings"
	"time"
)

type captcha struct {
	Captcha *base64Captcha.Captcha
}

func New(cache cache.Cache) Captcha {
	driver := base64Captcha.NewDriverString(
		80,  // height int
		240, // width int
		6,   // noiseCount int
		1,   // showLineOptions int
		4,   // length int
		"123456789abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ", // source string
		nil, // bgColor *color.RGBA
		nil, // fontsStorage FontsStorage
		nil, // fonts []string
	)

	store := &Store{
		cache: cache,
		ttl:   time.Minute * 30,
		ctx:   context.Background(),
	}

	newCaptcha := base64Captcha.NewCaptcha(driver, store)
	return &captcha{
		newCaptcha,
	}
}

func (c *captcha) Gen() (string, string, error) {
	id, b64s, _, err := c.Captcha.Generate()
	return id, b64s, err
}

func (c *captcha) Verify(id, answer string) bool {
	answer = strings.ToLower(answer)
	return c.Captcha.Verify(id, answer, true)
}
