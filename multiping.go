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
	var err error
	cmdName := "ping"
	cmdArgs := []string{"-c1", "-W1", host}

	defer wg.Done()
	if _, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Println(host, "is down")
	} else {
		fmt.Println(host, "is up")
	}
}

func main() {
	// get the starting time
	startTime := time.Now().UTC()

	// read the input file
	inFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inFile.Close() // close the file when main exits
	scanner := bufio.NewScanner(inFile)

	// use a goroutine for each ping
	var wg sync.WaitGroup
	for scanner.Scan() {
		wg.Add(1)
		host := scanner.Text()
		go ping(&wg, host)
	}
	wg.Wait()

	// get the ending time and calculate the total duration
	endTime := time.Now().UTC()
	fmt.Println(endTime.Sub(startTime))
}
