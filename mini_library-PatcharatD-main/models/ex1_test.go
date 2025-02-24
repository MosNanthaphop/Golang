package models

import (
	"testing"
)

// TestEx1IncreaseStock ทดสอบฟังก์ชัน increaseStock (ex1)
func TestEx1IncreaseStock(t *testing.T) {
	t.Run("Increase stock from 5 to 6", func(t *testing.T) {
		stock := 5
		expected := 6

		// เรียกใช้ฟังก์ชัน increaseStock
		increaseStock(&stock)

		if stock != expected {
			t.Errorf("Expected stock to be %d, but got %d", expected, stock)
		}
	})

	t.Run("Increase stock from 0 to 1", func(t *testing.T) {
		stock := 0
		expected := 1

		// เรียกใช้ฟังก์ชัน increaseStock
		increaseStock(&stock)

		if stock != expected {
			t.Errorf("Expected stock to be %d, but got %d", expected, stock)
		}
	})

	t.Run("Increase stock from -1 to 0", func(t *testing.T) {
		stock := -1
		expected := 0

		// เรียกใช้ฟังก์ชัน increaseStock
		increaseStock(&stock)

		if stock != expected {
			t.Errorf("Expected stock to be %d, but got %d", expected, stock)
		}
	})
}
