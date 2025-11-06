package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Для демонстрации скопируем структуру Book сюда
type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
	ReaderID *int //ID читателя, который взял книгу
}

func main() {
	//Создаем экземпляр книги
	book := Book{
		ID:       101,
		Title:    "Война и мир",
		Author:   "Лев Толстой",
		Year:     1869,
		IsIssued: false,
		ReaderID: nil,
	}

	//Сериализуем!
	//jsonData, err := json.Marshal(book)
	jsonData, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		log.Fatalf("Ошибка при сериализации: %v", err)
	}

	//jsonData - это []byte. Преобразуем его в строку для вывода
	fmt.Println(string(jsonData))
	//Вывод: {"ID":101,"Title":"Война и мир","Author":"Лев Толстой","Year":1869,"IsIssued":false,"ReaderID":null}

	//Допустим это наш json, который мы получили
	//например из файла

	jsonString := `{
	"ID": 101,
	"Title": "Война и мир",
	"Author": "Лев Толстой",
	"Year": 1869,
	"IsIssued": false,
	"ReaderID": null
	}`

	//Преобразуем строку в срез байт
	jsonDataBytes := []byte(jsonString)

	//Создаем переменную-приемник, куда будем загружать данные
	var loadedBook Book

	//Десериализируем!
	//Обратите внимание на & - Мы передаем Указатель на переменную
	//Это нужно, чтобы функция Unmarshal могла ИЗМЕНИТЬ нашу переменную loadedBook
	err = json.Unmarshal(jsonDataBytes, &loadedBook)
	if err != nil {
		log.Fatalf("Ошибка при десериализации: %v", err)
	}

	//Проверяем результат
	fmt.Printf("Загруженная книга: %+v\n", loadedBook)
}
