package storage

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/BatrazG/simple-library/domain"
)

type Storable interface {
	Save() error
	Load() error
}

// Сохраняет срез книг в csv-файл
func SaveBooksToCSV(filename string, books []*domain.Book) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл %s: %w", filename, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	//Записываем заголовок
	headers := []string{"ID", "Название", "Автор"}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("не удалось записать заголовок: %w", err)
	}

	//Записываем данные книг
	for _, book := range books {
		record := []string{
			strconv.Itoa(book.ID),
			book.Title,
			book.Author,
			strconv.Itoa(book.Year),
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("не удалось записать список книги с ID %d: %w", book.ID, err)
		}
	}
	return nil
}

// LoadBooksFromCSV загружает список книг из csv
func LoadBooksFromCSV(filename string) ([]*domain.Book, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл %s", filename)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать данные из файла: %w", err)
	}

	if len(records) < 2 {
		return []*domain.Book{}, nil
	}

	var books []*domain.Book
	for i, record := range records[1:] {
		if len(record) < 4 {
			fmt.Printf("Пропускаем строку %d, недостаточно полей", i+2)
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			continue //Неверный формат ID, пропускаем
		}

		year, err := strconv.Atoi(record[3])
		if err != nil {
			continue //Неверный формат года, пропускаем
		}

		book := domain.Book{
			ID:     id,
			Title:  record[1],
			Author: record[2],
			Year:   year,
		}

		books = append(books, &book)
	}

	fmt.Println("Список загруженных книг: ")
	for _, book := range books {
		fmt.Println(book)
	}
	return books, nil

}
