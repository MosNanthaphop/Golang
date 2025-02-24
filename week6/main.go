package main

/*
import (
	"errors"
	"fmt"
)

// ประกาศ Function "divide" ให้ส่งค่ากลับเป็น Interface "error"
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("หารด้วยศูนย์ไม่ได้")
	}
	return a / b, nil
}

func main() {
	result, err := divide(5, 0)
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาด", err)
	} else {
		fmt.Printf("5 / 0 = %.2f\n", result)
	}
}*/

import (
	"fmt"
)

/*
// Implement Error Interface (Implement Method Error())
type DivisionError struct {
	dividend float64
	divisor  float64
}

func (de *DivisionError) Error() string {
	return fmt.Sprintf("ไม่สามารถหาร %.2f ด้วย %.2f ได้", de.dividend, de.divisor)
}

func safeDivide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, &DivisionError{dividend, divisor}
	}
	return dividend / divisor, nil
}

func main() {
	result, err := safeDivide(1, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}*/
// Function ที่อาจเกิดข้อผิดพลาดร้ายแรง
func riskyFunction() {
	// defer Function ที่จะจับ Panic
	defer func() {
		// ใช้ recover() เพื่อจับ Panic
		r := recover()
		if r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	// จำลองการเกิดข้อผิดพลาดร้ายแรงด้วย Panic
	panic("เกิดข้อผิดพลาดร้ายแรง!")

	// บรรทัดนี้จะไม่ถูกทำงานเนื่องจากเกิด Panic ก่อนหน้านี้
	fmt.Println("บรรทัดนี้จะไม่ถูกพิมพ์")
}

func main() {
	// เรียกใช้ Function ที่เกิด panic
	riskyFunction()

	// แสดงข้อความว่าโปรแกรมยังทำงานต่อได้
	fmt.Println("โปรแกรมยังทำงานต่อไปได้")
}
