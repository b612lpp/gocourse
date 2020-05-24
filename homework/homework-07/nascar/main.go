package main

import (
	"fmt"
	"math/rand"
)

type carDefenition struct {
	speed       int
	preparation int
}

const s int = 1000

func main() {
	genStats()
}

func genStats() carDefenition {
	var getStats carDefenition
	getStats.speed = rand.Intn((10 - 50) + 10)
	getStats.preparation = rand.Intn(1-5) + 1
	fmt.Println(getStats)
	return getStats
}
