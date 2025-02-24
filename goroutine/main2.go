package main

import (
	"fmt"
	"time"
)

func prepareDish1(dish string, duration time.Duration) {
	fmt.Printf("เริ่มทำ%s...\n", dish)
	time.Sleep(duration)
	fmt.Printf("%sเสร็จแล้ว!\n", dish)
}

func main2() {
	start := time.Now()
	go prepareDish1("ผัดกะเพรา", 2*time.Second)
	go prepareDish1("ต้มยำกุ้ง", 5*time.Second)

	time.Sleep(5 * time.Second)
	fmt.Printf("ใช้เวลาทั้งหมด: %.2f วินาที\n", time.Since(start).Seconds())
}
