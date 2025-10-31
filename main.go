package main

import (
	"fmt"

	"github.com/Standartenfuhrer/simple-library/domain"
	"github.com/Standartenfuhrer/simple-library/library"
	"github.com/Standartenfuhrer/simple-library/storage"
)

func main() {
	myLibrary := library.New()
	book := &domain.Book{}
	reader := &domain.Reader{}

	reader, err := myLibrary.AddReader("Тамерлан", "Джигкаев")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Зарегестрирован новый читатель:", reader)
	}

	book, err = myLibrary.AddBook(1833, "Егвений Онегин", "Александр Пушкин")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Книга (%v) успешно добавлена\n", book.Title)
	}

	err = myLibrary.IssueBookToReader(1, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(myLibrary.Books[0])
	}

	err = storage.SaveBooksToCSV("test.csv", )
}
