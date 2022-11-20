package main

import "fmt"

/*
	Команда.
	Вид: Поведенческий.
	Суть паттерна - превращает запросы интерфейса программы к бизнес-логике в объекты.

	+: Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют
	+: Позволяет реализовать отмену операции
	+: Позволяет реализовать отложенный запуск операции
	+: Реализует принцип open/closed так как можно добавлять новый функционал не переписывая классы интерфейса программы
	-: Усложняет код из-за введения множества дополнительный классов
*/

type Device interface {
	on()
	off()
}

type Command interface {
	execute()
}

type onCommand struct {
	device Device
}

func (c *onCommand) execute() {
	c.device.on()
}

type offCommand struct {
	device Device
}

func (c *offCommand) execute() {
	c.device.off()
}

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

type Tv struct {
	isOn bool
}

func (t *Tv) on() {
	t.isOn = true
	fmt.Println("TV is now on")
}

func (t *Tv) off() {
	t.isOn = false
	fmt.Println("TV is now off")
}

func main() {
	tv := &Tv{}

	onCommand := &onCommand{device: tv}

	offCommand := &offCommand{device: tv}

	onButton := &Button{
		command: onCommand,
	}

	onButton.press()

	offButton := &Button{
		command: offCommand,
	}

	offButton.press()
}
