package main

import "fmt"

func main() {
	var namesEdadMap map[string]int
	namesEdadMap = map[string]int{}
	namesEdadMap["PabloZ"] = 32
	fmt.Println(namesEdadMap)
	fmt.Println(namesEdadMap["PabloZ"])

	animalesMap := make(map[string]string)
	animalesMap["canino"] = "Perro"
	animalesMap["felino"] = "Gato"

	if animal, founded := animalesMap["canino"]; founded {
		formatedString := fmt.Sprintf("Existe el animal %s", animal)
		fmt.Println(formatedString)
	}

	for key, value := range animalesMap {
		fmt.Println("key", key)
		fmt.Println("value", value)
	}

}
