package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func ping(wg *sync.WaitGroup, host string) {
	defer wg.Done()

	//temporary pseudo logic of ping
	if rand.Int()%2 == 1 {
		fmt.Println(host, "is down")
		time.Sleep(1000 * time.Millisecond) //sleep 1 sec to simulate unreachable host
	} else {
		fmt.Println(host, "is up")
	}
}

func main() {
	startTime := time.Now().UTC()

	rand.Seed(time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond)))
	var wg sync.WaitGroup

	hosts := []string{	"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y",
						"q", "w", "e", "r", "t", "y" }
	for i := 0; i < len(hosts); i++ {
		wg.Add(1)
		go ping(&wg, hosts[i])
	}
	wg.Wait()

	endTime := time.Now().UTC()
	fmt.Println(endTime.Sub(startTime))
}
