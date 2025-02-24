package main

import (
	"fmt"
	"sync"
)

func countLuckyNumber(name string, wg *sync.WaitGroup) {
	count := 0
	fmt.Printf("%s %d done\n", name, count)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go countLuckyNumber("goroutine", &wg)
	go countLuckyNumber("goroutine2", &wg)
	wg.Wait()
	fmt.Println("All done")
}
