package model

import (
	"database/sql"
	"example.com/projectApiClient"
	"fmt"
)

// ModuleModel используется для конструктора модели
type ModuleModel struct {
	dataBase *sql.DB
}

// NewModuleModel конструктор модели возвращающий указатель на структуру ModuleModel
func NewModuleModel(DB *sql.DB) *ModuleModel {
	return &ModuleModel{
		dataBase: DB,
	}
}

// GetModuleById метод модели по получению всех пользователей из БД возвращает массив структур Module по id документа и ошибку
func (m *ModuleModel) GetModuleById(documentId int64) ([]projectApiClient.Module, error) {
	fmt.Println(
		"m.dataBase",
		m.dataBase,
	)
	//QueryRow запрос возврата сроки выборки из таблицы значений значений по id
	var rows, err = m.dataBase.Query(
		"SELECT id, title FROM documentations.module where fk_document=$1",
		documentId,
	)
	if err != nil {
		err = fmt.Errorf(
			"Ошибка в выбора таблицы %s",
			err,
		)
		return nil, err
	}
	defer rows.Close()
	module := []projectApiClient.Module{}
	//получение данных из всей таблицы
	for rows.Next() {
		p := projectApiClient.Module{}
		err := rows.Scan(
			&p.Id,
			&p.Title,
		)
		if err != nil {
			err = fmt.Errorf(
				"Ошибка сканирования результата селекта %s",
				err,
			)
			return nil, err
		}
		//добавление новых данных в массив структур
		module = append(
			module,
			p,
		)
	}
	return module, err
}
