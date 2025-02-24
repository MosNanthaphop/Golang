package main

import (
	"fmt"
	"myproject/helper"
)

/*
	func createMultiplier(multiplier int) func(int) int {
		return func(x int) int {
			return x * multiplier
		}
	}

	func createCounter() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}

var globalVar = initGlobalVar()

	func initGlobalVar() string {
		fmt.Println("initGlobalVar")
		return "initGlobalVar"
	}

	func init() {
		fmt.Println("init")
	}

	func init() {
		fmt.Println("init2")
	}

func main() {
	s1 := make([]int, 5, 10)
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	s1[0] = 1
	s1[1] = 2
	s1[2] = 3
	fmt.Println(s1)
	s1 = append(s1, 4) //[1 2 3 0 0 4]
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	s2 := s1[2:5]
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	s2 = append(s2, 7, 8)
	fmt.Println(s2)
	fmt.Println(s1)

	s2 = append(s2, 9, 10, 11, 12, 13, 14)
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
*/
//fmt.Println(s1)
func init() {
	fmt.Println("Main initialized")
}

func main() {
	fmt.Println("Hello from main")
	helper.SayHello()
}
