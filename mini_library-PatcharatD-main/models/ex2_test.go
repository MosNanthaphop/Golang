package models

import (
	"testing"
)

// TestEx2NewBook ทดสอบฟังก์ชัน NewBook ว่าสร้าง Book ได้ถูกต้องหรือไม่ (ex2)
func TestEx2NewBook(t *testing.T) {
	t.Run("Create Book with valid inputs", func(t *testing.T) {
		title := "Go 101"
		author := "Nuttachot Promrit"
		stock := 3

		book := NewBook(title, author, stock)

		if book == nil {
			t.Fatalf("Expected NewBook to return a non-nil pointer, got nil")
		}

		if book.GetTitle() != title {
			t.Errorf("Expected title to be '%s', got '%s'", title, book.GetTitle())
		}

		if book.GetAuthor() != author {
			t.Errorf("Expected author to be '%s', got '%s'", author, book.GetAuthor())
		}

		if book.GetStock() != stock {
			t.Errorf("Expected stock to be %d, got %d", stock, book.GetStock())
		}
	})

	t.Run("Create Book with zero stock", func(t *testing.T) {
		title := "Zero Stock Book"
		author := "Author Zero"
		stock := 0

		book := NewBook(title, author, stock)

		if book.GetStock() != 0 {
			t.Errorf("Expected stock to be 0, got %d", book.GetStock())
		}
	})

	t.Run("Create Book with negative stock", func(t *testing.T) {
		title := "Negative Stock Book"
		author := "Author Negative"
		stock := -5

		book := NewBook(title, author, stock)

		if book.GetStock() != -5 {
			t.Errorf("Expected stock to be -5, got %d", book.GetStock())
		}
	})
}

// TestEx2GetterMethods ทดสอบ Getter Methods ว่าถูกต้องหรือไม่ (ex2)
func TestEx2GetterMethods(t *testing.T) {
	title := "Effective Go"
	author := "Robert Griesemer"
	stock := 10

	book := NewBook(title, author, stock)

	t.Run("GetTitle returns correct title", func(t *testing.T) {
		if book.GetTitle() != title {
			t.Errorf("Expected GetTitle to return '%s', got '%s'", title, book.GetTitle())
		}
	})

	t.Run("GetAuthor returns correct author", func(t *testing.T) {
		if book.GetAuthor() != author {
			t.Errorf("Expected GetAuthor to return '%s', got '%s'", author, book.GetAuthor())
		}
	})

	t.Run("GetStock returns correct stock", func(t *testing.T) {
		if book.GetStock() != stock {
			t.Errorf("Expected GetStock to return %d, got %d", stock, book.GetStock())
		}
	})
}
