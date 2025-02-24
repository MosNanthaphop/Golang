package main

import (
	"fmt"
	"time"
)

func main() {
	// กำหนดวันสรุปยอดเป็นวันที่ 20 มกราคม 2024
	statementDate := time.Date(2024, time.January, 20, 0, 0, 0, 0, time.UTC)
	fmt.Printf("วันสรุปยอด %v\n\n", statementDate)

	// คำนวณวันครบกำหนดชำระโดยใช้ Add
	dueDateAdd := statementDate.Add(25 * 24 * time.Hour)
	fmt.Printf("วันครบกำหนดชำระ %v\n", dueDateAdd)
	fmt.Printf("จำนวนวันระหว่างวันสรุปยอดและวันครบกำหนด: %d\n\n", int(dueDateAdd.Sub(statementDate).Hours()/24))

	// ทดสอบกับเดือนกุมภาพันธ์ (ปีอธิกสุรทิน)
	statementDateFeb := time.Date(2024, time.February, 20, 0, 0, 0, 0, time.UTC)
	fmt.Printf("วันสรุปยอด (กุมภาพันธ์) %v\n\n", statementDateFeb)

	dueDateAddFeb := statementDateFeb.Add(25 * 24 * time.Hour)
	fmt.Printf("วันครบกำหนดชำระ %v\n", dueDateAddFeb)
	fmt.Printf("จำนวนวันระหว่างวันสรุปยอดและวันครบกำหนด %d\n", int(dueDateAddFeb.Sub(statementDateFeb).Hours()/24))
}
