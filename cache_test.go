package generic_test

import (
	"github.com/graphikDB/generic"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	cache := generic.NewCache(1 * time.Minute)
	cache.Set("key", "value", 0)
	if cache.Len() != 1 {
		t.Fatal("expected 1 key")
	}
}
