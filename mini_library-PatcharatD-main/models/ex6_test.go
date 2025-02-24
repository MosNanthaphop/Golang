package models

import (
	"mini_library/interfaces" // เพิ่มการ import แพ็กเกจ interfaces
	"testing"
)

// TestEx6NewEBook ทดสอบฟังก์ชัน NewEBook ว่าสร้าง EBook ได้ถูกต้องหรือไม่ (ex6)
func TestEx6NewEBook(t *testing.T) {
	title := "Go In Depth"
	author := "Nuttachot Promrit"
	fileSizeMB := 10

	ebook := NewEBook(title, author, fileSizeMB)

	if ebook == nil {
		t.Fatalf("Expected NewEBook to return a non-nil pointer, got nil")
	}

	// ตรวจสอบฟิลด์จาก Book ที่ถูก embed
	if ebook.GetTitle() != title {
		t.Errorf("Expected Title to be '%s', got '%s'", title, ebook.GetTitle())
	}

	if ebook.GetAuthor() != author {
		t.Errorf("Expected Author to be '%s', got '%s'", author, ebook.GetAuthor())
	}

	if ebook.GetStock() != 0 {
		t.Errorf("Expected Stock to be 0, got %d", ebook.GetStock())
	}

	// ตรวจสอบฟิลด์เพิ่มเติม
	if ebook.GetFileSize() != fileSizeMB {
		t.Errorf("Expected FileSizeMB to be %d, got %d", fileSizeMB, ebook.GetFileSize())
	}
}

// TestEx6EBookInfo ทดสอบเมธอด Info() ที่ถูก Override สำหรับ EBook (ex6)
func TestEx6EBookInfo(t *testing.T) {
	title := "Go In Depth"
	author := "Nuttachot Promrit"
	fileSizeMB := 10

	ebook := NewEBook(title, author, fileSizeMB)
	expectedInfo := "Go In Depth by Nuttachot Promrit (E-Book, 10MB)"
	info := ebook.Info()

	if info != expectedInfo {
		t.Errorf("Expected Info() to return '%s', got '%s'", expectedInfo, info)
	}
}

// TestEx6GetFileSize ทดสอบ Getter Method สำหรับ fileSizeMB (ex6)
func TestEx6GetFileSize(t *testing.T) {
	title := "Mastering Go"
	author := "John Doe"
	fileSizeMB := 25

	ebook := NewEBook(title, author, fileSizeMB)

	if ebook.GetFileSize() != fileSizeMB {
		t.Errorf("Expected GetFileSize() to return %d, got %d", fileSizeMB, ebook.GetFileSize())
	}
}

// TestEx6LibraryItemInterface ทดสอบว่า EBook implement LibraryItem interface อย่างถูกต้อง (ex6)
func TestEx6LibraryItemInterface(t *testing.T) {
	var item interfaces.LibraryItem = NewEBook("Go In Depth", "Nuttachot Promrit", 10)

	expectedInfo := "Go In Depth by Nuttachot Promrit (E-Book, 10MB)"
	info := item.Info()

	if info != expectedInfo {
		t.Errorf("Expected Info() to return '%s', got '%s'", expectedInfo, info)
	}
}
