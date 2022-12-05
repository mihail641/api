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

// ModuleController структура используется для конструктора контроллер
type ModuleController struct {
	model *model.ModuleModel
}

// NewModuleModel конструктор контроллера, возращающий экземпляр структуры ModuleController
// со свойством model контроллера модели с аргументом DB
func NewModuleController(DB *sql.DB) *ModuleController {
	return &ModuleController{
		model: model.NewModuleModel(DB),
	}
}

// GetModuleById метод контроллера по получения значений из БД по id документа
func (dc *ModuleController) GetModuleById(res http.ResponseWriter, req *http.Request) {
	//установливаем заголовок «Content-Type: application/json», т.к. потому что мы отправляем данные JSON с запросом через роутер
	res.Header().Set(
		"Content-Type",
		"application/json",
	)
	//изъятия из заголовка URL id string
	params := mux.Vars(req) // we are extracting 'id' of the Course which we are passing in the url
	var id = params["id"]
	//конвертация string в int
	//парсинг id из строки в int 64
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
	//передача парметра id методу модели GetModuleById
	p, err := dc.model.GetModuleById(s)
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
