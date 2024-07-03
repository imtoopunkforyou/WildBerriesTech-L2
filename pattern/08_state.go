package main

import "fmt"

/*
Применение:

	Состояние — это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от
	своего состояния. Извне создаётся впечатление, что изменился класс объекта.

Плюсы:
 1. Упрощает код контекста.
 2. Выделяет код состояний в одном место, что упрощает поддержку кода.
 3. Избавляет от большей части услонвых операторов проверки состояния.

Минусы:
 1. Может неоправданно усложнить код, если мало состояний и они редко меняются
*/
func main() {
	phone := NewMobileAlert()
	phone.Alert()
	phone.SetState(&loudMode{})
	phone.Alert()
}

// MobileAlertState Интерфейс состояния.
type MobileAlertState interface {
	Alert()
}

// MobileAlert Структура содержащая текущее состояние.
type MobileAlert struct {
	State MobileAlertState
}

// Alert Вызов текущего состояния.
func (m *MobileAlert) Alert() {
	m.State.Alert()
}

// SetState Установка состояния для объекта.
func (m *MobileAlert) SetState(state MobileAlertState) {
	m.State = state
}

// NewMobileAlert Конструктор состояния.
func NewMobileAlert() *MobileAlert {
	return &MobileAlert{&silentMode{}}
}

type silentMode struct{}

func (s silentMode) Alert() {
	fmt.Println("Phone in silent mode")
}

type loudMode struct{}

func (l loudMode) Alert() {
	fmt.Println("Phone in loud mode")
}
