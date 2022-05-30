package main

import (
	"algorithmSystem/Cache/FIFO"
	"fmt"
	"strconv"
)

func main() {
	cache := FIFO.NewFIFOCache(1024)
	cache.Set("1", 100)
	cache.Set("2", 101)
	cache.Set("3", 103)

	for i := 0; i < 1000000 && cache.UseBytes() < 1024; i++ {
		cache.Set(strconv.Itoa(i), i)
	}

	fmt.Println(cache.Get("127"))
}
