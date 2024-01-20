package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Hakim"
	names[1] = "Asrori"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [3]int{
		90,
		80,
		100,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	values[2] = 70
	fmt.Println(values[2])

}
