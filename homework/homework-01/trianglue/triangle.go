package main

import (
	"fmt"
	"math"
)

func main() {
	var catetA float64
	var catetB float64
	fmt.Println("Provide lenght of side A, please")
	fmt.Scan(&catetA)
	fmt.Println("Provide lenght of side B, please")
	fmt.Scan(&catetB)

	gipotenuZa := math.Sqrt(math.Pow(catetA, 2) + math.Pow(catetB, 2))
	perimeTr := catetA + catetB + gipotenuZa
	sQare := (catetA * catetB) / 2
	fmt.Println("Lenght of hypotinuze is: ", convert(gipotenuZa))
	fmt.Println("Lenght of perimetr: ", convert(perimeTr))
	fmt.Println("Square of the triangle is: ", convert(sQare))
}

func convert(tringleSide float64) string {

	x := fmt.Sprintf("%.2f", tringleSide)
	return x
}
