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

// Library - наша центральная структура-агрегатор
type Library struct {
	Books   []*Book
	Readers []*Reader

	//Счетчики для генерации уникальных ID
	lastBookID   int
	lastReaderID int
}

func (lib *Library) AddReader(firstName, lastName string) *Reader {
	lib.lastReaderID++

	//Создаем нового читателя
	newReader := &Reader{
		ID:        lib.lastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true, //Новый читатель всегда активный
	}

	//Добавляем читателя в срез
	lib.Readers = append(lib.Readers, newReader)

	fmt.Printf("Зарегистрирован новый читатель: %s %s \n", firstName, lastName)
	return newReader
}

/*// NewLibrary Создает экземпляр новой библиотеки
func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]*Book),
		Readers: make(map[int]*Reader),
	}
}*/

// AddBook добавляет новую книгу в библиотеку
func (lib *Library) AddBook(title, author string, year int) *Book {
	lib.lastBookID++

	//Создаем новую книгу
	newBook := &Book{
		ID:       lib.lastBookID,
		Title:    title,
		Author:   author,
		Year:     year,
		IsIssued: false, //Новая книга всегда в наличии
	}

	//Добавляем новую книгу в библиотеку
	lib.Books = append(lib.Books, newBook)

	fmt.Printf("Добавлена новая книга: %s\n", newBook)
	return newBook
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
		return
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
