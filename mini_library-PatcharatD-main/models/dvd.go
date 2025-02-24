package models

import (
	"errors"
	"strconv"
)

type DVD struct {
	// TODO: implement
	Title    string
	Duration int
	Stock    int
}

func (d DVD) Info() string {
	// TODO: implement
	return "DVD: " + d.Title + " (" + strconv.Itoa(d.Duration) + " mins)"
}

func NewDVD(title string, duration int, stock int) *DVD {
	return &DVD{Title: title, Duration: duration, Stock: stock}
}

func (d *DVD) validateStock(amount int) error {
	if d.Stock+amount < 0 {
		return errors.New("stock cannot be negative")
	}
	return nil
}

func (d *DVD) AddStock(amount int) error {
	// TODO: ตรวจสอบ stock ด้วย validateStock ก่อนเพิ่ม
	if err := d.validateStock(amount); err != nil {
		return err
	} else {
		if amount > 0 {
			for i := 0; i < amount; i++ {
				increaseStock(&d.Stock)
			}
		}
		if amount <= 0 {
			d.Stock += amount
		}

	}
	return nil
}

func (d *DVD) Borrow() error {
	if d.Stock > 0 {
		d.AddStock(-1)
	} else if d.Stock == 0 {
		return errors.New("out of stock")
	}
	return nil
}

func (d *DVD) Return() error {
	d.AddStock(1)
	return nil
}
