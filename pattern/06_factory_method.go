package main

import (
	"fmt"
	"log"
)

/*
Применение:
 Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов в
 суперклассе, позволяя подклассам изменять тип создаваемых объектов. Нужен когда система должна быть независима от
 способа создания объектов.

Плюсы:
 1. Избавляет класс от привязки к конерктным продуктам.
 2. Реализует принцип открытости/закрытости.
 3. Упрощает добавление новых продуктов в программу.
 4. Выделяет код производства в одном место, что упрощает поддержку кода.

Минусы:
 1. Может привести к созданию больших параллельеых иерархий классов, так как для каждого класса продукта, нужно создать
	свой подкласс создателя.
*/

func main() {
	factory := NewCreator()
	product := factory.CreateProduct(1)
	product.Use()
	product = factory.CreateProduct(3)
	product.Use()
	product = factory.CreateProduct(2)
	product.Use()
}

// Интерфейс для создания продукта фабрикой.
type Creator interface {
	CreateProduct(product int) Product
}

// Интерфейс созданных продуктов фабрикой.
type Product interface {
	Use()
}

type ConcreteCreator struct{}

// Конструктор фабрики.
func NewCreator() Creator {
	return &ConcreteCreator{}
}

func (c *ConcreteCreator) CreateProduct(create int) Product {
	var product Product
	switch create {
	case 1:
		product = &ConcreteProductA{}
	case 2:
		product = &ConcreteProductB{}
	case 3:
		product = &ConcreteProductC{}
	default:
		log.Fatal("Create product fail")
	}
	return product
}

type ConcreteProductA struct{}

func (c *ConcreteProductA) Use() {
	fmt.Println("Used first product")
}

type ConcreteProductB struct{}

func (c *ConcreteProductB) Use() {
	fmt.Println("Used second product")
}

type ConcreteProductC struct{}

func (c *ConcreteProductC) Use() {
	fmt.Println("Used third product")
}
