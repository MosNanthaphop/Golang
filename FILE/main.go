package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	/*
		s := "hello สวัสดี"
			for i := 0; i <= len(s); i++ {
				fmt.Printf("%c %X\n", s[i], s[i])
			}
	*/

	/*
		for _, ch := range s {
			fmt.Printf("%c ", ch)
		}
	*/
	/*
		content, err := os.ReadFile("text.txt")
		if err != nil {
			fmt.Println("Error reading file")
			return
		}
		fmt.Println(string(content))

		newContent := []byte("Hello World!")
		err = os.WriteFile("newtext.txt", newContent, os.FileMode(0644))
		if err != nil {
			fmt.Println("Error writing file")
		}
	*/
	/*
		content, err := os.ReadFile("car1.jpg")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(len(content))

		err = os.WriteFile("new_car.jpg", content, os.FileMode(0644))
		if err != nil {
			fmt.Println(err)
			return
		}
	*/
	// อ่านไฟล์ JSON
	content, err := os.ReadFile("person.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var person Person
	err = json.Unmarshal(content, &person)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	fmt.Printf("Person: %+v\n", person)
	fmt.Println("name: ", person.Name)

	newPerson := Person{Name: "Mslick", Age: 20}
	newContent, err := json.Marshal(newPerson)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = os.WriteFile("newPerson.json", newContent, os.FileMode(0644))
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
