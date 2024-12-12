package captcha

import (
	"context"
	"crmeb_go/constants"
	"crmeb_go/pkg/cache"
	"fmt"
	"strings"
	"time"
)

type Store struct {
	cache cache.Cache
	ctx   context.Context
	ttl   time.Duration
}

// Set 将验证码的 id 和 value 存储到 Redis 中
func (s *Store) Set(id string, value string) error {
	value = strings.ToLower(value)
	key := fmt.Sprintf(constants.CaptchaKeyPrefix, id)
	return s.cache.SetCache(s.ctx, key, value, s.ttl)
}

// Get 根据 id 获取验证码的值
// clear = true 时获取后删除该 key，防止重复使用
func (s *Store) Get(id string, clear bool) string {
	key := fmt.Sprintf(constants.CaptchaKeyPrefix, id)
	val, err := s.cache.GetCache(s.ctx, key)
	if err != nil {
		return ""
	}

	if clear {
		s.cache.DelCache(key)
	}

	return val
}

func (s *Store) Verify(id, answer string, clear bool) bool {

	key := fmt.Sprintf(constants.CaptchaKeyPrefix, id)
	val, err := s.cache.GetCache(s.ctx, key)
	if err != nil {
		return false
	}

	if clear {
		s.cache.DelCache(key)
	}

	return answer == val
}
