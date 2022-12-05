// Package model, извлечение из БД, выполнение манипуляций с БД
package model

import (
	"database/sql"
	"fmt"
)

// User структура используется инициализации данные в структуры
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Sale int    `json:"sale"`
}

// UserModel используется для конструктора модели
type UserModel struct {
	dataBase *sql.DB
}

// NewUserModel конструктор модели возвращающий указатель на структуру UserModel
func NewUserModel(DB *sql.DB) *UserModel {
	return &UserModel{
		dataBase: DB,
	}
}

// Getusers метод модели по получению всех пользователей из БД возвращает массив
//структур User и ошибку
func (m *UserModel) Getusers() ([]User, error) {
	//rows запрос возврата сроки выборки из таблицы значений
	var rows, err = m.dataBase.Query("SELECT id, name, sale FROM Misha2")
	if err != nil {
		err = fmt.Errorf(
			"Ошибка в выбора таблицы %s",
			err,
		)
		return nil, err
	}
	defer rows.Close()
	//users инициализация массива структур User
	users := []User{}
	//получение данных из всей таблицы
	for rows.Next() {
		p := User{}
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Sale,
		)
		if err != nil {
			err = fmt.Errorf(
				"Ошибка сканирования результата селекта %s",
				err,
			)
			return nil, err
		}
		//добавление новых данных в массив структур
		users = append(
			users,
			p,
		)
	}
	//возврат массива структур и ошибки
	return users, err
}

// GetSingleUser метод модели по получению всех пользователей из БД по id, возвращает структуру User и ошибку
func (m *UserModel) GetSingleUser(id int) (User, error) {
	var p User
	//QueryRow запрос возврата сроки выборки из таблицы значений значений по id
	row1 := m.dataBase.QueryRow(
		"SELECT id, name, sale FROM Misha2 where id=$1",
		id,
	)
	//row1.Scan сканирование полученных результатов
	err := row1.Scan(
		&p.ID,
		&p.Name,
		&p.Sale,
	)
	if err == sql.ErrNoRows {
		err = fmt.Errorf(
			"Нет такого id=%d",
			id,
		)
		return p, err
	}

	return p, err
}

// CreateUser метод модели по созданию нового пользователя используя уникальный id возвращает структу User и ошибку
func (m *UserModel) CreateUser(name string, sale int) (User, error) {
	//user инициализация данных переданных из контроллера в структуру User
	var user = User{Name: name, Sale: sale}
	//создание нового значения используя данные переданных из контроллера с уникальным id
	err := m.dataBase.QueryRow(
		"INSERT INTO Misha2 (name, sale) VALUES($1,$2) returning id",
		user.Name,
		user.Sale,
	).Scan(&user.ID)
	if err != nil {
		err = fmt.Errorf(
			"Ошибка создания строки %s",
			err,
		)
		return user, err
	}
	return user, nil
}

// UpdateUser метод модели по изменению конкретного пользователя из БД, возвращает структуру User и ошибку
func (m *UserModel) UpdateUser(id int, name string, sale int) (User, error) {
	//инициализация данных переданных из контроллера в структуру User
	user := User{ID: id, Name: name, Sale: sale}
	fmt.Println(
		"Печать из модели",
		id,
		name,
		sale,
	)
	//изменения в БД значений при совпадении id
	row1 := m.dataBase.QueryRow(
		"SELECT id FROM Misha2 where id=$1",
		user.ID,
	)
	err := row1.Scan(&user.ID)
	if err == sql.ErrNoRows {
		err = fmt.Errorf(
			"Нет такого id=%d",
			id,
		)
		return user, err
	} else {
		_, err := m.dataBase.Exec(
			"update Misha2 set name =$1, sale= $2 where id = $3",
			user.Name,
			user.Sale,
			user.ID,
		)
		if err != nil {
			err = fmt.Errorf(
				"Нет такого id=%d",
				user.ID,
			)
		}
		return user, err
	}
}

// DeleteUser метод модели по удалению конкретного пользователя из БД по id, возвращает структуру User и ошибку
func (m *UserModel) DeleteUser(id int) (User, error) {
	var s *string
	//row1 нахождение строки с id переданного из контроллера
	row1 := m.dataBase.QueryRow(
		"SELECT FROM Misha2 where id=$1",
		id,
	)
	err := row1.Scan(&s)
	if err == sql.ErrNoRows {
		err = fmt.Errorf(
			"Нет такого id=%d",
			id,
		)
		return User{}, err
	} else {
		var k string
		//удаление строки в случае нахождения данной строки
		row := m.dataBase.QueryRow(
			"DELETE  FROM Misha2 where id=$1",
			id,
		)
		err = row.Scan(&k)
		if err == sql.ErrNoRows {
			err = fmt.Errorf(
				"Ошибка удаления строки из БД %s",
				err,
			)
			return User{}, err
		}
	}
	return User{}, err
}
