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

	if animal, ok := animalesMap["canino"]; ok {
		fmt.Printf("Existe: %s", animal)
	}

	for key, value := range animalesMap {
		fmt.Println(key)
		fmt.Println(value)
	}



}
