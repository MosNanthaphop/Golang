package main

import (
	"fmt"
	"runtime"
	"time"
)

func main1() {
	// อ่านข้อมูลการใช้หน่วยความจำก่อนสร้าง Goroutine จำนวนมาก
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	fmt.Printf("จำนวน Goroutine เริ่มต้น: %d\n", runtime.NumGoroutine())

	// สร้าง Goroutine จำนวนมาก
	for i := 0; i < 5000000; i++ {
		go func() { //Go หน้า func เป็น Goroutine
			time.Sleep(5 * time.Second)
		}()
	}

	// รอให้ Goroutine เริ่มทำงาน
	time.Sleep(time.Second)

	// อ่านข้อมูลการใช้หน่วยความจำหลังสร้าง Goroutine
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	// จำนวนและผลต่างการใช้หน่วยความจำ
	fmt.Printf("จำนวน Goroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("การใช้หน่วยความจำก่อนสร้าง Goroutine: %d Byte\n", m1.Alloc)
	fmt.Printf("การใช้หน่วยความจำหลังสร้าง Goroutine: %d Byte\n", m2.Alloc)
	fmt.Printf("ความแตกต่างของหน่วยความจำที่ใช้เพิ่มขึ้น: %d Byte\n", m2.Alloc-m1.Alloc)
	fmt.Printf("ประมาณการใช้หน่วยความจำต่อ Goroutine: %d Byte\n", (m2.Alloc-m1.Alloc)/5000000)
}
