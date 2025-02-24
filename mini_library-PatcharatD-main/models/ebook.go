package models

import "strconv"

type EBook struct {
	// TODO: implement
	Book
	fileSizeMB int
}

func NewEBook(title, author string, fileSizeMB int) *EBook {
	// TODO: implement
	return &EBook{Book: Book{title: title, author: author}, fileSizeMB: fileSizeMB}
}

func (eb *EBook) GetFileSize() int {
	// TODO: implement
	return eb.fileSizeMB
}

func (eb *EBook) Info() string {
	// TODO: implement using getter methods
	return eb.GetTitle() + " by " + eb.GetAuthor() + " (E-Book, " + strconv.Itoa(eb.GetFileSize()) + "MB)"
}
