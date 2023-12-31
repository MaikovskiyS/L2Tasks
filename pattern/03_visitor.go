package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/
/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern


Посетитель — это поведенческий паттерн, который позволяет добавить новую операцию для целой иерархии классов, не изменяя код этих классов.

Проблема:
	Добавление нового функционала для объектов разнородных классов без изменения этих классов. Новый функционал обычно не уместен
	внутри этих разнородных классов (например, экспорт в XML данных из объектов "Физ.лицо", "Юр.лицо" в банковской сфере)
Решение:
	Паттерн Посетитель предлагает разместить новое поведение в отдельном классе, вместо того чтобы множить его сразу в нескольких
	классах. Объекты, с которыми должно было быть связано поведение, не будут выполнять его самостоятельно. Вместо этого вы будете
	передавать эти объекты в методы посетителя. Код поведения, скорее всего, должен отличаться для объектов разных классов,
	поэтому и методов у посетителя должно быть несколько


	В качестве примера рассмотрим внедрение "посетителя" в структуры "Круг" и "Прямоугольник".
	Посетитель добавляет поведение "Вычислить площадь" в структуры без изменения внутренней логики работы фигур.
	В случае добавления другого функционала в "посетителя", например getNumSides(), мы будем использовать все тот же метод accept(v visitor)
	без новых изменений структур фигур.
*/

type Shape interface {
	getType() string
	accept(Visitor)
}

// реализация интерфейса Shape: элемент "Прямоугольник"
type Rectangle struct {
	length float64
	width  float64
}

// "внедрение" метода для структур фигур:
func (r *Rectangle) accept(v Visitor) {
	v.visit(r)
}

func (r *Rectangle) getType() string {
	return "Rectangle"
}

// реализация интерфейса Shape: элемент "Круг"
type Circle struct {
	radius float64
}

// "внедрение" метода для структур фигур:
func (c *Circle) accept(v Visitor) {
	v.visit(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

// Интерфейс "Посетитель"
type Visitor interface {
	visit(Shape)
}

// Некоторая реализация интерфейса "Посетитель" - вычисление площади фигуры
type AreaCalculator struct {
	area float64
}

// создание "универсального" метода, который считает площадь для фигуры
func (a *AreaCalculator) visit(s Shape) {
	switch shapeType := s.(type) {
	case *Rectangle:
		// приводим к типу
		rec := s.(*Rectangle)
		// получаем доступ даже к закрытым полям структуры, ведь они доступны внутри самой структуры
		a.area = rec.length * rec.width
		fmt.Printf("Area of the rectangle is %v\n", a.area)
	case *Circle:
		cir := s.(*Circle)
		a.area = 3.14 * cir.radius * cir.radius
		fmt.Printf("Area of the circle is %v\n", a.area)
	default:
		fmt.Printf("Unknown figure: %T", shapeType)
	}
}

func InitVisitor() {
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{width: 2, length: 3}

	// создаем посетителя
	areaCalculator := &AreaCalculator{}
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)
}
