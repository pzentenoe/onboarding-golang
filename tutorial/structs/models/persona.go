package models

type Persona struct {
	Name     string `json:"name" xml:"name"`
	LastName string `json:"last_name" xml:"last_name"`
	Age      int    `json:"age" xml:"age"`
}
