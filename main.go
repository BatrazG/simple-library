package main

import "fmt"

func main() {
	user1 := Reader{
		ID:        1,
		FirstName: "Agunda",
		LastName:  "Kokoyti",
		IsActive:  true,
	}

	user1.Deactivate()

	fmt.Println(user1)

	book1 := Book{
		ID:       1,
		Title:    "1984",
		Author:   "Джордж Оруэлл",
		Year:     1949,
		IsIssued: false,
	}

	fmt.Println(book1)
	book1.IssueBook(&user1)
	fmt.Println(book1)
	book1.ReturnBook()
	fmt.Println(book1)
}
