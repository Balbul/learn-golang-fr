package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	jsonFromApi := `
[
	{
  "name": "John",
  "age": 30,
  "city": "New York"
  },
	{
  "name": "Jane",
  "age": 31,
  "city": "London"
	}
]`

	var users []User

	err := json.Unmarshal([]byte(jsonFromApi), &users)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Printf("json: %+v\n", users)

	var myStruct []User

	var user_one User
	user_one.Name = "John"
	user_one.Age = 30
	user_one.City = "New York"

	myStruct = append(myStruct, user_one)

	jsonFromSlice, err := json.MarshalIndent(myStruct, "", "  ")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Printf("json: %+v\n", string(jsonFromSlice))

}
