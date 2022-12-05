package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"my_project/model"
	"net/http"
	"strconv"
)

// ErrorController структура используется для конструктора контроллер
type ErrorController struct {
	model *model.ErrorModel
}

// NewErrorModel конструктор контроллера, возращающий экземпляр структуры ErrorController
// со свойством users контроллера модели с аргументом DB
func NewErrorController(DB *sql.DB) *ErrorController {
	return &ErrorController{
		model: model.NewErrorModel(DB),
	}
}

// GetErrorById метод контроллера по получModuения всех значений из БД
func (dc *ErrorController) GetErrorById(res http.ResponseWriter, req *http.Request) {
	//установливаем заголовок «Content-Type: application/json», т.к. потому что мы отправляем данные JSON с запросом через роутер
	res.Header().Set(
		"Content-Type",
		"application/json",
	)
	//изъятия из заголовка URL id string
	params := mux.Vars(req) // we are extracting 'id' of the Course which we are passing in the url

	var id = params["id"]
	//конвертация string в int
	s, err := strconv.ParseInt(
		id,
		10,
		64,
	)
	if err != nil {
		m := "Ошибка перевода id из string в int64 "
		fmt.Println(
			m,
			err,
		)
		fmt.Fprintf(
			res,
			m,
			err,
		)
		return
	}
	//передача парметра id методу модели GetErrorById
	p, err := dc.model.GetErrorById(s)
	if err != nil {
		m := "Ошибка выполнения функции выбора по id: "
		fmt.Println(
			m,
			err,
		)
		fmt.Fprintf(
			res,
			m,
			err,
		)
		return
	}
	//кодирование в json результата выполнения метода и передача в пакет main
	json.NewEncoder(res).Encode(&p)
}
