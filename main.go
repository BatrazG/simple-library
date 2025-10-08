package main

import "fmt"

func main() {
	user1 := Reader{
		ID:        1,
		FirstName: "Agunda",
		LastName:  "Kokoyti",
		IsActive:  true,
	}

	//user1.Deactivate()

	fmt.Println(user1)

	book1 := Book{
		ID:       1,
		Title:    "1984",
		Author:   "Джордж Оруэлл",
		Year:     1949,
		IsIssued: false,
	}

	//Выдать книгу конкретному читателю
	book1.IssueBook(&user1)
	fmt.Println(book1)
	fmt.Println("---")

	//Пробуем выдать уже выданную книгу
	reader2 := Reader{
		ID:        2,
		FirstName: "Sergey",
		LastName:  "Meniaylo",
		IsActive:  true,
	}
	book1.IssueBook(&reader2)

	//Возвращаем книгу
	book1.ReturnBook()
	fmt.Println(book1)
	fmt.Println("---")

	//Дективируем читатаеля
	user1.Deactivate()
	fmt.Println(user1)
	fmt.Println("---")

	//Пробуем выдать книгу неактивному читателю
	book1.IssueBook(&user1)

	//Интеграция уведомителя
	fmt.Println("Проверяем интеграцию уведомителя")

	// Создаем срез типа интерфейса Notifier
	notifers := []Notifer{
		EmailNotifer{EmailAdress: "a.kokoity@example.ru"},
		SMSNotifer{PhoneNumber: "799912345678"},
	}

	// Создаем сообщение
	message := "Ваша книга '1984' просрочена!"

	// Проходимся по всем уведомителям и вызываем их общий метод
	// Нам не важно, email это или sms, мы просто знаем, что они умеют уведомлять.
	for _, notifer := range notifers {
		notifer.Notify(message)
	}
}
