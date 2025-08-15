package jsonstudies

import (
	"encoding/json"
	"fmt"
)

func main() {
	personJson := []byte(`{
		"name": "Vinicius",
		"surname": "Takedi",
		"age": 22,
		"job_role": "Software Engineer"
	}`)

	type Person struct {
		Name    string `json:"name"`
		Surname string `json:"surname"`
		Age     int    `json:"age"`
		JobRole string `json:"job_role"`
	}

	fmt.Printf("%T %v \n", personJson, personJson)

	var person Person
	err := json.Unmarshal(personJson, &person)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%T %v \n", person, person)

	backToJson, err := json.Marshal(person)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%T %v \n", backToJson, backToJson)
}
