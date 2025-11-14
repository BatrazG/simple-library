package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BatrazG/simple-library/cmd/cli"
	"github.com/BatrazG/simple-library/library"
	"github.com/BatrazG/simple-library/storage"
)

func main() {

	const dbFile = "books.json" // Выносим имя файла в константу

	// Пытаемся загрузить библиотеку из файла
	myLibrary, err := storage.LoadLibraryFromJSON(dbFile)
	if err != nil {
		// Если файл не найден, это не ошибка, а первый запуск.
		// Создаем новую пустую библиотеку.
		if os.IsNotExist(err) {
			fmt.Println("Файл данных не найден, создана новая библиотека.")
			myLibrary = library.New()
		} else {
			// Если произошла другая ошибка (например, JSON некорректен), завершаем работу.
			log.Fatalf("Ошибка при загрузке библиотеки: %v", err)
		}
	} else {
		fmt.Println("Библиотека успешно загружена из файла.")
	}

	// Удалите старую отладочную печать книг, если она есть
	// fmt.Println("--------------------------")
	// for i, book := range myLibrary.Books {
	// 	fmt.Println(i, book)
	// }

	cli.Run(myLibrary, dbFile)
}
