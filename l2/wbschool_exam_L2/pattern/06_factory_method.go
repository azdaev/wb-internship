package main

import "fmt"

/*
	Фабричный метод.
	Вид: Порождающий.
	Суть паттерна - позволяет определяет общий интерфейс для создания объектов в суперклассе, позволяя подклассам
изменять тип создаваемых объектов.

	+: Избавляет класс от привязки к конкретным классам продуктов
	+: Выделяет код производства продуктов в одно место, упрощая поддержку кода
	+: Упрощает добавление новых продуктов в программу
	+: Реализуется принцип Open/Closed
	-: Может привести к созданию больших параллельных иерархий классов, так
как для каждого класса продукта надо создать свой подкласс создателя.
*/

type ITransport interface {
	setName(name string)
	setPath(path string)
	getName() string
	getPath() string
}

type Transport struct {
	name string
	path string
}

func (g *Transport) setName(name string) {
	g.name = name
}

func (g *Transport) getName() string {
	return g.name
}

func (g *Transport) setPath(path string) {
	g.path = path
}

func (g *Transport) getPath() string {
	return g.path
}

type Boat struct {
	Transport
}

func newBoat() ITransport {
	return &Boat{
		Transport: Transport{
			name: "Boat",
			path: "Sea",
		},
	}
}

type Airplane struct {
	Transport
}

func newAirplane() ITransport {
	return &Airplane{
		Transport: Transport{
			name: "Boeing",
			path: "Air",
		},
	}
}

func getTransport(TransportType string) (ITransport, error) {
	if TransportType == "boat" {
		return newBoat(), nil
	}
	if TransportType == "airplane" {
		return newAirplane(), nil
	}
	return nil, fmt.Errorf("wrong Transport type passed")
}

func printDetails(g ITransport) {
	fmt.Printf("Transport: %s", g.getName())
	fmt.Println()
	fmt.Printf("Path: %s", g.getPath())
	fmt.Println()
}

func main() {
	boat, _ := getTransport("boat")
	airplane, _ := getTransport("airplane")

	printDetails(boat)
	printDetails(airplane)
}
