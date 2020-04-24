package main

import "fmt"

func main() {
	var targetLen int
	fmt.Println("How many prime numbers you need")
	fmt.Scanln(&targetLen)
	//initial size of initial slice
	n := 2
	//here we check if slice with result has necessary len. If not fill initial slice again and again ..
	for len(removeDevidedNumbers(fillInitialArray(n))) < targetLen {

		removeDevidedNumbers(fillInitialArray(n))
		n++
		//fmt.Println(n)

	}
	fmt.Println(removeDevidedNumbers(fillInitialArray(n)))
}

//filling initial slice by natural numbers
func fillInitialArray(n int) []int {
	var firstSlice []int
	for i := 2; i < n; i++ {
		firstSlice = append(firstSlice, i)

	}

	return firstSlice
}

//remove devided numbers from initial slice. It gets InitialSlice as argument and return result in "b" variable
func removeDevidedNumbers(b []int) []int {

	for z := 0; z < len(b); z++ {

		for i := 1; i < len(b); i++ {

			if b[i] != b[z] && b[i]%b[z] == 0 {
				b = append(b[:i], b[i+1:]...)

			}

		}
	}
	//fmt.Println(len(b))
	return b
}
