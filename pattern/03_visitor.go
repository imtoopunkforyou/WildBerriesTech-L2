package main

import "fmt"

/*
Применение:
  Посетитель — это поведенческий паттерн проектирования, который позволяет добавлять в программу новые операции, не изменяя
  классы объектов, над которыми эти операции могут выполняться. Добавляем новый функционал, не меняя изначальный объект.

Плюсы:
 1. Посетитель может накапливать состояние при обходе структуры элементов.
 2. Упрощает добавление новых операций над объектами.
 3. Объединяет родственные операции в одном классе.

Минусы:
 1. Может нарушить инкапсуляцию элементов.
 2. Не имеет смысла, если иерархия элементов часто меняется.
*/

func main() {
	visitor := &townVisitor{}
	places := []Place{&fabric{}, &garage{}, &hospital{}}

	for _, place := range places {
		place.Accept(visitor)
	}
}

// Place Интерфейс, чтобы посетитель мог посетить нужный класс через него.
type Place interface {
	Accept(v Visitor)
}

// Visitor Интерфейс посетителя.
type Visitor interface {
	visitGarage(_ *garage)
	visitHospital(_ *hospital)
	visitFabric(_ *fabric)
}

type townVisitor struct{}

func (v *townVisitor) visitGarage(_ *garage) {
	fmt.Println("I'm visit garage")
}

func (v *townVisitor) visitHospital(_ *hospital) {
	fmt.Println("I'm visit hospital")
}

func (v *townVisitor) visitFabric(_ *fabric) {
	fmt.Println("I'm visit fabric")
}

type garage struct{}

func (h *garage) Accept(v Visitor) {
	v.visitGarage(h)
}

type hospital struct{}

func (h *hospital) Accept(v Visitor) {
	v.visitHospital(h)
}

type fabric struct{}

func (f *fabric) Accept(v Visitor) {
	v.visitFabric(f)
}
