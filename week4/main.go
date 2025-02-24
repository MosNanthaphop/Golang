package main

import (
	"fmt"
)

func main() {
	/*x := 5
	var p *int = &x // p เก็บที่อยู่ของ x
	fmt.Println("ค่าของ x:", x)
	fmt.Println("ที่อยู่ของ x:", p)
	fmt.Println("ค่าที่ p ชี้ไป:", *p)*/

	/*
		ptr := new(int) // ประกาศแบบอัตโนมัติ
		fmt.Println(ptr)
		fmt.Println(*ptr)*/

	/*
		num := 5
		doubleValue(num)
		fmt.Println("หลัง doubleValue:", num) // ยังคงเป็น 5

		doublePointer(&num)
		fmt.Println("หลัง doublePointer:", num) // กลายเป็น 10*/

	/*
		var ptr *int // ptr เป็น nil pointer
		if ptr != nil {
			fmt.Println(*ptr)
		} else {
			fmt.Println("Pointer is nil")
		}*/

	/*
		myCat := Cat{Name: "ตุนตัง", Breed: "มันช์กิ้น", Age: 2}
		fmt.Println(myCat)

		// การเข้าถึงฟิลด์
		fmt.Println("ชื่อแมว:", myCat.Name)*/

	/*
		myCat := Cat{Name: "ตุนตัง", Breed: "มันช์กิ้น", Age: 2}
		birthday(&myCat)
		fmt.Println(myCat) // {ตุนตัง มันช์กิ้น 3}*/

	/*
		printAnything(42)
		printAnything("Hello")
		printAnything(true)
		printAnything([]int{1, 2, 3})*/

}

func animalMove(m Mover) {
	m.Move()
}

type Mover interface {
	Move()
}

type cat struct {
	name string
	age  int
}

func (c cat) Mover() {
	fmt.Println("cat is moving")
}

func doubleValue(x int) {
	x *= 2
}

func doublePointer(x *int) {
	*x *= 2
}

type Cat struct {
	Name  string
	Breed string
	Age   int
}

func birthday(p *Cat) {
	p.Age++ // ไม่จำเป็นต้องใช้ (*p).Age
}

func printAnything(v interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", v, v)
}
