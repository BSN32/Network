/*
Json Dosyadan okuyup decode yapma işlemi
7/03/2022
Burak Şen
*/
package jsontest

import (
	"encoding/json"
	"fmt"
	"os"
)

// User struct which contains a name
// a type and a list of social links
type Users struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Age  int    `json:"Age"`
}

func Demo1() {

	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	var users []Users

	var userDecoder *json.Decoder = json.NewDecoder(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	err = userDecoder.Decode(&users)
	if err != nil {
		fmt.Println(err)
	}

	for i, user := range users {
		fmt.Println(i + 1)
		fmt.Println("user name:", user.Name)
		fmt.Println("user type:", user.Type)
		fmt.Println("user age:", user.Age)
	}
}
