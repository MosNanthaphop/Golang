package models

import (
	"mini_library/interfaces"
	"testing"
)

// TestEx5BookBorrow ทดสอบเมธอด Borrow() สำหรับ Book (ex5)
func TestEx5BookBorrow(t *testing.T) {
	t.Run("Borrow when stock > 0", func(t *testing.T) {
		book := NewBook("Intro to Go", "Alice", 2)
		err := book.Borrow()
		if err != nil {
			t.Errorf("Expected no error when borrowing, got %v", err)
		}
		expectedStock := 1
		if book.GetStock() != expectedStock {
			t.Errorf("Expected stock to be %d, got %d", expectedStock, book.GetStock())
		}
	})

	t.Run("Borrow when stock == 0", func(t *testing.T) {
		book := NewBook("Advanced Go", "Bob", 0)
		err := book.Borrow()
		if err == nil {
			t.Errorf("Expected error when borrowing with stock 0, got nil")
		} else {
			expectedError := "out of stock"
			if err.Error() != expectedError {
				t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
			}
		}
		if book.GetStock() != 0 {
			t.Errorf("Expected stock to remain 0, got %d", book.GetStock())
		}
	})
}

// TestEx5BookReturn ทดสอบเมธอด Return() สำหรับ Book (ex5)
func TestEx5BookReturn(t *testing.T) {
	t.Run("Return a borrowed book", func(t *testing.T) {
		book := NewBook("Intro to Go", "Alice", 1)
		err := book.Borrow()
		if err != nil {
			t.Errorf("Expected no error when borrowing, got %v", err)
		}
		err = book.Return()
		if err != nil {
			t.Errorf("Expected no error when returning, got %v", err)
		}
		expectedStock := 1
		if book.GetStock() != expectedStock {
			t.Errorf("Expected stock to be %d after return, got %d", expectedStock, book.GetStock())
		}
	})
}

// TestEx5DVDBorrow ทดสอบเมธอด Borrow() สำหรับ DVD (ex5)
func TestEx5DVDBorrow(t *testing.T) {
	t.Run("Borrow when stock > 0", func(t *testing.T) {
		dvd := NewDVD("Go Language Tour", 120, 3)
		err := dvd.Borrow()
		if err != nil {
			t.Errorf("Expected no error when borrowing, got %v", err)
		}
		expectedStock := 2
		if dvd.Stock != expectedStock {
			t.Errorf("Expected stock to be %d, got %d", expectedStock, dvd.Stock)
		}
	})

	t.Run("Borrow when stock == 0", func(t *testing.T) {
		dvd := NewDVD("Go Concurrency", 90, 0)
		err := dvd.Borrow()
		if err == nil {
			t.Errorf("Expected error when borrowing with stock 0, got nil")
		} else {
			expectedError := "out of stock"
			if err.Error() != expectedError {
				t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
			}
		}
		if dvd.Stock != 0 {
			t.Errorf("Expected stock to remain 0, got %d", dvd.Stock)
		}
	})
}

// TestEx5DVDReturn ทดสอบเมธอด Return() สำหรับ DVD (ex5)
func TestEx5DVDReturn(t *testing.T) {
	t.Run("Return a borrowed DVD", func(t *testing.T) {
		dvd := NewDVD("Go Language Tour", 120, 2)
		err := dvd.Borrow()
		if err != nil {
			t.Errorf("Expected no error when borrowing, got %v", err)
		}
		err = dvd.Return()
		if err != nil {
			t.Errorf("Expected no error when returning, got %v", err)
		}
		expectedStock := 2
		if dvd.Stock != expectedStock {
			t.Errorf("Expected stock to be %d after return, got %d", expectedStock, dvd.Stock)
		}
	})
}

// TestEx5BorrowableInterface ตรวจสอบว่า Book และ DVD implement Borrowable interface (ex5)
func TestEx5BorrowableInterface(t *testing.T) {
	var _ interfaces.Borrowable = (*Book)(nil)
	var _ interfaces.Borrowable = (*DVD)(nil)
}

// TestEx5ErrorHandling ทดสอบการจัดการข้อผิดพลาดใน Borrow() และ Return() (ex5)
func TestEx5ErrorHandling(t *testing.T) {
	t.Run("Borrow reduces stock correctly", func(t *testing.T) {
		book := NewBook("Effective Go", "Robert Griesemer", 1)
		err := book.Borrow()
		if err != nil {
			t.Errorf("Expected no error when borrowing, got %v", err)
		}
		if book.GetStock() != 0 {
			t.Errorf("Expected stock to be 0 after borrowing, got %d", book.GetStock())
		}
	})

	t.Run("Return increases stock correctly", func(t *testing.T) {
		dvd := NewDVD("Go DVD", 90, 1)
		err := dvd.Borrow()
		if err != nil {
			t.Errorf("Expected no error when borrowing, got %v", err)
		}
		err = dvd.Return()
		if err != nil {
			t.Errorf("Expected no error when returning, got %v", err)
		}
		if dvd.Stock != 1 {
			t.Errorf("Expected stock to be 1 after returning, got %d", dvd.Stock)
		}
	})

	t.Run("Borrowing with insufficient stock returns error", func(t *testing.T) {
		book := NewBook("Go Programming", "Jane Doe", 0)
		err := book.Borrow()
		if err == nil {
			t.Errorf("Expected error when borrowing with stock 0, got nil")
		} else if err.Error() != "out of stock" {
			expectedError := "Out of stock!"
			if err.Error() != expectedError {
				t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
			}
		}
	})
}
