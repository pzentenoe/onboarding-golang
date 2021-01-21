package main

import (
	"fmt"
	"github.com/pzentenoe/onboarding-golang/tutorial/functions/utils"
	"github.com/pzentenoe/onboarding-golang/tutorial/structs/models"
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

	p1 := models.Persona{
		Name:     "Pablo",
		LastName: "Z",
		Age:      32,
	}

	ChangeAge(&p1, 33)
	fmt.Println(p1)

	numero2 := 33
	var numero *int
	numero = &numero2
	fmt.Println(*numero)
	fmt.Println(&numero)

}

func ChangeAge(p *models.Persona, newAge int) {
	p.Age = newAge
}
