package cache

import (
	"context"
	"testing"
	"time"
)

func TestGetCache(t *testing.T) {

	cache, err := GetCache(context.Background(), "test")
	if err != nil {
		t.Error(err)
	}
	t.Log(cache)
}

func TestGetCacheOrElse(t *testing.T) {
	value, err := GetCacheOrElse(context.Background(), "test", 1*time.Minute, FetcherFunc(func(ctx context.Context, key string) (string, error) {
		return "FetcherFunc", nil
	}))
	if err != nil {
		t.Error(err)
	}
	t.Log(value)
}

func TestSetCache(t *testing.T) {

	err := SetCache(context.Background(), "test", "test", time.Hour)
	if err != nil {
		t.Error(err)
	}
}
