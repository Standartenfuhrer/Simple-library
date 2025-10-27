package main

import (
	"fmt"
	//"github.com/Standartenfuhrer/simple-library/domain"
	"github.com/Standartenfuhrer/simple-library/library"
)

func main() {
	myLibrary := library.New()

	reader, err := myLibrary.AddReader("Тамерлан", "Джигкаев")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Зарегестрирован новый читатель:", reader)
	}

	b, err := myLibrary.AddBook(1833, "Егвений Онегин", "Александр Пушкин")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Книга (%v) успешно добавлена\n", b.Title)
	}

	err = myLibrary.IssueBookToReader(1, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(myLibrary.Books[0])
	}
}
