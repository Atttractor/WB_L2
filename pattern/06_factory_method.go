package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type creator struct {
	factory factory
}

func (c *creator) create() product {
	p := c.factory.factoryMethod()
	p.getInfo()
	return p
}

type factory interface {
	factoryMethod() product
}

type creator1 struct{}

func (c1 *creator1) factoryMethod() product {
	return new(product1)
}

type creator2 struct{}

func (c1 *creator2) factoryMethod() product {
	return new(product2)
}

type product interface {
	getInfo()
}

type product1 struct{}

func (p1 *product1) getInfo() {
	fmt.Println("product1")
}

type product2 struct{}

func (p2 *product2) getInfo() {
	fmt.Println("product2")
}

//func main() {
//	c := creator{new(creator1)}
//	c.create()
//	c = creator{new(creator2)}
//	c.create()
//}
