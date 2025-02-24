package main

import (
	"fmt"
	"time"
)

func prepareDish(dish string, duration time.Duration, done chan bool) {
	fmt.Printf("เริ่มทำ%s...\n", dish)
	time.Sleep(duration)
	fmt.Printf("%sเสร็จแล้ว!\n", dish)
	done <- true
}

func main4() {
	start := time.Now()
	done := make(chan bool)

	go prepareDish("ผัดกะเพรา", 2*time.Second, done)
	go prepareDish("ต้มยำกุ้ง", 3*time.Second, done)

	// รอให้ทั้งสองจานเสร็จ
	// เราไม่สนใจค่าที่รับมา (เพราะเราใช้ bool channel เพียงเพื่อส่งสัญญาณ)
	<-done
	<-done

	// คำนวณและแสดงเวลาที่ใช้ทั้งหมด
	elapsedTime := time.Since(start).Seconds()
	fmt.Printf("ใช้เวลาทั้งหมด: %.2f วินาที\n", elapsedTime)
}
