package library

import (
	"errors"
	"fmt"
)

type Book struct {
	ID     int
	Title  string
	Author string
}

type Library struct {
	Books []Book
}

func (l *Library) AddBook(b Book) error {
	for _, book := range l.Books {
		if book.ID == b.ID {
			return errors.New("duplicate book ID")
		}
	}
	l.Books = append(l.Books, b)
	return nil
}

func (l *Library) GetBookByID(id int) (Book, error) {
	for _, book := range l.Books {
		if book.ID == id {
			return book, nil
		}
	}
	return Book{}, errors.New("book not found")
}

type Describable interface {
	Describe() string
}

func (b Book) Describe() string {
	return fmt.Sprintf("Book[ID=%d, Title=%s, Author=%s]", b.ID, b.Title, b.Author)
}
