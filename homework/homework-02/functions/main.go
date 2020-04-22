package main

import "fmt"

func main() {
	drawWelcome()

}

func drawWelcome() {
	fmt.Println("---------------------------------")

	var calledFunc string
	fmt.Println("We have following functions\n 1. Check if number is even \n 2. Check if numder divided by 3\n 3. Show Fibo numbers\n 4. Exit\nInput number what you want to run")
	fmt.Scan(&calledFunc)

	switch calledFunc {
	case "1":
		checkEven()
	case "2":
		devidedThree()
	case "3":
		fibo()
	case "4":
		break
	default:
		fmt.Println("Undifined action")
		drawWelcome()
	}
}

func checkEven() {
	var number int
	fmt.Println("Input the number you want to check if it is even")
	fmt.Scanln(&number)
	if number%2 == 0 {
		fmt.Println("Your number is even")
	} else {
		fmt.Println("Your number is odd")
	}
	drawWelcome()

}

func devidedThree() {
	var number int
	fmt.Println("Input the number you want to check if it is divided by 3 without remainder")
	fmt.Scanln(&number)
	if number%3 == 0 {
		fmt.Println("Your number is divided by 3 without remainder")
	} else {
		fmt.Println("nope")
	}
	drawWelcome()

}

func fibo() {
	/*fibo := []float64{0, 1}
	for n := 2; n < 100; n++ {
		fibo = append(fibo, fibo[n-1]+fibo[n-2])
		//fmt.Println(n, fibo[n])
	}
	for i := 0; i < 100; i++ {

		fmt.Println(i+1, fibo[i])

	} */
	var x float64
	var y float64 = 1

	for i := 1; i < 101; i++ {
		x += y
		y = x - y
		fmt.Println(x)
	}
	drawWelcome()

}
