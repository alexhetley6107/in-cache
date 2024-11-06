# in-memory-cache

## Example #1

```
package main

import (
  "fmt"
  cache "github.com/alexhetley6107/in-cache"
)

func main() {
	c := cache.New()

	c.Set("userId", 42)
	userId := c.Get("userId")

	fmt.Println(userId)

	c.Delete("userId")
	userId := c.Get("userId")

	fmt.Println(userId)
}

```