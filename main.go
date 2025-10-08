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
}
