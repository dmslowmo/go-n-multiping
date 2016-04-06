package main

import (
	"fmt"
	"os/exec"
	"sync"
	"time"
)

func ping(wg *sync.WaitGroup, host string) {
	var (
		err    error
	)
	cmdName := "ping"
	cmdArgs := []string{"-c1", "-W1", host}

	defer wg.Done()
	if _, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil{
		fmt.Println(host, "is down")
	} else {
		fmt.Println(host, "is up")
	}
}

func main() {
	startTime := time.Now().UTC()

	hosts := []string{
			"192.168.2.1", 
			"127.0.0.1", 
			"www.google.com",
			"10.0.10.156",
			}

	var wg sync.WaitGroup

	for i := 0; i < len(hosts); i++ {
		wg.Add(1)
		go ping(&wg, hosts[i])
	}
	wg.Wait()

	endTime := time.Now().UTC()
	fmt.Println(endTime.Sub(startTime))
}
