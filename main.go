package main

import (
	"fmt"

	//"github.com/Standartenfuhrer/simple-library/domain"
	"github.com/Standartenfuhrer/simple-library/cmd/cli"
	"github.com/Standartenfuhrer/simple-library/library"
	"github.com/Standartenfuhrer/simple-library/storage"
)

func main() {
	//Создаем пустые экземпляры наших структур
	myLibrary := library.New()

	//Создаем пустую ошибку
	var err error

	//Загружаем список читателей и выводим его в случае успеха
	myLibrary.Readers, err = storage.LoadReadersFromCSV("readers.csv")
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(myLibrary.Readers)
	}

	//Загружаем список книг и выводим его в случае успеха
	myLibrary.Books, err = storage.LoadBooksFromCSV("books.csv")
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(myLibrary.Books)
	}

	err = myLibrary.ReturnBook(1)
	if err != nil{
		fmt.Println(err)
	}

	//Запускаем наш проект
	cli.Run(myLibrary)

}
