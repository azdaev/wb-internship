package main

import "fmt"

type Charger interface {
	ChargeTypeC()
}

// Samsung - структура, которая может заряжаться от type-c
type Samsung struct {
}

func (s *Samsung) ChargeTypeC() {
	fmt.Println("Samsung: Getting power from type-c")
}

// Iphone - структура, которая не может заряжаться от type-c, но может от lightning
type Iphone struct {
}

func (i *Iphone) ChargeLightning() {
	fmt.Println("iPhone: Getting power from lightning")
}

// IphoneToTypeCAdapter - адаптер, чтобы заряжать Iphone от type-c
type IphoneToTypeCAdapter struct {
	iphone *Iphone
}

func (adapter *IphoneToTypeCAdapter) ChargeTypeC() {
	fmt.Println("Connected adapter to iPhone")
	adapter.iphone.ChargeLightning()
}

func ConnectTypeC(device Charger) {
	device.ChargeTypeC()
}

func main() {
	samsung := &Samsung{}
	iphone := &Iphone{}
	adapter := &IphoneToTypeCAdapter{iphone: iphone}

	ConnectTypeC(samsung)
	ConnectTypeC(adapter)
}
