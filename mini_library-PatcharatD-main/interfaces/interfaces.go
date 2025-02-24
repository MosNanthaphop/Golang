package interfaces

import "fmt"

type LibraryItem interface {
	Info() string
}

func PrintItemInfo(item LibraryItem) {
	fmt.Println(item.Info())
}

type Borrowable interface {
	Borrow() error
	Return() error
}
