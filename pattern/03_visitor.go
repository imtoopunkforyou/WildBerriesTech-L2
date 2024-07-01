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
	visitor := &TownVisitor{}
	places := []Place{&Fabric{}, &Garage{}, &Hospital{}}

	for _, place := range places {
		place.Accept(visitor)
	}
}

// интерфейс, чтобы посетитель мог посетить нужный класс через него
type Place interface {
	Accept(v Visitor)
}

// интерфейс посетителя
type Visitor interface {
	VisitGarage(g *Garage)
	VisitHospital(h *Hospital)
	VisitFabric(f *Fabric)
}

// класс, имплеминтирующая интерфейс посетителя
type TownVisitor struct{}

func (v *TownVisitor) VisitGarage(g *Garage) {
	fmt.Println("I'm visit garage")
}

func (v *TownVisitor) VisitHospital(h *Hospital) {
	fmt.Println("I'm visit hospital")
}

func (v *TownVisitor) VisitFabric(f *Fabric) {
	fmt.Println("I'm visit fabric")
}

type Garage struct{}

func (h *Garage) Accept(v Visitor) {
	v.VisitGarage(h)
}

type Hospital struct{}

func (h *Hospital) Accept(v Visitor) {
	v.VisitHospital(h)
}

type Fabric struct{}

func (f *Fabric) Accept(v Visitor) {
	v.VisitFabric(f)
}
