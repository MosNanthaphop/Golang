package models

import (
	"testing"
)

// TestEx3Details ทดสอบเมธอด Details() (ex3)
func TestEx3Details(t *testing.T) {
	t.Run("Details returns correct string", func(t *testing.T) {
		title := "Golang Concurrency"
		author := "Bob"
		stock := 2

		book := NewBook(title, author, stock)
		expected := "Golang Concurrency by Bob (Stock: 2)"
		details := book.Details()

		if details != expected {
			t.Errorf("Expected Details to return '%s', got '%s'", expected, details)
		}
	})

	t.Run("Details after stock change", func(t *testing.T) {
		title := "Advanced Go"
		author := "Alice"
		stock := 5

		book := NewBook(title, author, stock)
		expectedInitial := "Advanced Go by Alice (Stock: 5)"
		if book.Details() != expectedInitial {
			t.Errorf("Expected initial Details to be '%s', got '%s'", expectedInitial, book.Details())
		}

		// เพิ่ม stock
		if err := book.AddStock(3); err != nil {
			t.Errorf("Unexpected error when adding stock: %v", err)
		}
		expectedAfterAdd := "Advanced Go by Alice (Stock: 8)"
		if book.Details() != expectedAfterAdd {
			t.Errorf("Expected Details after adding stock to be '%s', got '%s'", expectedAfterAdd, book.Details())
		}

		// ลด stock
		if err := book.AddStock(-2); err != nil {
			t.Errorf("Unexpected error when subtracting stock: %v", err)
		}
		expectedAfterSubtract := "Advanced Go by Alice (Stock: 6)"
		if book.Details() != expectedAfterSubtract {
			t.Errorf("Expected Details after subtracting stock to be '%s', got '%s'", expectedAfterSubtract, book.Details())
		}
	})
}

// TestEx3AddStock ทดสอบเมธอด AddStock(amount int) error (ex3)
func TestEx3AddStock(t *testing.T) {
	t.Run("Add positive stock", func(t *testing.T) {
		book := NewBook("Test Book", "Tester", 10)
		err := book.AddStock(5)
		if err != nil {
			t.Errorf("Expected no error when adding stock, got %v", err)
		}
		expected := 15
		if book.GetStock() != expected {
			t.Errorf("Expected stock to be %d, got %d", expected, book.GetStock())
		}
	})

	t.Run("Subtract stock without going negative", func(t *testing.T) {
		book := NewBook("Test Book", "Tester", 10)
		err := book.AddStock(-3)
		if err != nil {
			t.Errorf("Expected no error when subtracting stock, got %v", err)
		}
		expected := 7
		if book.GetStock() != expected {
			t.Errorf("Expected stock to be %d, got %d", expected, book.GetStock())
		}
	})

	t.Run("Subtract stock resulting in zero", func(t *testing.T) {
		book := NewBook("Test Book", "Tester", 3)
		err := book.AddStock(-3)
		if err != nil {
			t.Errorf("Expected no error when subtracting stock to zero, got %v", err)
		}
		expected := 0
		if book.GetStock() != expected {
			t.Errorf("Expected stock to be %d, got %d", expected, book.GetStock())
		}
	})

	t.Run("Subtract stock resulting in negative", func(t *testing.T) {
		book := NewBook("Test Book", "Tester", 2)
		err := book.AddStock(-5)
		if err == nil {
			t.Errorf("Expected error when subtracting stock resulting in negative, got nil")
		} else {
			expectedError := "stock cannot be negative"
			if err.Error() != expectedError {
				t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
			}
		}
		// ตรวจสอบว่า stock ไม่เปลี่ยนแปลง
		expected := 2
		if book.GetStock() != expected {
			t.Errorf("Expected stock to remain %d, got %d", expected, book.GetStock())
		}
	})

	t.Run("Add zero stock", func(t *testing.T) {
		book := NewBook("Test Book", "Tester", 5)
		err := book.AddStock(0)
		if err != nil {
			t.Errorf("Expected no error when adding zero stock, got %v", err)
		}
		expected := 5
		if book.GetStock() != expected {
			t.Errorf("Expected stock to remain %d, got %d", expected, book.GetStock())
		}
	})
}
