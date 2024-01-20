package main

import "fmt"

func main() {
	person := map[string]string{
		"firstName": "Hakim",
		"lastName":  "Asrori",
	}

	person["title"] = "Programmer"

	fmt.Println(len(person))
	fmt.Println(person)
	fmt.Println(person["firstName"])
	fmt.Println(person["lastName"])

	delete(person, "title")
	fmt.Println(person)
}
