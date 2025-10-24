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
func SaveBooksToCSV(filename string, books []domain.Book) error {
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
func LoadBooksFromCSV(filename string) ([]domain.Book, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл %s:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	//Пропускаем строку заголовка
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("не удалось прочитать заголовок из файла: %w", err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Errorf("не удалось прочитать данные из файла: %w", err)
	}

	var books []domain.Book
	for i, record := range records {
		if len(record) < 4 {
			continue //если строка содержит меньше 4 полей - данные о книге неподные, пропускаем
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

		books = append(books, book)
	}
	return books, nil

}
