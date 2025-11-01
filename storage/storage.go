package storage

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
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
	headers := []string{"ID", "Название", "Автор", "Год", "Выдана", "ID читателя"}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("не удалось записать заголовок: %w", err)
	}

	//Записываем данные книг
	for _, book := range books {
		var readerID string
		if book.ReaderID != nil {
			readerID = strconv.Itoa(*book.ReaderID)
		}
		record := []string{
			strconv.Itoa(book.ID),
			book.Title,
			book.Author,
			strconv.Itoa(book.Year),
			strconv.FormatBool(book.IsIssued),
			readerID,
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
		return nil, fmt.Errorf("не удалось открыть файл %s: %w", filename, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	//Ожидаемое количество столбцов
	var expectedColumns = 6

	//Читаем заголовок отдельно, чтобы пропустить его
	//Заодно убедимся, что файл не пустой
	if _, err := reader.Read(); err != nil {
		if errors.Is(err, io.EOF) {
			//Файл пустой или содержит только заголовок
			//Это не ошибка
			return []*domain.Book{}, nil
		}
		return nil, fmt.Errorf("не удалось прочитать заголовок: %w", err)
	}

	var books []*domain.Book
	//Добавляем счетчик строк для более информативных логов
	var lineNum int

	//Читаем построчно
	//База данных может сильно разрастись
	for {
		lineNum++
		record, err := reader.Read()
		if err != nil {
			//Если мы достигла конца файла
			//-это нормальное завершение цикла
			if errors.Is(err, io.EOF) {
				break
			}
			//Любая другая ошибка при чтении является критической
			return nil, fmt.Errorf("ошибка чтения файла на строке %d: %w", lineNum, err)
		}

		//Проверяем на точное соответствие количества колонок
		if len(record) != expectedColumns {
			log.Printf("ПРЕДУПРЕЖДЕНИЕ: Пропущена строка %d, неверное количество колонок (ожидалось %d, получено %d)", lineNum, expectedColumns, len(record))
		}

		//Ошибки будем логировать
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("ПРЕДУПРЕЖДЕНИЕ: Пропускаем строку %d, неверный формат ID: %v", lineNum, err)
			continue //Неверный формат ID, пропускаем
		}

		year, err := strconv.Atoi(record[3])
		if err != nil {
			log.Printf("ПРЕДУПРЕЖДЕНИЕ: Пропускаем строку %d, неверный формат года: %v", lineNum, err)
			continue
		}

		isIssued, err := strconv.ParseBool(record[4])
		if err != nil {
			log.Printf("ПРЕДУПРЕЖДЕНИЕ: Пропускаем строку %d, неверный формат поля 'Выдана': %v", lineNum, err)
			continue
		}

		// Если поле пустое, указатель должен быть nil. Иначе - парсим значение.
		var readerIDPtr *int
		if record[5] != "" {
			readerID, err := strconv.Atoi(record[5])
			if err != nil {
				log.Printf("ПРЕДУПРЕЖДЕНИЕ: Пропускаем строку %d, неверный формат ID читателя: %v", lineNum, err)
				continue
			}
			readerIDPtr = &readerID
		}

		book := domain.Book{
			ID:       id,
			Title:    record[1],
			Author:   record[2],
			Year:     year,
			IsIssued: isIssued,
			ReaderID: readerIDPtr, // Используем созданный указатель.
		}

		books = append(books, &book)
	}
	return books, nil
}
