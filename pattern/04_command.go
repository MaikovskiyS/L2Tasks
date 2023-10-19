package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Command_pattern

Команда — это поведенческий паттерн проектирования, который превращает запросы в объекты, позволяя передавать их как аргументы
при вызове методов, ставить запросы в очередь, логировать их, а также поддерживать отмену операций.

Цель:
Создание структуры, в которой класс-отправитель и класс-получатель не зависят друг от друга напрямую.
Организация обратного вызова к классу, который включает в себя класс-отправитель.

Плюсы:

	Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
	Позволяет реализовать простую отмену и повтор операций.
	Позволяет реализовать отложенный запуск операций.
	Позволяет собирать сложные команды из простых.
	Реализует принцип открытости/закрытости.

	 Например, Паттерн Команда предлагает больше не отправлять вызовы из UI (кнопки) напрямую в бизнес-логику.
	 Вместо этого каждый вызов, отличающийся от других, следует завернуть в собственный класс с единственным методом,
	который и будет осуществлять вызов. Такие объекты называют командами.

Минусы:

	Усложняет код программы из-за введения множества дополнительных классов.

В качестве примера рассмотрим ресторан, в котором есть определенное количество поваров, и блюда на кухне.
Каждый повар может одновременно выполнять одно из следующих заданий:

	Готовить пиццу
	Сделать салат
	Мыть посуду

Каждый раз, когда готовится пицца или салат, используется тарелка. При мытье посуды общее количество посуды (в ресторане) сбрасывается.

Шаблон Command полезен, когда вам нужно выполнить задачи, но вы хотите отделить управление задачами от выполнения самой задачи.
В примере ниже мы отделили исполнителей (поваров) от задач, инкапсулируя каждую задачу в общий интерфейс.

// Базовой единицей для реализации шаблона команд является командный интерфейс:
// Классы команд можно объединить под общим интерфейсом c единственным методом запуска.
// После этого одни и те же отправители (Invoker, Повар) смогут работать с различными командами, не привязываясь к их классам.
// Даже больше: команды можно будет взаимозаменять на лету, изменяя итоговое поведение отправителей
*/
type Command interface {
	execute()
}

type Restaurant struct {
	TotalDishes   int
	CleanedDishes int
}

// Конструктор ресторана, который имеет 10 тарелок. И они чистые
func NewResteraunt() *Restaurant {
	const totalDishes = 10
	return &Restaurant{
		TotalDishes:   totalDishes,
		CleanedDishes: totalDishes,
	}
}

type MakePizzaCommand struct {
	n          int
	restaurant *Restaurant // Объект Command знает о приёмнике и вызывает метод приемника.
}

func (c *MakePizzaCommand) execute() {
	c.restaurant.CleanedDishes -= c.n // Значения параметров приёмника изменяются в команде.
	fmt.Println("made", c.n, "pizzas")
}

type MakeSaladCommand struct {
	n          int
	restaurant *Restaurant // Объект Command знает о приёмнике и вызывает метод приемника.
}

func (c *MakeSaladCommand) execute() {
	c.restaurant.CleanedDishes -= c.n // Значения параметров приёмника изменяются в команде.
	fmt.Println("made", c.n, "salads")
}

type CleanDishesCommand struct {
	restaurant *Restaurant // Объект Command знает о приёмнике и вызывает метод приемника.
}

func (c *CleanDishesCommand) execute() {
	c.restaurant.CleanedDishes = c.restaurant.TotalDishes // Значения параметров приёмника изменяются в команде.
	fmt.Println("dishes cleaned")
}

// Теперь мы можем добавить методы в Restaurant чтобы создать экземпляры этих команд.
// Таким образом, Restaurant действует как своего рода фабрика команд.
func (r *Restaurant) MakePizza(n int) Command {
	return &MakePizzaCommand{
		restaurant: r,
		n:          n,
	}
}

func (r *Restaurant) MakeSalad(n int) Command {
	return &MakeSaladCommand{
		restaurant: r,
		n:          n,
	}
}

func (r *Restaurant) CleanDishes() Command {
	return &CleanDishesCommand{
		restaurant: r,
	}
}

type Cook struct {
	Commands []Command
}

// Метод выполнения команд поваром
func (c *Cook) executeCommands() {
	for _, c := range c.Commands {
		c.execute()
	}
}

func InitCommand() {
	r := NewResteraunt()

	// create the list of tasks to be executed
	tasks := []Command{
		r.MakePizza(2),
		r.MakeSalad(1),
		r.MakePizza(3),
		r.CleanDishes(),
		r.MakePizza(4),
		r.CleanDishes(),
	}

	cooks := []*Cook{
		{},
		{},
	}

	for i, task := range tasks {
		cook := cooks[i%len(cooks)]
		cook.Commands = append(cook.Commands, task)
	}

	for i, c := range cooks {
		fmt.Println("cook", i, ":")
		c.executeCommands()
		fmt.Println()
	}
}
