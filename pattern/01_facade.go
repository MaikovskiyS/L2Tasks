package pattern

import (
	"errors"
	"fmt"
	"log"
	"time"
)

/*
	Реализовать паттерн «фасад».

Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Фасад -упрошенный интерфейс для управления сложными подсистемами
+ - упрошает работу с большими объектами
- - может стать god_object
*/

// examples for shop facade
var (
	inShowCaseItem = "toy"
	inStorageItem  = "milk"
	wrongItem      = "item"
)

type Shop struct {
	orderCh  chan string
	showcase Showcase
	storage  Storage
}

func InitFacade() {
	orderch := make(chan string)
	showCaseitem := &Item{Name: inShowCaseItem}
	storageItem := &Item{Name: inStorageItem}

	factory := NewFactory(orderch)
	shop := NewShop(orderch)
	shop.showcase.store[inShowCaseItem] = showCaseitem
	shop.storage.store[inStorageItem] = storageItem
	go func(*Factory) {
		for {
			select {
			case order := <-factory.orderCh:
				factory.CreateItem(order)
			default:
				fmt.Println("doing smth on factory")
			}
			time.Sleep(2 * time.Second)
		}
	}(factory)

	//facade
	good, err := shop.SellItem(wrongItem)
	if err != nil {
		log.Fatal("товар производится")
	}
	fmt.Printf("%s sold\n", good.Name)
}

func NewShop(ch chan string) Shop {

	return Shop{
		orderCh:  ch,
		showcase: NewShowcase(),
		storage:  NewStorage(),
	}
}
func (s Shop) SellItem(name string) (*Item, error) {
	fmt.Println("продаем товар...проверяем на витрине")
	item, err := s.showcase.GetItem(name)
	if err != nil {
		fmt.Println("товара нет на витрине, идем на склад")
		item, err := s.storage.GetItem(name)
		if err != nil {
			fmt.Println("товара нет на складе, делаем заказ на фабрику")
			s.CreateOrder()
			return nil, err
		}
		return item, nil
	}
	return item, nil
}
func (s *Shop) CreateOrder() {
	s.orderCh <- "new order"
}

type Showcase struct {
	store map[string]*Item
}

func NewShowcase() Showcase {
	return Showcase{
		store: make(map[string]*Item),
	}

}
func (c *Showcase) GetItem(name string) (*Item, error) {
	item, ok := c.store[name]
	if !ok {
		return nil, errors.New("not found")
	}
	return item, nil
}

type Storage struct {
	store map[string]*Item
}

func NewStorage() Storage {
	return Storage{
		store: make(map[string]*Item),
	}
}
func (d *Storage) GetItem(name string) (*Item, error) {
	item, ok := d.store[name]
	if !ok {
		return nil, errors.New("not found")
	}
	return item, nil
}

type Factory struct {
	orderCh chan string
}

func NewFactory(ch chan string) *Factory {
	return &Factory{
		orderCh: ch,
	}
}
func (f *Factory) CreateItem(string) *Item {
	fmt.Println("изготовление товара")
	return &Item{Name: "New item"}
}

type Item struct {
	Name string
}
