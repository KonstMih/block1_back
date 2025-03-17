package main

import (
	"block1_http/handler/diapason"
	"block1_http/handler/duration"
	"block1_http/row"
	"fmt"
	"log"
	"net/http"

  "github.com/rs/cors"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/last", LastRow)                               // запрос последней строки б/д
	r.HandleFunc("/duration/{signals}/{minutes}", Graf)          // запрос данных столбца для построения графика
	r.HandleFunc("/archive/{signals}/{start}/{finish}", Archive) // запрос данных столбца в диапазоне времени для архива

  // Настраиваем CORS middleware
  c := cors.New(cors.Options{
    AllowedOrigins:   []string{"*"}, // Разрешенные домены
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Разрешенные методы
    AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization"}, // Разрешенные заголовки
    AllowCredentials: true, // Разрешить передачу куки и заголовков авторизации
    Debug:            true, // Включить отладку (опционально)
  })

  // Обертываем роутер в CORS middleware
  handler := c.Handler(r)

	http.Handle("/", handler)

	fmt.Println("Start server...")
	log.Println(http.ListenAndServe(":9090", nil)) // номер порта на котором был запущен сервер

}

// функция запроса последней строки
func LastRow(w http.ResponseWriter, r *http.Request) {
	json_last_row := row.Get_last_row("./data.db")

	_, err := w.Write(json_last_row)
	if err != nil {
		fmt.Println(err)
	}
}

// функция запроса данных столбца для построения графика
func Graf(w http.ResponseWriter, r *http.Request) {
	json_column := duration.Get_graph("./data.db", r)

	_, err := w.Write(json_column)
	if err != nil {
		fmt.Println(err)
	}
}

// запрос данных столбца в диапазоне времени для архива
func Archive(w http.ResponseWriter, r *http.Request) {
	json_column := diapason.Get_diapason("./data.db", r)

	_, err := w.Write(json_column)
	if err != nil {
		fmt.Println(err)
	}
}

