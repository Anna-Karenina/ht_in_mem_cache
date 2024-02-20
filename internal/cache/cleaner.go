package cache

import "time"

func (c *Cache) cleanerRun() {
	for {
		time.Sleep(time.Duration(time.Second * 5))

		for key, item := range c.items {
			if isOverdue(item) {
				c.Delete(key)
			}
		}
	}
}
