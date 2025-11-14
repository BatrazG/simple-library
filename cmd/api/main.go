package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	//Регистрируем обрабочик для пути "/"
	//Все запросы на http://localhost:8080/ будут обрабатываться функцией homeHandler
	http.HandleFunc("/", homeHandler)
	//Для лабораторной:
	http.HandleFunc("/healthcheck", healthCheckHandler)

	//Вызываем функции из ДЗ
	http.HandleFunc("/api/v1/info", infoHandler)
	http.HandleFunc("/notfound", notFoundHandler)

	//Определяем порт, на котором будет работать сервер
	port := ":8080"
	fmt.Printf("Запускаем сервер на порту %s", port)

	//Запускаем сервер
	//ListenAndServe блокирует выполнение, поэтому код дальше не пойдет,
	//Пока сервер работает или не случится ошибка
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err) //Если не удалось запустить сервер - логируем и выходим
	}
}

// homeHandler - это наш первый обработчик
// Он принимает два аргумента: ResponseWriter для записи ответа
// и Request с данными для запроса
func homeHandler(w http.ResponseWriter, r *http.Request) {
	//Просто пишем строку в тело ответа
	//w.Write() ожидает срез байт, поэтому мы преобразуем строку
	w.Write([]byte("Hello from Library API"))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	//1. Создаем структуру для ответа, чтобы ее легко было превратить в JSON
	respone := struct {
		Status      string `json:"status"`
		ProjectName string `json:"project_name"`
	}{
		Status:      "available",
		ProjectName: "LibraryAPI",
	}

	//2. Сериализуем структуру в JSON.
	jsonData, err := json.Marshal(respone)
	//Если не удалось создать JSON - это ошибка на нашей стороне
	//Отправляем статус 500 InternalServerError
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	//3. Устанавливаем заголовок, чтобы клиент знал, что мы отправляем JSON
	w.Header().Set("Content-Type", "application/json")

	//4. Устанавливаем статус-код 200 ОК
	w.WriteHeader(http.StatusOK)

	//5. Пишем JSON в тело ответа
	w.Write(jsonData)
}

// Домашнее задание

// infoHandler обрабатывает запросы к "/api/v1/info"
func infoHandler(w http.ResponseWriter, r *http.Request) {
	//Создаем структуру для ответа
	response := struct {
		Version    string `json:"version"`
		ServerTime string `json:"server_time"`
	}{
		Version:    "1.0.0",
		ServerTime: time.Now().Format(time.RFC3339),
	}

	//Сериализируем структуру в JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	//Устанавливаем заголовок
	w.Header().Set("Content-Type", "application/json")
	//Устанавливаем статус 200 ОК
	w.WriteHeader(http.StatusOK)
	//Пишем JSON в тело ответа
	w.Write(jsonData)
}

// notFoundHandler всегда возвращает ошибку 404
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	//Устанавливаем статус-код 404 Not Found
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("This page doesn't exist"))
}
