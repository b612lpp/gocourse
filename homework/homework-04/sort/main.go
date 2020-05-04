/*Создать тип, описывающий контакт в телефонной книге.
Создать псевдоним типа телефонной книги (массив контактов) и реализов*******+ать для него интерфейс Sort{}.
map[KeyType]ValueType*/
package main

import (
	"fmt"
	"sort"
)

type Contact struct {
	Name  string
	Phone int
}

type PhoneBook []Contact

func (b PhoneBook) Len() int {
	return len(b)
}

func (b PhoneBook) Less(i, j int) bool {
	return b[i].Phone < b[j].Phone
}

func (b PhoneBook) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	var YellowPages = PhoneBook{{Name: "Mr. WHite", Phone: 5432}, {Name: "Mr. Red", Phone: 2138}, {Name: "Mr. Blue", Phone: 9621}}
	fmt.Println("Default list", YellowPages)
	sort.Sort(YellowPages)
	fmt.Println("Sorted list", YellowPages)
	sort.Sort(sort.Reverse(YellowPages))
	fmt.Println("Reversed sort", YellowPages)
	//smth from the net
	sort.SliceStable(YellowPages, func(i, j int) bool { return YellowPages[i].Name < YellowPages[j].Name })
	fmt.Println("Sorted by name", YellowPages)

}
