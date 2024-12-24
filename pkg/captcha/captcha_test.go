package captcha

import (
	"crmeb_go/internal/conf"
	"crmeb_go/internal/redis"
	"crmeb_go/pkg/cache"
	"crmeb_go/pkg/logs"
	"fmt"
	"os"
	"testing"
)

var localCache cache.Cache

func TestMain(m *testing.M) {
	conf.InitConfig("D:\\goproject\\xxfasu\\crmeb_go\\config")
	logs.InitLog()
	client, _ := redis.InitRedis()
	localCache = cache.InitLocalCache(client)
	code := m.Run()
	fmt.Println("test end")
	os.Exit(code)
}

func TestGen(t *testing.T) {
	c := New(localCache)
	gen, s, err := c.Gen()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(gen)
	t.Log(s)
}

func TestVerify(t *testing.T) {
	t.Log(New(localCache).Verify("8EtocHy6jhhgssfNNQoa", "HauH"))
}
