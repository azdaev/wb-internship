package main

import "fmt"

/*
	Строитель.
	Вид: порождающий.
	Суть паттерна - предоставление простого или урезанного интерфейса для работы
	со сложной подсистемой (фреймворка например).

	+: позволяет создавать тяжелые объекты пошагово
	+: переиспользование кода
	-: усложняет код программы из-за введения дополнительных классов
*/

type Computer struct {
	processorType string
	ramQuantity   int
	GPUs          []string
}

type IBuilder interface {
	setProcessorType()
	setRamQuantity()
	setGPUs()
	getComputer() Computer
}

func getBuilder(builderType string) IBuilder {
	if builderType == "budget" {
		return newBudgetBuilder()
	}

	if builderType == "expensive" {
		return newExpensiveBuilder()
	}
	return nil
}

type BudgetBuilder struct {
	processorType string
	ramQuantity   int
	GPUs          []string
}

func newBudgetBuilder() *BudgetBuilder {
	return &BudgetBuilder{}
}

func (b *BudgetBuilder) setProcessorType() {
	b.processorType = "Intel i3"
}

func (b *BudgetBuilder) setRamQuantity() {
	b.ramQuantity = 8
}

func (b *BudgetBuilder) setGPUs() {
	b.GPUs = []string{"Nvidia"}
}

func (b *BudgetBuilder) getComputer() Computer {
	return Computer{
		processorType: b.processorType,
		ramQuantity:   b.ramQuantity,
		GPUs:          b.GPUs,
	}
}

type ExpensiveBuilder struct {
	processorType string
	ramQuantity   int
	GPUs          []string
}

func newExpensiveBuilder() *ExpensiveBuilder {
	return &ExpensiveBuilder{}
}

func (b *ExpensiveBuilder) setProcessorType() {
	b.processorType = "Intel i7"
}

func (b *ExpensiveBuilder) setRamQuantity() {
	b.ramQuantity = 32
}

func (b *ExpensiveBuilder) setGPUs() {
	b.GPUs = []string{"Nvidia", "Nvidia"}
}

func (b *ExpensiveBuilder) getComputer() Computer {
	return Computer{
		processorType: b.processorType,
		ramQuantity:   b.ramQuantity,
		GPUs:          b.GPUs,
	}
}

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildComputer() Computer {
	d.builder.setProcessorType()
	d.builder.setRamQuantity()
	d.builder.setGPUs()
	return d.builder.getComputer()
}

func main() {
	budgetBuilder := getBuilder("budget")
	expensiveBuilder := getBuilder("expensive")

	director := newDirector(budgetBuilder)
	budgetComputer := director.buildComputer()

	fmt.Printf("Budget Computer:\n\tProcessor: %s\n\tRAM: %dGB\n\tGPU: %v\n", budgetComputer.processorType, budgetComputer.ramQuantity, budgetComputer.GPUs)

	director.setBuilder(expensiveBuilder)
	expensiveComputer := director.buildComputer()
	fmt.Printf("Expensive Computer:\n\tProcessor: %s\n\tRAM: %dGB\n\tGPU: %v\n", expensiveComputer.processorType, expensiveComputer.ramQuantity, expensiveComputer.GPUs)

}
