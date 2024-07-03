package main

import "fmt"

/*
Применение:

	Команда — это поведенческий паттерн проектирования, который превращает запросы в объекты, позволяя передавать их как
	аргументы при вызове методов, ставить запросы в очередь, логировать их, а также поддерживать отмену операций.

Плюсы:
 1. Позволяет собирать сложные команды из простых.
 2. Позволяет реализовать отложенный запуск операции.
 3. Позволяет реализовать простой повтор и отмену операции.
 4. Убирает прямую зависимость между объектами, вызывающими операции и выполняющие их.
 5. Реализует принцип открытости/закрытости

Минусы:
 1. Усложняет код, введением множества дополнительных классов.
*/
func main() {
	receiver := &receiver{}
	invoker := &invoker{}
	invoker.AddCommand(&buttonON{receiver: receiver})
	invoker.AddCommand(&buttonOFF{receiver: receiver})
	invoker.Execute()
}

type command interface {
	Execute()
}

// Получатель, имеющий набор действий, которые команда может запрашивать.
type receiver struct{}

func (r *receiver) LightON() {
	fmt.Println("Light is on")
}

func (r *receiver) LightOFF() {
	fmt.Println("Light is off")
}

// Классы, реализующие конкретные команды.
type buttonON struct {
	receiver *receiver
}

func (b *buttonON) Execute() {
	b.receiver.LightON()
}

type buttonOFF struct {
	receiver *receiver
}

func (b *buttonOFF) Execute() {
	b.receiver.LightOFF()
}

// Инициатор, записывающий команды в стек и провоцирует их выполнение.
type invoker struct {
	commands []command
}

func (i *invoker) AddCommand(command command) {
	i.commands = append(i.commands, command)
}

func (i *invoker) DeleteCommand() {
	if len(i.commands) > 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

func (i *invoker) Execute() {
	for _, command := range i.commands {
		command.Execute()
	}
}
