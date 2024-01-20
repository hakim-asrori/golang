package main

import "fmt"

func main() {
	const firstName string = "Hakim"
	const lastName = "Asrori"

	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		mom string = "Rini Julianti"
		dad        = "Kusen"
	)
	fmt.Println(mom)
	fmt.Println(dad)
}
