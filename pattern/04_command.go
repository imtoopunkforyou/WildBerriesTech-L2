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
	receiver := &Receiver{}
	invoker := &Invoker{}
	invoker.AddCommand(&ButtonON{receiver: receiver})
	invoker.AddCommand(&ButtonOFF{receiver: receiver})
	invoker.Execute()
}

type Command interface {
	Execute()
}

// Получатель, имеющий набор действий, которые команда может запрашивать.
type Receiver struct{}

func (r *Receiver) LightON() {
	fmt.Println("Light is on")
}

func (r *Receiver) LightOFF() {
	fmt.Println("Light is off")
}

// Классы, реализующие конкретные команды.
type ButtonON struct {
	receiver *Receiver
}

func (b *ButtonON) Execute() {
	b.receiver.LightON()
}

type ButtonOFF struct {
	receiver *Receiver
}

func (b *ButtonOFF) Execute() {
	b.receiver.LightOFF()
}

// Инициатор, записывающий команды в стэк и провоцирует их выполнение.
type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) DeleteCommand() {
	if len(i.commands) > 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

func (i *Invoker) Execute() {
	for _, command := range i.commands {
		command.Execute()
	}
}
