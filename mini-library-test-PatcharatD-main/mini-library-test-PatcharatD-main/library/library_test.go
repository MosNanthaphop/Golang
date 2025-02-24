// file: library_test.go
package library

import (
	"fmt"
	"testing"
)

// TestAddBook ทดสอบการเพิ่มหนังสือ (AddBook)
func TestAddBook(t *testing.T) {
	l := Library{}

	tests := []struct {
		name       string
		inputBook  Book
		wantErr    bool
		wantErrMsg string
	}{
		{
			name:      "เพิ่มหนังสือเล่มแรก (ไม่มีปัญหา)",
			inputBook: Book{ID: 1, Title: "Golang 101", Author: "Alice"},
			wantErr:   false,
		},
		{
			name:      "เพิ่มหนังสือเล่มที่ 2 (ไม่มีปัญหา)",
			inputBook: Book{ID: 2, Title: "Docker Deep Dive", Author: "Bob"},
			wantErr:   false,
		},
		{
			name:       "เพิ่มหนังสือ ID ซ้ำ (ต้อง error)",
			inputBook:  Book{ID: 1, Title: "Kubernetes Basic", Author: "Charlie"},
			wantErr:    true,
			wantErrMsg: "duplicate book ID",
		},
	}

	for _, tt := range tests {
		tt := tt // ป้องกันการ capture ผิดใน loop
		t.Run(tt.name, func(t *testing.T) {
			err := l.AddBook(tt.inputBook)
			if tt.wantErr && err == nil {
				t.Errorf("AddBook(%v) expected error but got nil", tt.inputBook)
				return
			}
			if !tt.wantErr && err != nil {
				t.Errorf("AddBook(%v) unexpected error: %v", tt.inputBook, err)
				return
			}
			if tt.wantErr {
				// เช็กว่า error message ตรงตามที่คาดไหม
				if err.Error() != tt.wantErrMsg {
					t.Errorf("AddBook(%v) error msg = %q, want %q",
						tt.inputBook, err.Error(), tt.wantErrMsg)
				}
			}
		})
	}

	if got, want := len(l.Books), 2; got != want {
		t.Fatalf("Library has %d books, want %d", got, want)
	}
}

// TestGetBookByID ทดสอบการค้นหาหนังสือ (GetBookByID)
func TestGetBookByID(t *testing.T) {
	// เตรียม library พร้อมหนังสือ 2 เล่ม
	l := Library{
		Books: []Book{
			{ID: 10, Title: "Microservices in Go", Author: "Dan"},
			{ID: 20, Title: "Clean Code", Author: "Bob Martin"},
		},
	}

	// ใช้ Subtests แยกแต่ละเคส
	t.Run("ค้นหาเล่มที่มีอยู่ (ID=10)", func(t *testing.T) {
		got, err := l.GetBookByID(10)
		if err != nil {
			t.Errorf("GetBookByID(10) unexpected error: %v", err)
			return
		}
		if got.ID != 10 {
			t.Errorf("GetBookByID(10) got.ID = %d, want 10", got.ID)
		}
		if got.Title != "Microservices in Go" {
			t.Errorf("GetBookByID(10) got.Title = %q, want %q",
				got.Title, "Microservices in Go")
		}
	})

	t.Run("ค้นหาเล่มที่มีอยู่ (ID=20)", func(t *testing.T) {
		got, err := l.GetBookByID(20)
		if err != nil {
			t.Errorf("GetBookByID(20) unexpected error: %v", err)
			return
		}
		if got.Title != "Clean Code" {
			t.Errorf("GetBookByID(20) got.Title = %q, want %q",
				got.Title, "Clean Code")
		}
	})

	t.Run("ค้นหาเล่มที่ไม่มีอยู่ (ID=999)", func(t *testing.T) {
		_, err := l.GetBookByID(999)
		if err == nil {
			t.Errorf("GetBookByID(999) expected error but got nil")
		} else {
			// ตรวจสอบ Error Message หรือ type ได้
			if err.Error() != "book not found" {
				t.Errorf("GetBookByID(999) error = %q, want %q",
					err.Error(), "book not found")
			}
		}
	})
}

// TestDescribe ทดสอบการ implement Interface Describable ของ Book
func TestDescribe(t *testing.T) {
	// สมมติ Book implement func (b Book) Describe() string
	b := Book{ID: 100, Title: "Interface in Go", Author: "Jane"}

	// ตรวจสอบว่า Book ตรงตาม Describable interface ได้
	var d Describable = b // ถ้า Book ไม่มี Describe() ก็จะคอมไพล์ Error

	desc := d.Describe()
	wantSubstring := "Book[ID=100"
	if desc == "" {
		t.Errorf("Describe() got empty string, want something meaningful")
	}
	if !contains(desc, wantSubstring) {
		t.Errorf("Describe() = %q, want substring %q", desc, wantSubstring)
	}
}

// ฟังก์ชันช่วยตรวจว่าตัว s มี substring sub หรือไม่
func contains(s, sub string) bool {
	return len(s) >= len(sub) && (fmt.Sprintf("%v", s[0:len(sub)]) == sub ||
		func() bool {
			for i := 0; i+len(sub) <= len(s); i++ {
				if s[i:i+len(sub)] == sub {
					return true
				}
			}
			return false
		}())
}
