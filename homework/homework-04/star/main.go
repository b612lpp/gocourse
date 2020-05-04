/*Написать функцию, которая будет получать позицию коня на шахматной доске,
 а возвращать массив из точек, в которые конь сможет сделать ход.
Точку следует обозначить как структуру, содержащую x и y типа int
Получение значений и их запись в точку должны происходить только с помощью отдельных методов.
В них надо проводить проверку на то, что такая точка может существовать на шахматной доске.
*/

package main

import "fmt"

type horsePosition struct {
	HorixontWay int
	VerticalWay int
}

//CorrectPosition checks if start position is correct
var CorrectPosition bool

//ListPositions contains list of available positions to move
var ListPositions []horsePosition

func (p horsePosition) checkPosition() bool {

	if p.HorixontWay < 8 && p.VerticalWay < 8 {
		CorrectPosition = true
		//fmt.Println(x)
	}

	return CorrectPosition
}

func (p horsePosition) nextPosition() []horsePosition {
	var MoveX = []int{2, 2, -2, -2, 1, 1, -1, -1}
	var MoveY = []int{1, -1, 1, -1, -2, 2, 2, -2}
	for i := 0; i < 8; i++ {
		a := p.HorixontWay + MoveX[i]
		b := p.VerticalWay + MoveY[i]
		if a >= 0 && a < 8 && b >= 0 && b < 8 {
			var c = horsePosition{a, b}
			ListPositions = append(ListPositions, c)
		}

	}
	return ListPositions
}

func main() {
	var z = horsePosition{3, 7}
	if z.checkPosition() == true {
		z.nextPosition()
	}
	fmt.Println(ListPositions)
}
