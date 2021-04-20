package main

import "fmt"

var Name string
var Age int

func main() {
	fmt.Scanln(&Name)
	fmt.Scanln(&Age)
	fmt.Println("Имя:", Name)
	if Age < 7 {
		fmt.Println("Вам", Age, ".Идите в садик")
	} else if (Age >= 7) && (Age < 18) {
		var classes int
		classes = Age - 6
		fmt.Println("Вам", Age, ". Вам в", classes, "класс")
	} else {
		fmt.Println("Вам", Age, ". Пора идти работать")
	}
}
