package models

import (
	"testing"
)

// TestEx4BookInfo ทดสอบว่า Book implement LibraryItem interface อย่างถูกต้อง (ex4)
func TestEx4BookInfo(t *testing.T) {
	book := NewBook("Go Patterns", "Sajjaporn", 5)
	expected := "Go Patterns by Sajjaporn"
	info := book.Info()
	if info != expected {
		t.Errorf("Expected Info() to return '%s', got '%s'", expected, info)
	}
}

// TestEx4DVDInfo ทดสอบว่า DVD implement LibraryItem interface อย่างถูกต้อง (ex4)
func TestEx4DVDInfo(t *testing.T) {
	dvd := NewDVD("Go Language Tour", 120, 3)
	expected := "DVD: Go Language Tour (120 mins)"
	info := dvd.Info()
	if info != expected {
		t.Errorf("Expected Info() to return '%s', got '%s'", expected, info)
	}
}

// TestEx4NewDVD ทดสอบฟังก์ชัน NewDVD ว่าสร้าง DVD ได้ถูกต้องหรือไม่ (ex4)
func TestEx4NewDVD(t *testing.T) {
	title := "Effective Go DVD"
	duration := 90
	stock := 2
	dvd := NewDVD(title, duration, stock)

	if dvd == nil {
		t.Fatalf("Expected NewDVD to return a non-nil pointer, got nil")
	}

	if dvd.Title != title {
		t.Errorf("Expected DVD.Title to be '%s', got '%s'", title, dvd.Title)
	}

	if dvd.Duration != duration {
		t.Errorf("Expected DVD.Duration to be %d, got %d", duration, dvd.Duration)
	}

	if dvd.Stock != stock {
		t.Errorf("Expected DVD.Stock to be %d, got %d", stock, dvd.Stock)
	}
}
