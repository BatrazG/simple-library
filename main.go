package main

import (
	"fmt"

	//"github.com/BatrazG/simple-library/cli"
	"github.com/BatrazG/simple-library/cmd/cli"
	"github.com/BatrazG/simple-library/library"
)

func main() {

	myLibrary := library.New()

	myLibrary.AddBook("1984", "Джордж Оруэлл", 1949)
	myLibrary.AddBook("Мастер и Маргарита", "Михаил Булгаков", 1967)
	myLibrary.AddBook("мастер и маргарита", "Михаил Булгаков", 1998)

	_, err := myLibrary.AddReader("Агунда", "Кокойти")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Читатель успешно добавлен")
	}
	_, err = myLibrary.AddReader("Сергей", "Меняйло")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Читатель успешно добавлен")
	}

	allBooks := myLibrary.GetAllBooks()
	for _, book := range allBooks {
		fmt.Println(book)

		fmt.Println("Ищем книгу по названию 'мастер и маргарита'")
		foundBooks, err := myLibrary.FindBookByTitle("мастер и маргарита")
		if err != nil {
			fmt.Println("Неудачный поиск:", err)
		} else {
			for i, book := range foundBooks {
				fmt.Println(i+1, book)
			}
		}
	}

	cli.Run(myLibrary)

}

/*fmt.Println("Запуск системы управления библиотекой")

//1. Создаем экземпляр библиотеки
myLibrary := &Library{} //Пустая библиотека готова к работе

fmt.Println("Наполняем библиотеку")
//2. Добавляем читателей
_, err := myLibrary.AddReader("Агунда", "Кокойти")
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Читатель успешно добавлен")
}
_, err = myLibrary.AddReader("Сергей", "Меняйло")
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Читатель успешно добавлен")
}

//3. Добавляем книги
myLibrary.AddBook("1984", "Джордж Оруэлл", 1949)
myLibrary.AddBook("Мастер и Маргарита", "Михаил Булгаков", 1967)

fmt.Println("\n---Библиотека готова к работе---")
fmt.Println("Количество читателей:", len(myLibrary.Readers))
fmt.Println("Количество книг:", len(myLibrary.Books))

//Модуль 16. Практикум
fmt.Println("---Тестируем выдачу книг---")
//Выдаем книгу 1 читателю 1
fmt.Println("Попытка выдать книгу 1 читателю 1")
err = myLibrary.IssueBookToReader(1, 1)
if err != nil {
	fmt.Println("Ошибка выдачи:", err)
} else {
	fmt.Println("Книга успешно выдана")
}

//Попытка выдать ту же книгу еще раз
fmt.Println("Поптка выдать уже выданную книгу")
err = myLibrary.IssueBookToReader(1, 2)
if err != nil {
	fmt.Println("Ошибка выдачи:", err)
} else {
	fmt.Println("Книга успешно выдана")
}

fmt.Println("Попытка выдать несуществующую книгу")
err = myLibrary.IssueBookToReader(99, 1)
if err != nil {
	fmt.Println("Ошибка выдачи:", err)
} else {
	fmt.Println("Книга успешно выдана")
}

fmt.Println("Попытка выдать книгу несуществующему читателю")
err = myLibrary.IssueBookToReader(2, 99)
if err != nil {
	fmt.Println("Ошибка выдачи:", err)
} else {
	fmt.Println("Книга успешно выдана")
}

fmt.Println()

//Смотрим все книги в библиотеке
//myLibrary.ListAllBooks()
//Рефактторинг 5
fmt.Println("------")
fmt.Println("Выводим список книг с помощью универсального метода:")
allBooks := myLibrary.GetAllBooks()
if len(allBooks) == 0 {
	fmt.Println("Библиотека пуста")
} else {
	for i, book := range allBooks {
		fmt.Printf("%d: %s\n", i+1, book)
	}
}
fmt.Println("------\n")

//Тест возврата книги в библиотеку с помощью метода-дирижера Library.ReturnBook
//Возвращаем успешно выданную книгу в библиотеку
fmt.Println("Тест возврата книг")
fmt.Println(myLibrary.Books[0])
err = myLibrary.ReturnBook(1)
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Книга успешно возвращена в библиотеку")
	fmt.Println(myLibrary.Books[0])
}
//Тест попытки еще раз вернуть книгу, которая уже в библиотеке
err = myLibrary.ReturnBook(1)
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Книга успешно возвращена в библиотеку")
	fmt.Println(myLibrary.Books[0])
}

//------------------
//Тестируем config
fmt.Println("------")
fmt.Println("Поиск порта")
config := map[string]string{
	"PORT": "456",
}

port, err := GetPortFromConfig(config)
if err != nil {
	fmt.Println("Ошибка:", err)
} else {
	fmt.Println(port)
}

config = map[string]string{
	"tort": "medivik",
}
port, err = GetPortFromConfig(config)
if err != nil {
	fmt.Println("Ошибка:", err)
} else {
	fmt.Println(port)
}
*/
