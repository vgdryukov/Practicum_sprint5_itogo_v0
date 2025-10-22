package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for index, value := range dataset {
		err := dp.Parse(value)
		if err != nil {
			log.Printf("Ошибка парсинга для элемента [%d]: %v", index, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка получения информации для элемента [%d]: %v", index, err)
		}
		fmt.Println(info)
	}
}
