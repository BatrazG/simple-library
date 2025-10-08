package main

import (
	"fmt"
)

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
	ReaderID *int //ID читателя, который взял книгу
}

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

type Library struct {
	Books   map[int]*Book
	Readers map[int]*Reader
	//и т.д.
}

// DisplayReader выводит полную информацию о пользователе
func (r Reader) DisplayReader() {
	fmt.Printf("Читатель: %s %s (ID: %d)\n", r.FirstName, r.LastName, r.ID)
}

func (r Reader) String() string {
	status := ""
	if r.IsActive {
		status = "активен"
	} else {
		status = "не активен"
	}
	return fmt.Sprintf("Пользователь %s %s, № %d, пользователь %s", r.FirstName, r.LastName, r.ID, status)
}

// IssueBook выдает книгу читателю
func (b *Book) IssueBook(reader *Reader) {
	if b.IsIssued {
		fmt.Printf("Книга '%s' уже кому-то выдана\n", b.Title)
		return
	}
	if !reader.IsActive {
		fmt.Printf("Читатель %s %s не активен и не может получить книгу.", reader.FirstName, reader.LastName)
	}
	b.IsIssued = true
	b.ReaderID = &reader.ID
	fmt.Printf("Книга '%s' была выдана читателю %s %s\n", b.Title, reader.FirstName, reader.LastName)
}

// ReturnBook возвращает книгу в библиотеку
func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("Книга '%s' и так в библиотеке", b.Title)
		return
	}
	b.IsIssued = false
	b.ReaderID = nil
	fmt.Printf("Книга '%s' возвращена в библиотеку\n", b.Title)
}

// Deactivate делает пользователя неактивным
func (r *Reader) Deactivate() {
	r.IsActive = false
}

func (b Book) String() string {
	status := "в библиотеке"
	if b.IsIssued && b.ReaderID != nil {
		status = fmt.Sprintf("на руках у читателя с ID %d", *b.ReaderID)
	}
	return fmt.Sprintf("%s (%s, %d), статус %s", b.Title, b.Author, b.Year, status)
}
