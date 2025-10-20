package main

import "fmt"

func main() {
	fmt.Println("Запуск системы управления библиотекой...")

	myLibrary := &Library{}

	fmt.Println("\n--- Наполняем библиотеку ---")
	myLibrary.AddReader("Тамерлан", "Джигкаев")
	myLibrary.AddReader("Линда", "Элбакянц")

	myLibrary.AddBook(1984, "1984", "Джордж Оруэлл")
	myLibrary.AddBook(1967, "Мастер и Маргарита", "Михаил Булгаков")

	fmt.Println("\n--- Библиотека готова к работе ---")
	fmt.Println("Количество читателей:", len(myLibrary.Readers))
	fmt.Println("Количество книг:", len(myLibrary.Books))
}
