package main

import (
	"fmt"
	"strings"
)

/*
Применение:
  Фасад — это структурный паттерн проектирования, который предоставляет простой интерфейс к сложной системе классов,
  библиотеке или фреймворку и позволяет изолировать код клиента от сложной подсистемы.

Плюсы:
 1. Снижает зависимость кода клиента от подсистемы.
 2. Изолирует клиента от компонентов подсистемы.
 3. Упрощает взаимодействие с подсистемой.

Минусы:
 1. Фасад может стать объектом, который делает слишком много/связан со всеми классами программы и будет слишком связанный
     объект, что усложнит его поддержание, тк может что-то сломаться.
*/

func main() {
	computer := NewComputer()
	fmt.Println(computer.StartComputer())
}

// NewComputer Конструктор  фасада.
func NewComputer() *Computer {
	return &Computer{cpu: &CPU{}, ram: &RAM{}, vram: &VideoRAM{}}
}

// Computer Передаем объекты в фасад, с которыми он будет работать.
type Computer struct {
	cpu  *CPU
	ram  *RAM
	vram *VideoRAM
}

// StartComputer Метод фасада, который будет использовать объекты других классов.
func (c *Computer) StartComputer() string {
	result := []string{c.cpu.startCPU(), c.ram.startRAM(), c.vram.startVideoRAM(), "Computer is started"}
	return strings.Join(result, "\n")
}

// CPU объект, входящий в состав фасада.
type CPU struct {
}

func (c *CPU) startCPU() string {
	return "CPU is started"
}

// RAM объект, входящий в состав фасада.
type RAM struct {
}

func (r *RAM) startRAM() string {
	return "RAM is started"
}

// VideoRAM объект, входящий в состав фасада.
type VideoRAM struct {
}

func (vr *VideoRAM) startVideoRAM() string {
	return "Video RAM is started"
}
