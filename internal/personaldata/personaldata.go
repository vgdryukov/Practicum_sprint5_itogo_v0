package personaldata

import "fmt"

type Personal struct {
	// TODO: добавить поля
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	// TODO: реализовать функцию
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.2f\n", p.Weight)
	fmt.Printf("Рост: %.2f\n", p.Height)
}
