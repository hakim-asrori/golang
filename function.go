package main

import "fmt"

func main() {
	name()
	age(22)
	fmt.Println(calculateAdd(10, 2))

	name, _, title := person("Programmer")
	fmt.Println("Nama", name, "Title", title)
}

func name() {
	fmt.Println("Hakim Asrori")
}

func person(title string) (string, int, string) {
	return "Hakim Asrori", 22, title
}

func age(number int) {
	fmt.Println("Umur", number)
}

func calculateAdd(number1 int, number2 int) int {
	return number1 + number2
}
