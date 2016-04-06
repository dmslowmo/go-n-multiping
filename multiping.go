package main

import (
	"bufio"
	"fmt"
	"os"
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

	inFile, _ := os.Open(os.Args[1])

	scanner := bufio.NewScanner(inFile)

	var wg sync.WaitGroup

	for scanner.Scan() {
		wg.Add(1)
		host := scanner.Text()
		go ping(&wg, host)
	}
	wg.Wait()

	endTime := time.Now().UTC()
	fmt.Println(endTime.Sub(startTime))
}
