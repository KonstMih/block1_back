package diapason

import (
	"block1_http/handler"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Get_diapason(adress string, r *http.Request) []byte {
	vars := mux.Vars(r)
	signals := vars["signals"]
	start := vars["start"]
	finish := vars["finish"]

	signals_list := strings.Split(signals, ":") // формируем список сигналов

	db := handler.Open_db(adress) // открываем базу данных
	defer db.Close()              // после всех манипуляций базу данных закрываем

	archive_data := handler.Create_byte_request(signals_list, db, start, finish)
	return archive_data
}
