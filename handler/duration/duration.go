package duration

import (
	"block1_http/handler"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func Get_graph(adress string, r *http.Request) []byte {
	vars := mux.Vars(r)
	signals := vars["signals"]
	minutes, err := strconv.Atoi(vars["minutes"]) // парсим количество минут от конца таблицы, если значение не корректно по умолчанию 1 минута
	if err != nil {
		fmt.Println("minutes not a number")
		minutes = 1
	}
	signals_list := strings.Split(signals, ":") // формируем список сигналов ":"- разделитель списка сигналов

	db := handler.Open_db(adress) // открываем базу данных
	defer db.Close()              // после всех манипуляций базу данных закрываем

	start, finish := handler.Minutes_to_range(db, minutes)
	graph_data := handler.Create_byte_request(signals_list, db, start, finish)
	return graph_data
}
