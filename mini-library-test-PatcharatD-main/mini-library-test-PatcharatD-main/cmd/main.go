// ไฟล์: cmd/main.go
package main

import (
	"fmt"
	"log"

	"myproject/library" // เปลี่ยนเป็นโมดูล/เส้นทาง Import ของคุณตามจริง
)

func main() {
	// สร้าง Library ว่าง
	lib := library.Library{}

	// เพิ่มหนังสือเล่มแรก
	err := lib.AddBook(library.Book{
		ID:     101,
		Title:  "Golang Fundamentals",
		Author: "Alice",
	})
	if err != nil {
		log.Fatalf("AddBook error: %v", err)
	}

	// เพิ่มหนังสือเล่มที่สอง
	err = lib.AddBook(library.Book{
		ID:     102,
		Title:  "Learn Docker",
		Author: "Bob",
	})
	if err != nil {
		log.Fatalf("AddBook error: %v", err)
	}

	// ลองเพิ่มหนังสือด้วย ID ซ้ำ
	err = lib.AddBook(library.Book{
		ID:     101,
		Title:  "Duplicate ID Example",
		Author: "Charlie",
	})
	if err != nil {
		fmt.Println("Error when adding duplicate book ID:", err)
	} else {
		fmt.Println("AddBook with duplicate ID (unexpected, should fail).")
	}

	// ค้นหาหนังสือตาม ID
	bk, err := lib.GetBookByID(102)
	if err != nil {
		fmt.Println("GetBookByID error:", err)
	} else {
		fmt.Printf("Found book (ID=%d): %s by %s\n", bk.ID, bk.Title, bk.Author)
	}

	// ค้นหาหนังสือที่ไม่มีใน Library
	bk, err = lib.GetBookByID(999)
	if err != nil {
		fmt.Println("GetBookByID(999) error:", err)
	} else {
		fmt.Printf("Found book: %+v (unexpected)\n", bk)
	}

	// แสดงจำนวนหนังสือที่อยู่ใน Library
	fmt.Printf("Current library has %d books.\n", len(lib.Books))
}
