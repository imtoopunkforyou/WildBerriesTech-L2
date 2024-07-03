package main

import "fmt"

/*
Применение:

	Цепочка обязанностей — это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно
	по цепочке обработчиков. Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли
	передавать запрос дальше по цепи. Нужен когда обработчиков больше одного.

Плюсы:
 1. Уменьшает зависимость между клиентом и обработчиками.
 2. Реализует принцип открытости/закрытости.
 3. Реализует принцип единой обязанности.

Минусы:
 1. Запрос может остаться ни кем не обработанный и потеряться.
*/
func main() {
	handlers := NewChain()
	handlers.handleRequest(3)
}

// NewChain Конструктор цепочки запросов.
func NewChain() *ConcreteHandlerA {
	return &ConcreteHandlerA{next: &ConcreteHandlerB{next: &ConcreteHandlerC{}}}
}

type handler interface {
	handleRequest(request int)
}

// ConcreteHandlerA Первый обработчик.
type ConcreteHandlerA struct {
	next handler
}

func (h *ConcreteHandlerA) handleRequest(request int) {
	if request == 1 {
		fmt.Println("ConcreteHandlerA")
	} else if h.next != nil {
		h.next.handleRequest(request)
	}
}

// ConcreteHandlerB Второй обработчик.
type ConcreteHandlerB struct {
	next handler
}

func (h *ConcreteHandlerB) handleRequest(request int) {
	if request == 2 {
		fmt.Println("ConcreteHandlerB")
	} else if h.next != nil {
		h.next.handleRequest(request)
	}
}

// ConcreteHandlerC Третий обработчик.
type ConcreteHandlerC struct {
	next handler
}

func (h *ConcreteHandlerC) handleRequest(request int) {
	if request == 3 {
		fmt.Println("ConcreteHandlerC")
	} else if h.next != nil {
		h.next.handleRequest(request)
	}
}
