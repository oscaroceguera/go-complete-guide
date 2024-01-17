package main

import (
	"fmt"
	"testExample/helpers"
)

func main() {
	fmt.Println("Primer programa")

	h := helpers.NewHelpers()
	fmt.Printf("Mi primer numero es %d\n", h.Sum(1,2))

	mn := h.SumMultiple([]int{1,1,1,1})
	fmt.Printf("La suma de mis multiples numeros es %d",  mn)
}