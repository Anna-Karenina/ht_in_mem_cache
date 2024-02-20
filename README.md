# cache-lib

A home-task to write in-memory key:value cache library for Go.

## Quickstart

### Installing

Using libcache is easy. First, use go get to install the latest version of the library.

```sh
go get github.com/Anna-Karenina/ht_in_mem_cache
```

Next, include lib in your application:

```go
import (
    c "github.com/Anna-Karenina/ht_in_mem_cache"
)
```

### Examples

#### Basic

```go
package main
import (
    "fmt"
    c "github.com/Anna-Karenina/ht_in_mem_cache"
)

func main() {
    cache := c.NewCache()

    cache.Set("<key>": interface)
    value, err := cache.Get("<key>")
    cache.Delete("<key>")

    cacheSize:=cache.Len()
}
```
