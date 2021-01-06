package main

import (
	"fmt"
	"github.com/pzentenoe/onboarding-golang/tutorial/structs/models"
)

type DriverSQL string

const (
	oracle   DriverSQL = "Oracle"
	postgres DriverSQL = "postgres"
)

func getUrl(driver DriverSQL) string {
	return fmt.Sprintf("url:%s", driver)
}

type siglo int

func main() {

	/*	fmt.Println(getUrl(oracle))

		var siglo20 siglo = 20
		siglo21 := 21

		fmt.Println(siglo20 + siglo(siglo21))*/

	var user1 models.User

	user2 := models.User{
		ID:       "2",
		UserName: "user2",
		Password: "1212121",
	}

	fmt.Println(user1)

	fmt.Println(user2.UserName)
	user2.SetUserName("user3")
	fmt.Println(user2.UserName)

	persona1 := new(models.Persona)
	persona1.Name = "Pablo"
	user3 := &models.User{
		ID:       "3",
		UserName: "user-3",
		Password: "121221",
		Persona:  persona1,
	}

	fmt.Println(user3.Persona)

	user3Json, err := user3.ToJSON()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("User json: %s\n", user3Json)

	var user4 *models.User

	user4, err = user3.FromJSON(user3Json)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(user4.Persona.Name)

}
