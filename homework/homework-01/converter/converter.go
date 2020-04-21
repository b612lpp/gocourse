package main

import (
	"fmt"
	"math"
)

func main() {
	const dollarCurency float64 = 70.00
	var usersRubles float64
	fmt.Println("How much rubbles you want to exchange?")
	fmt.Scan(&usersRubles)
	fmt.Println("You will get ", math.Round((usersRubles/dollarCurency)*100)/100, "US Dollars")
}
