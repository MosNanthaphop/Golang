package main

import (
	"fmt"
	"sync"
	"time"
)

func PrintNumbers(prefix string, count int) {
	for i := 0; i <= count; i++ {
		fmt.Printf("%s: %d\n", prefix, i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main5() {
	var wg sync.WaitGroup
	wg.Add(2)
	//Job A
	go func() {
		defer wg.Done()
		PrintNumbers("X", 9)
	}()

	//Job B
	go func() {
		defer wg.Done()
		PrintNumbers("Y", 9)
	}()

	wg.Wait()
}
