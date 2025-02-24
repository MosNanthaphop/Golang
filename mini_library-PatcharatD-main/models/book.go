package models

import (
	"errors"
	"strconv"
)

// IncreaseStock เพิ่มจำนวน stock ขึ้น 1 ผ่าน Pointer
func increaseStock(stockPtr *int) {
	*stockPtr += 1
}

// Book แทนข้อมูลหนังสือพื้นฐาน
type Book struct {
	title  string
	author string
	stock  int
}

// NewBook เป็นเหมือน Constructor Function
func NewBook(title, author string, stock int) *Book {
	// TODO: สร้าง Book และ return Pointer
	return &Book{title: title, author: author, stock: stock}
}

// Getter Methods
func (b *Book) GetTitle() string {
	// TODO: implement
	return b.title
}

func (b *Book) GetAuthor() string {
	// TODO: implement
	return b.author
}

func (b *Book) GetStock() int {
	// TODO: implement
	return b.stock
}

// validateStock ตรวจสอบความถูกต้องของจำนวน stock
func (b *Book) validateStock(amount int) error {
	if b.stock+amount < 0 {
		return errors.New("stock cannot be negative")
	}
	return nil
}

func (b Book) Details() string {
	// TODO: คืนข้อความ <Title> by <Author> (Stock: <Stock>)
	return b.title + " by " + b.author + " (Stock: " + strconv.Itoa(b.stock) + ")"
}

func (b *Book) AddStock(amount int) error {
	// TODO: ตรวจสอบ stock ด้วย validateStock ก่อนเพิ่ม
	if err := b.validateStock(amount); err != nil {
		return err
	} else {
		if amount > 0 {
			for i := 0; i < amount; i++ {
				increaseStock(&b.stock)
			}
		}
		if amount <= 0 {
			b.stock += amount
		}

	}
	return nil
}

func (b Book) Info() string {
	// TODO: implement
	return b.title + " by " + b.author
}

func (b *Book) Borrow() error {
	if b.stock > 0 {
		b.AddStock(-1)
	} else if b.stock == 0 {
		return errors.New("out of stock")
	}
	return nil
}

func (b *Book) Return() error {
	b.AddStock(1)
	return nil
}
