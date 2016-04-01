package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const maxConcurrency = 6 //default

func ping(wg *sync.WaitGroup, host string) {
	defer wg.Done()

	//temporary pseudo logic of ping
	if rand.Int()%2 == 1 {
		fmt.Println(host, "is down")
	} else {
		fmt.Println(host, "is up")
	}
}

func main() {
	rand.Seed(42)
	var wg sync.WaitGroup

	hosts := []string{"q", "w", "e", "r", "t", "y"}
	for i := 0; i < len(hosts); i++ {
		wg.Add(1)
		go ping(&wg, hosts[i])
	}
	wg.Wait()
}
