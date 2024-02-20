package cache_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Anna-Karenina/ht_in_mem_cache/pkg/cache"
)

func TestCacheCleaner(t *testing.T) {
	block := make(chan string)
	c := cache.NewCache()
	c.Set("user", 1, time.Duration(time.Second*2))

	go func() {
		for {

			time.Sleep(time.Duration(time.Second * 1))
			fmt.Println(c.Len())
		}

	}()
	fmt.Println(<-block)
}
