package main

import (
	"fmt"
	"time"
)

func procedure(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i                     // ส่งค่า i เข้า channel
		time.Sleep(1 * time.Second) //หน่วงเวลา
	}
	close(ch) // ปิด channel เมื่อส่งครบ
}

func consumer(ch chan int) {
	for v := range ch { // รับค่าจาก channel จนกว่าจะปิด
		fmt.Println("รับค่า:", v)
	}
}

func main3() {
	ch := make(chan int)
	go procedure(ch)
	consumer(ch)
}
