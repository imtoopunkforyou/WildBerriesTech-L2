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
	ctx.SetOperation(&Addition{})
	ctx.Calculation(5, 2)
	ctx.SetOperation(&Subtraction{})
	ctx.Calculation(5, 2)
	ctx.SetOperation(&Multiplication{})
	ctx.Calculation(5, 2)
	ctx.SetOperation(&Division{})
	ctx.Calculation(5, 2)
}

// Интерфейс, который объединяет в себе математические операции над двумя int.
type Strategy interface {
	Calculation(x, y int)
}

// Интерфейс для работы с интерфейсом стратегии.
type Context struct {
	ctx Strategy
}

// Метод установки нужной стратегии.
func (c *Context) SetOperation(strategy Strategy) {
	c.ctx = strategy
}

// Выполнение математический операци под выбранную стратегию.
func (c *Context) Calculation(x, y int) {
	c.ctx.Calculation(x, y)
}

type Addition struct{}

func (a Addition) Calculation(x, y int) {
	fmt.Println(x + y)
}

type Subtraction struct{}

func (s Subtraction) Calculation(x, y int) {
	fmt.Println(x - y)
}

type Multiplication struct{}

func (m Multiplication) Calculation(x, y int) {
	fmt.Println(x * y)
}

type Division struct{}

func (a Division) Calculation(x, y int) {
	fmt.Println(x / y)
}
