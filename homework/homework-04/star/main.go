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

func (p horsePosition) nextPosition() []horsePosition {
	var c horsePosition
	var listPositions []horsePosition
	var positionDelta = []horsePosition{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {1, -2}, {1, 2}, {-1, 2}, {-1, -2}}

	for i := 0; i < 8; i++ {
		c.HorixontWay = p.HorixontWay + positionDelta[i].HorixontWay
		c.VerticalWay = p.VerticalWay + positionDelta[i].VerticalWay
		if c.checkCorrect() == true {
			listPositions = append(listPositions, c)
		}

	}
	return listPositions
}

func (p horsePosition) checkCorrect() bool {
	var r bool
	if p.HorixontWay < 8 && p.VerticalWay < 8 && p.HorixontWay >= 0 && p.VerticalWay >= 0 {
		r = true
	}
	return r
}

func main() {

	fmt.Println(inputData().nextPosition())
}

//function gets input values and generate position necessary type
func inputData() horsePosition {
	var x, y int
	var c horsePosition
	fmt.Println("input x")
	fmt.Scanln(&x)
	fmt.Println("input y")
	fmt.Scanln(&y)
	if x < 8 && y < 8 {
		c = horsePosition{x, y}

	} else {
		fmt.Println("wrong data")
		inputData()

	}
	return c

}
