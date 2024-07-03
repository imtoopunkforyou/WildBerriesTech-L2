package main

import "fmt"

/*
Применение:
Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.
Его можно использовать как с классом директор, так и без. Директор только задает последовательность постройки для строителя,
методы строителя можно вызывать напрямую из клиентского кода, директор удобен при нескольких шаблонах пастройки.
Используется когда объект имеет сложное создание и конфигурацию.

Плюсы:
 1. Позволяет изолировать логику постройки объекта.
 2. Позволяет переиспользование кода для постройки разных объектов.
 3. Создает объект пошагово

Минусы:
 1. Увеличивает и усложняет код из-за введения новых классов и объектов.
 2. Если слишком много конфигураций для постройки, то сложно станет управлять и расширять его
*/
func main() {
	builder := &ConcreteBuilder{house: &House{}}
	director := &Director{}
	director.setBuilder(builder)
	director.build("Wood floor", "Brick Walls", "Tile roof")
	house := builder.getHouse()
	fmt.Println(house.floor, "\n", house.walls, "\n", house.roof)

	builder.reset()
	builder.buildFloor("Brick floor")
	builder.buildWalls("Wood floor")
	builder.buildRoof("Tile roof")
	house = builder.getHouse()
	fmt.Println(house.floor, "\n", house.walls, "\n", house.roof)
}

// Builder Интерфейс строителя, который будет строить дом.
type Builder interface {
	buildFloor(floor string)
	buildWalls(walls string)
	buildRoof(roof string)
	getHouse() *House
	reset()
}

// Director Класс директор, который будет отдавать приказы строителю.
type Director struct {
	builder Builder
}

func (d *Director) setBuilder(builder Builder) {
	d.builder = builder
}

func (d *Director) build(floor, walls, roof string) {
	d.builder.buildFloor(floor)
	d.builder.buildWalls(walls)
	d.builder.buildRoof(roof)
}

// House Объект, который будет делать строитель.
type House struct {
	floor string
	walls string
	roof  string
}

// ConcreteBuilder Реализация строителя.
type ConcreteBuilder struct {
	house *House
}

func (b *ConcreteBuilder) buildFloor(floor string) {
	b.house.floor = floor
}

// BuildWalls Создает часть объекта House(стены)
func (b *ConcreteBuilder) buildWalls(walls string) {
	b.house.walls = walls
}

func (b *ConcreteBuilder) buildRoof(roof string) {
	b.house.roof = roof
}

func (b *ConcreteBuilder) getHouse() *House {
	return b.house
}

func (b *ConcreteBuilder) reset() {
	b.house = &House{}
}
