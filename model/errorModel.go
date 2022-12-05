package model

import (
	"database/sql"
	"fmt"
)

// User структура используется инициализации данные в структуры
type Error struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

// UserModel используется для конструктора модели
type ErrorModel struct {
	dataBase *sql.DB
}

// NewUserModel конструктор модели возвращающий указатель на структуру ErrorModel
func NewErrorModel(DB *sql.DB) *ErrorModel {
	return &ErrorModel{
		dataBase: DB,
	}
}

// GetErrorById метод модели по получению всех пользователей из БД возвращает массив структур Error и ошибку
func (m *ErrorModel) GetErrorById(moduleId int64) ([]Error, error) {
	//Query запрос возврата срок выборки из таблицы значений значений по id
	var rows, err = m.dataBase.Query(
		"SELECT id, title FROM documentations.error where fk_module=$1",
		moduleId,
	)
	if err != nil {
		err := fmt.Errorf(
			"Ошибка в выбора таблицы %s ",
			err,
		)
		return nil, err
	}
	defer rows.Close()
	error := []Error{}
	//получение данных из всей таблицы
	for rows.Next() {
		p := Error{}
		err := rows.Scan(
			&p.Id,
			&p.Title,
		)
		if err != nil {
			err := fmt.Errorf(
				"Ошибка сканирования результата селекта %s ",
				err,
			)
			return nil, err
		}
		//добавление новых данных в массив структур
		error = append(
			error,
			p,
		)
	}
	return error, err
}
