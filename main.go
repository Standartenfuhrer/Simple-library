package main

import "fmt"

func main() {
	fmt.Println("Запуск системы управления библиотекой...")

	//Создаем библиотеку
	myLibrary := &Library{}

	fmt.Println("\n--- Наполняем библиотеку ---")
	//Добавляем пользователей
	myLibrary.AddReader("Тамерлан", "Джигкаев")
	myLibrary.AddReader("Линда", "Элбакянц")

	//Добавляем пользователей
	b, err := myLibrary.AddBook(1833, "Евгений Онегин", "Александр Пушкин")
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Printf("Книга (%v) успешно добавлена.", b)
	}
	b, err = myLibrary.AddBook(1984, "1984", "Джордж Оруэлл")
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Printf("\nКнига (%v) успешно добавлена.", b)
	}
	b, err = myLibrary.AddBook(1967, "Мастер и Маргарита", "Михаил Булгаков")
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Printf("\nКнига (%v) успешно добавлена.\n", b)
	}
	
	//Выдаем книгу пользователю
	fmt.Println("----Успешная выдача книги----")
	err = myLibrary.IssueBookToReader(1, 1)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(myLibrary.Books[1])
	}
	fmt.Println("----Попытка выдать уже выданную книгу----")
	err = myLibrary.IssueBookToReader(1, 1)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(myLibrary.Books[1])
	}
	fmt.Println("----Попытка выдать книгу несуществующему читателю----")
	err = myLibrary.IssueBookToReader(2, 15)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(myLibrary.Books[1])
	}
	fmt.Println("----Успешный возврат книги----")
	err = myLibrary.ReturnBook(1)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Статус книги после возварат:", myLibrary.Books[1])
	}
	fmt.Println("----Попытка вернуть книгу, которая уже в библиотеке----")
	err = myLibrary.ReturnBook(1)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Статус книги после возварат:", myLibrary.Books[1])
	}
}
