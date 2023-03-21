package main

import (
	"fmt"
	"log"

	gonimbus "github.com/szymon676/go-nimbus"
)

func main() {
	engine := gonimbus.New()

	engine.Get("/people", handleGetUsers)
	engine.Post("/people", handleCreateUser)
	engine.Put("/people/:username", handleUpdateUser)

	engine.Serve("4000")
}

func handleCreateUser(c gonimbus.Context) {
	var bindPerson *person
	if err := c.BindJSON(&bindPerson); err != nil {
		log.Fatal(err)
	}
	people = append(people, *bindPerson)
	c.String(201, "user created successfully")
}

func handleGetUsers(c gonimbus.Context) {
	var users []person
	for _, user := range people {
		users = append(users, user)
	}
	c.Return(200, users)
}

func handleUpdateUser(c gonimbus.Context) {
	var bindPerson *person
	username := c.Param("username")
	fmt.Println(username)
	for i, person := range people {
		if person.Name == username {
			if err := c.BindJSON(&bindPerson); err != nil {
				c.String(400, err.Error())
				return
			}
			people[i] = *bindPerson
			c.String(202, "user updated successfully")
			return
		}
	}
	c.String(404, "user not found")
}

var people = []person{}

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
