package daos

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

var once sync.Once
var ch *cache.Cache

func init() {
	once.Do(func() {
		ch = cache.New(5*time.Minute, 10*time.Minute)
	})
}

type Pool struct {
	*cache.Cache
}

func New() *Pool {
	return &Pool{
		Cache: ch,
	}
}
