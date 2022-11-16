package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

/*
	Фасад.
	Вид: Структурный.
	Суть паттерна - предоставление простого или урезанного интерфейса для работы
	со сложной подсистемой (фреймворка например).

	+: изоляция клиента от сложной подсистемы, тем самым ее облегченное использование
	-: риск создания перегруженного объекта
*/

type PaymentService struct {
	AuthToken string
}

func NewPaymentsService(token string) *PaymentService {
	return &PaymentService{token}
}

func (ps *PaymentService) Transfer(quantity int, sender string) {
	fmt.Printf("Transfered %d to %s", quantity, sender)
}

type DeliveryService struct {
	AuthToken string
}

func NewDeliveryService() *DeliveryService {
	return &DeliveryService{}
}

func (ds *DeliveryService) SendDelivery(parcels []string) string {
	packageId := strconv.Itoa(rand.Intn(20000-10000) + 10000)
	fmt.Printf("Started delivery of order #%s\n", packageId)
	return packageId
}

type Shop struct {
}

func NewShop() *Shop {
	return &Shop{}
}

// Фасад

func (s *Shop) Order(BuyerAuthToken string, order []string) {
	prices := map[string]int{
		"MacBook 14":        849,
		"MacBook 14 Pro":    949,
		"iPhone 14 Pro":     399,
		"iPhone 14 Pro Max": 449,
		"AirPods Pro":       249,
	}
	ps := NewPaymentsService("restore42_token")
	ds := NewDeliveryService()
	totalPrice := 0
	for _, product := range order {
		totalPrice += prices[product]
	}
	ps.Transfer(totalPrice, BuyerAuthToken)
	trackID := ds.SendDelivery(order)
	fmt.Printf("Succesfully ordered. Total: %d\nTrack-number: %s", totalPrice, trackID)
}

func main() {
	order := []string{"MacBook 14 Pro", "iPhone 14 Pro Max", "AirPods Pro"}
	token := "amady833toktok"
	s := NewShop()
	s.Order(token, order)
}
