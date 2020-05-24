package main

import (
	"fmt"
	"math/rand"
	"time"
)

type carDefenition struct {
	speed       int
	preparation int
}

const s int = 1000

func main() {
	carOne := genStats()
	fmt.Println(carOne)
}

func genStats() carDefenition {
	var getStats carDefenition
	rand.Seed(time.Now().UnixNano())
	getStats.speed = rand.Intn(50)
	getStats.preparation = rand.Intn(10)
	//fmt.Println(getStats)
	return getStats
}
