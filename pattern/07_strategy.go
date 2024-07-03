package main

import "fmt"

/*
Применение:
 Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает каждый
 из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.
Плюсы:
 1. Уход от наследования к делегированию.
 2. Реализует принцип открытости/закрытости.
 3. Быстрая замена алгоритмов на ходу.
 4. Изолирует код и данные алгоритмов от других классов.

Минусы:
 1. Клиент должен знать разницу между стратегиями, для выбора нужной.
 2. Усложняет из-за введения дополнительных классов.
*/

func main() {
	ctx := Context{}
	ctx.SetOperation(&addition{})
	ctx.Calculation(5, 2)
	ctx.SetOperation(&subtraction{})
	ctx.Calculation(5, 2)
	ctx.SetOperation(&multiplication{})
	ctx.Calculation(5, 2)
	ctx.SetOperation(&division{})
	ctx.Calculation(5, 2)
}

// Strategy Интерфейс, который объединяет в себе математические операции над двумя int.
type Strategy interface {
	Calculation(x, y int)
}

// Context Интерфейс для работы с интерфейсом стратегии.
type Context struct {
	ctx Strategy
}

// SetOperation Метод установки нужной стратегии.
func (c *Context) SetOperation(strategy Strategy) {
	c.ctx = strategy
}

// Calculation Выполнение математический операции под выбранную стратегию.
func (c *Context) Calculation(x, y int) {
	c.ctx.Calculation(x, y)
}

type addition struct{}

func (a addition) Calculation(x, y int) {
	fmt.Println(x + y)
}

type subtraction struct{}

func (s subtraction) Calculation(x, y int) {
	fmt.Println(x - y)
}

type multiplication struct{}

func (m multiplication) Calculation(x, y int) {
	fmt.Println(x * y)
}

type division struct{}

func (a division) Calculation(x, y int) {
	fmt.Println(x / y)
}
