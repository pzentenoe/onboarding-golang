package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var numeroInt int
	numeroInt = -2000
	fmt.Println(numeroInt)

	numero2Int := int64(9223372036854775807)
	fmt.Println(numero2Int)

	var numeroUint uint64
	numeroUint = 1000
	fmt.Println(numeroUint)

	numero2Uint := uint(1000)
	fmt.Println(numero2Uint)

	suma := int(numeroUint) - numeroInt
	fmt.Println(suma)

	var numeroFloat float32
	numeroFloat = 13.2
	fmt.Println(numeroFloat)

	numeroFloat64 := float64(100.2)
	fmt.Println(numeroFloat64)

	var isMartes bool
	fmt.Println(isMartes)

	var caracter rune
	caracter = '.'
	fmt.Println("caracter code", caracter)
	fmt.Println("caracter :", strconv.QuoteRune(caracter))

	//Strings
	var palabra string
	palabra = "blabaklb,labla,babl,a"
	fmt.Println(palabra)

	palabras := strings.Split(palabra, ",")
	fmt.Println(palabras[0])

	palabra2 := strings.Join(palabras, ",")
	fmt.Println(strings.ToUpper(palabra2))

	//strconv
	numeroString := "10"
	numeroConverted, err := strconv.Atoi(numeroString)
	if err != nil {
		fmt.Printf("Fallo al convertir : %s", err.Error())
	} else {
		fmt.Println(numeroConverted)
	}

}
