package main

import "fmt"

func main() {

	//Array
	/*	var listaMeses [6]string
		listaMeses = [6]string{"enero", "febrero", "marzo", "abril", "mayo", "Junio"}	*/
	var listaMeses []string
	listaMeses = []string{"enero", "febrero", "marzo", "abril", "mayo", "Junio"}

	//Slices
	listaDias := make([]string, 0)
	listaDias = append(listaDias, "Lunes")
	listaDias = append(listaDias, "Martes", "Miercoles", "Jueves", "Viernes")

	fmt.Println(len(listaMeses))
	fmt.Println(len(listaDias))
	fmt.Println(cap(listaDias))

	for i := 0; i < len(listaMeses); i++ {
		fmt.Println(listaMeses[i])
	}

	fmt.Println("-------")

	for _, mes := range listaMeses {
		fmt.Println(mes)
	}

	fmt.Println("--------")
	listaMeses2 := []string{"Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

	listaMeses2 = append(listaMeses, listaMeses2...)

	for _, mes := range listaMeses2 {
		fmt.Println(mes)
	}

	i := 0

	//For infinito
	for {
		fmt.Println(i)
		if i == 5 {
			break
		}
		i++
	}

	isLessThan50 := true

	//Similar to (do while)
	j := 0
	for isLessThan50 {
		if j >= 50 {
			fmt.Println("Es mayor que 50")
			isLessThan50 = false
		}
		j++
	}

}
