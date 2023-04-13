package main

import (
	"fmt"
	"sync"
)

type MapItem struct {
	Key   string
	Value int
}

type Container struct {
	Map map[string]int
	mu  sync.Mutex
}

func (c *Container) set(item MapItem) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Map[item.Key] = item.Value
}

var wg sync.WaitGroup

func main() {
	container := Container{Map: make(map[string]int)}

	map_items := []MapItem{{"a", 1}, {"b", 2}, {"c", 3}}

	for _, v := range map_items {
		wg.Add(1)
		go func(item MapItem) {
			container.set(item)
			wg.Done()
		}(v)
	}

	wg.Wait()

	fmt.Println(container.Map)
}
