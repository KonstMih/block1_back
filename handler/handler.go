package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

// открытие базы данных
func Open_db(adress string) *sql.DB {
	db, err := sql.Open("sqlite3", adress)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

// преобразование количества минут в начало диапазона выборки для графика
func Minutes_to_range(db *sql.DB, minutes int) (string, string) {
	var last_date_str string
	row_last := db.QueryRow("SELECT date FROM block1 ORDER BY id DESC LIMIT 1") // запрос данных последней строки
	err := row_last.Scan(&last_date_str)
	if err != nil {
		fmt.Println(err)
	}

	last_date, err := time.Parse(time.DateTime, last_date_str)
	start_date := last_date.Add(time.Minute * time.Duration(minutes*(-1)))
	start_date_str := start_date.Format("2006-01-02 15:04:05.000")

	if err != nil {
		fmt.Println(err)
	}
	return start_date_str, last_date_str
}

// формирование хэш таблицы с данными сигнала в необходимом диапазоне
func Map_signal(db *sql.DB, signal string, start string, finish string) map[string]sql.Null[float64] {
	req_db := fmt.Sprintf("SELECT date, %s FROM block1 WHERE date > '%s' AND date < '%s'", signal, start, finish)
	column_signal, err := db.Query(req_db)
	if err != nil {
		fmt.Println(err)
	}
	defer column_signal.Close()

	var responce = make(map[string]sql.Null[float64])
	for column_signal.Next() {
		var signals_date string
		var signal_var sql.Null[float64]
		err := column_signal.Scan(&signals_date, &signal_var)
		if err != nil {
			fmt.Println(err)
		}
		responce[signals_date] = signal_var
	}
	return responce
}

// формирование массива байт из данных
func Create_byte_request(signals_list []string, db *sql.DB, start string, finish string) []byte {
	var responce_name = make(map[string](map[string]sql.Null[float64]))
	for _, signal := range signals_list {
		var responce = Map_signal(db, signal, start, finish)
		responce_name[signal] = responce
	}

	type data_response struct {
		Response map[string]map[string]sql.Null[float64] `json:"response"`
	}
	var json_response = data_response{}
	json_response.Response = responce_name
	json_byte, err := json.Marshal(json_response)
	if err != nil {
		fmt.Println(err)
	}
	return json_byte
}
