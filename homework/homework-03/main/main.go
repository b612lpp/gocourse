package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	//StatusFile contains file location
	StatusFile = "status.json"
)

//OpenedWindows contains windows status
type OpenedWindows struct{ LeftWindowOpen, RightWindowOpen bool }

//PassangerCar has structure of cars
type PassangerCar struct {
	CarMark     string
	YearMade    int
	TrunkSize   float32
	EngineStart bool
	WindowsOpen OpenedWindows
}

func main() {
	var CurrentCar PassangerCar
	CurrentCar = readFromJSON()
	fmt.Println(CurrentCar)
	drawHello(CurrentCar)

}

func drawHello(c PassangerCar) {
	var calledFunc string
	//CurrentCar := readFromJSON()
	fmt.Println("---------------------------")
	fmt.Println("Today we have following car: ", c.CarMark)
	fmt.Println("It is possible to:\n 1. Put something in it's Trunk\n 2. Change it's Mark \n 3. Check windows status\n 4. Close the catalog")
	fmt.Scan(&calledFunc)

	switch calledFunc {
	case "1":
		PassangerCar.changeTrunkCont(c)
	case "2":
		PassangerCar.ChangeMark(c)
	case "3":
		PassangerCar.funWithWindows(c)
	case "4":
		break
	default:
		fmt.Println("Undifined action")
		drawHello(c)
	}
}

func (c PassangerCar) funWithWindows() {

	left := c.WindowsOpen.LeftWindowOpen
	right := c.WindowsOpen.RightWindowOpen

	if left == true {
		fmt.Println("Left window is open.")
	} else {
		fmt.Println("Left window is closed")
	}
	if right == true {
		fmt.Println("Right window is open")
	} else {
		fmt.Println("Right window is closed")
	}
	drawHello(c)

}

//ChangeMark changes mark title
func (c PassangerCar) ChangeMark() {
	var MarkChangeTo string
	fmt.Println("Current Mark's title: ", c.CarMark)
	fmt.Println("Which title you want to swap to?")
	fmt.Scanln(&MarkChangeTo)
	c.CarMark = MarkChangeTo
	saveToJSON(c)
	fmt.Println("Your awesome name is assigned to current Car", c.CarMark)
	drawHello(c)

}

func (c PassangerCar) changeTrunkCont() {
	var litersToUpload float32

	fmt.Println("Current free space of Trunk is: ", c.TrunkSize)
	fmt.Println("How many liters you want to upload?")
	fmt.Scanln(&litersToUpload)
	c.TrunkSize -= litersToUpload
	saveToJSON(c)
	fmt.Println("Your things are uploaded successfully! Now current free space of Trunk is", c.TrunkSize)
	drawHello(c)
}

func readFromJSON() PassangerCar {

	file, _ := ioutil.ReadFile(StatusFile)

	data := PassangerCar{}

	_ = json.Unmarshal([]byte(file), &data)
	return data

}

func saveToJSON(s PassangerCar) {

	file, _ := json.MarshalIndent(s, "", " ")

	_ = ioutil.WriteFile(StatusFile, file, 0644)
}
