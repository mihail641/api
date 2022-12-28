package model

import (
	"example.com/projectApiClient"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

//постоянная с адресом из какой директории надо получить информацию о вложенности
const path = "/home/kate/Music"

//DirectoriesModel используется для конструктора модели
type DirectoriesModel struct {
}

// NewDirectoriesModel конструктор модели
func NewDirectoriesModel() *DirectoriesModel {
	return &DirectoriesModel{}
}

// GetDirectories метод модели получающий слайс Директорий находящихся по адресу path
func (dir *DirectoriesModel) GetDirectories(path string) ([]projectApiClient.Directory, error) {
	//создание экземпляра структуры Directory
	directories := []projectApiClient.Directory{}
	var firstDirectoryTitle string
	//разбиение path на слова по разделителю "/"
	str := strings.Split(path, "/")
	str = strings.Split(path, "\\")
	num := 1
	//длина строки path
	lastDir := len(str)
	//цикл для определения "первоначальной" диретории
	for i := 0; i < lastDir; i++ {
		firstDirectoryTitle = ``
		firstDirectory := str[lastDir-1]
		firstDirectoryTitle = firstDirectoryTitle + firstDirectory
	}
	//инициализация данных, полученых выше в структуру Directory
	directory := projectApiClient.Directory{Id: num, Title: firstDirectoryTitle}
	//обращение к рекурсионной функции readDir
	err := readDir(path, &directory.Directories, num)
	if err != nil {
		fmt.Println("Ошибка извлечения директорий", err)
		return nil, err
	}
	//добавление структуры в слайс структур
	directories = append(directories, directory)
	return directories, nil
}

//readDir рекурсионная функция необходимая длянахождения всех вложенных Диреторий
func readDir(path string, directories *[]projectApiClient.Directory, num int) error {
	//чиение диреторий находящихся в path, в случае если бы необходимо
	//прочитать вместе с файлами нужно использовать Reddirnames
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	directory := projectApiClient.Directory{}
	//цикл для изменения адреса path, определения является ли объект директорией или файлом,
	//а так е рексурсионным вызовом самой себя, для получения вложенных объектов
	for _, file := range files {
		path := filepath.Join(path, file.Name())
		fmt.Println("path ", path)
		if file.IsDir() == true {
			num = num + 1
			fmt.Println("file.Name()", file.Name())
			directory = projectApiClient.Directory{Id: num, Title: file.Name()}
			fmt.Println("directory", directory)
			err = readDir(path, &directory.Directories, num)
			fmt.Println("directory.Directories", &directory.Directories)
			//directories = append(directories, directory)
			fmt.Println("directories", directories)
			*directories = append(*directories, directory)
		}
	}
	fmt.Println("directories", directories)
	return err
}
