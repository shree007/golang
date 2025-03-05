package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name,omitzero"`
	Age   int    `json:"age,omitzero"`
	Email string `json:"email,omitzero"`
}

func main() {
	user := User{Name: "Alice"}
	data, _ := json.MarshalIndent(user, "", " ")
	fmt.Println(string(data))
}
