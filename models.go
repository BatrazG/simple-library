package main

import "fmt"

type Book struct {
	ID       int
	Title    string
	Author   string
	IsIssued bool
}

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

//DisplayReader выводит полную информацию о пользователе
func (r Reader) DisplayReader() {
	fmt.Printf("Читатель: %s %s (ID: %d)\n", r.FirstName, r.LastName, r.ID)
}

//IssueBook выдает книгу читателю
func (b *Book) IssueBook() {
	if b.IsIssued {
		fmt.Printf("Книга %s уже кому-то выдана\n", b.Title)
		return
	}
	b.IsIssued = true
	fmt.Printf("Книга %s была выдана\n", b.Title)
}

//ReturnBook возвращает книгу в библиотеку
func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("Книга %s и так в библиотеке", b.Title)
		return
	}
	b.IsIssued = false
	fmt.Printf("Книга %s возвращена в библиотеку\n", b.Title)
}

func (u *Reader) Deactivate() {
	u.IsActive = false
}
