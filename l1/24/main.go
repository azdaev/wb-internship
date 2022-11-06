package main

import (
	"fmt"
	"github.com/azdaev/wb-internship/l1/24/data"
)

func main() {
	p1 := *data.NewPoint(1, 2)
	p2 := *data.NewPoint(4, 5)
	fmt.Println(data.Distance(p1, p2))
}
