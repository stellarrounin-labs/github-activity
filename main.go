package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id       int    `json:id`
	Name     string `json:name`
	Username string `json:username`
	Email    string `json:email`
}

func main() {
	client := &http.Client{}

	resp, err := client.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var users []User

	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		fmt.Println("Read error: ", err)
		return
	}

	for _, user := range users {
		fmt.Printf("ID: %d | Name: %s\n", user.Id, user.Name)
	}
}
