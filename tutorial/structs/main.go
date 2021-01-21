package main

import (
	"encoding/json"
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

	var user1 *models.User
	fmt.Println(user1)

	user2 := models.User{
		ID:       "2",
		UserName: "user2",
		Password: "1212121",
	}

	fmt.Println(user2.UserName)
	user2.SetUserName("user3")
	fmt.Println(user2.UserName)

	persona1 := new(models.Persona)

	var p1 *models.Persona
	fmt.Println("p1", p1)
	p2 := &models.Persona{
		Name:     "qasas",
		LastName: "asas",
		Age:      0,
	}
	fmt.Println("p2", p2)
	p3 := new(models.Persona)
	fmt.Println("p3", p3)

	persona1.Name = "Pablo"



	p2Bytes, err := json.Marshal(&p2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p2Bytes)
	fmt.Println(string(p2Bytes))

	stringJSON := "{\"name\":\"hector\",\"last_name\":\"varas\",\"age\":30}"

	var hector *models.Persona
	errUnmarshall := json.Unmarshal([]byte(stringJSON), &hector)
	if errUnmarshall != nil {
		fmt.Println(errUnmarshall)
		return
	}
	fmt.Println(hector.Name)
	fmt.Println(hector.Age)


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
