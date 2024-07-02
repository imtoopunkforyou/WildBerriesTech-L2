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
	director.SetBuilder(builder)
	director.Build("Wood floor", "Brick Walls", "Tile roof")
	house := builder.GetHouse()
	fmt.Println(house.floor, "\n", house.walls, "\n", house.roof)

	builder.Reset()
	builder.BuildFloor("Brick floor")
	builder.BuildWalls("Wood floor")
	builder.BuildRoof("Tile roof")
	house = builder.GetHouse()
	fmt.Println(house.floor, "\n", house.walls, "\n", house.roof)
}

// Интерфейс строителя, который будет строить дом.
type Builder interface {
	BuildFloor(floor string)
	BuildWalls(walls string)
	BuildRoof(roof string)
	GetHouse() *House
	Reset()
}

// Класс директор, который будет отдавать приказы строителю.
type Director struct {
	builder Builder
}

// Метод для установки конкретного строителя для объекта.
func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

// Метод для отдачи приказов директором строителю.
func (d *Director) Build(floor, walls, roof string) {
	d.builder.BuildFloor(floor)
	d.builder.BuildWalls(walls)
	d.builder.BuildRoof(roof)
}

type House struct {
	floor string
	walls string
	roof  string
}

type ConcreteBuilder struct {
	house *House
}

func (b *ConcreteBuilder) BuildFloor(floor string) {
	b.house.floor = floor
}

func (b *ConcreteBuilder) BuildWalls(walls string) {
	b.house.walls = walls
}

func (b *ConcreteBuilder) BuildRoof(roof string) {
	b.house.roof = roof
}

func (b *ConcreteBuilder) GetHouse() *House {
	return b.house
}

func (b *ConcreteBuilder) Reset() {
	b.house = &House{}
}
