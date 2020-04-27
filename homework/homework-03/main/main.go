package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	drawHello()

}

func drawHello() {
	var calledFunc string
	CurrentCar := readFromJSON()
	fmt.Println("---------------------------")
	fmt.Println("Today we have following car: ", CurrentCar.CarMark)
	fmt.Println("It is possible to:\n 1. Put something in it's Trunk\n 2. Change it's Mark \n 3. Check windows status\n 4. Close the catalog")
	fmt.Scan(&calledFunc)

	switch calledFunc {
	case "1":
		changeTrunkCont()
	case "2":
		changeMark()
	case "3":
		funWithWindows()
	case "4":
		break
	default:
		fmt.Println("Undifined action")
		drawHello()
	}
}

func funWithWindows() {
	left := readFromJSON().WindowsOpen.LeftWindowOpen
	right := readFromJSON().WindowsOpen.RightWindowOpen
	//CurrentCar := readFromJSON()
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
	drawHello()

}
func changeMark() {
	var MarkChangeTo string
	CurrentCar := readFromJSON()
	fmt.Println("Current Mark title is is: ", readFromJSON().CarMark)
	fmt.Println("Which title you want to swap to?")
	fmt.Scanln(&MarkChangeTo)
	CurrentCar.CarMark = MarkChangeTo
	saveToJSON(CurrentCar)
	fmt.Println("Your awesome name is assigned to current Car", readFromJSON().CarMark)
	drawHello()

}

func changeTrunkCont() {
	var litersToUpload float32
	CurrentCar := readFromJSON()
	fmt.Println("Current free space of Trunk is: ", readFromJSON().TrunkSize)
	fmt.Println("How many liters you want to upload?")
	fmt.Scanln(&litersToUpload)
	CurrentCar.TrunkSize -= litersToUpload
	saveToJSON(CurrentCar)
	fmt.Println("Your things are uploaded successfully! Now current free space of Trunk is", readFromJSON().TrunkSize)
	drawHello()
}

func readFromJSON() PassangerCar {

	file, _ := ioutil.ReadFile("status.json")

	data := PassangerCar{}

	_ = json.Unmarshal([]byte(file), &data)
	return data

}

func saveToJSON(s PassangerCar) {

	file, _ := json.MarshalIndent(s, "", " ")

	_ = ioutil.WriteFile("status.json", file, 0644)
}
