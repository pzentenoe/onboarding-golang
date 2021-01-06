package main

import (
	"fmt"
	"github.com/pzentenoe/onboarding-golang/tutorial/functions/utils"
)

func main() {

	page, size, err := utils.ConvertPaginationParams("1", "10")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("page: %d, size: %d\n", page, size)

	sumar := func(num1, num2 int) int {
		return num1 + num2
	}

	func() {
		fmt.Println("Funcion anonima")
	}()

	fmt.Println(sumar(10, 8))

}
