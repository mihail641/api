package controller

import (
	"encoding/json"
	"fmt"
	"my_project/model"
	"net/http"
)

// DirectoriesController труктура используется для конструктора контроллер
type DirectoriesController struct {
	model *model.DirectoriesModel
}

// NewDirectoriesController конструктор контроллера, возращающий экземпляр структуры DirectoriesController
// со свойством model контроллера модели с аргументом DB
func NewDirectoriesController() *DirectoriesController {
	return &DirectoriesController{
		model: model.NewDirectoriesModel(),
	}
}

// GetDirectories метод контроллера по получению  значений Директорий
func (dir *DirectoriesController) GetDirectories(res http.ResponseWriter, req *http.Request) {
	//установливаем заголовок «Content-Type: application/json», т.к. потому что мы отправляем данные JSON с запросом через роутер
	res.Header().Set("Content-Type", "application/json")
	path := "/home/kate/Music"
	fmt.Println("path", path)
	//обращение к методу модели GetDirectories
	directories, err := dir.model.GetDirectories(path)
	if err != nil {
		m := "Ошибка выполнеия функции получения информации о всех пользователях: %s"
		fmt.Println(m, err)
		fmt.Fprintf(res, m, err)
		return
	}
	//кодирование в json результата выполнения метода и передача в пакет main
	json.NewEncoder(res).Encode(&directories)
}
