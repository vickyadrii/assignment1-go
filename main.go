package main

import (
	"fmt"
	"os"
)

type Friend struct {
	ID         string
	Name       string
	Address    string
	Occupation string
	Reason     string
}

func SearchFriendById(id string) {
	friends := []Friend{
		{
			ID:         "F1",
			Name:       "Naufal",
			Address:    "Jln Akasia 1",
			Occupation: "Web Developer",
			Reason:     "Ingin memperdalam ilmu golang dan ilmu padi",
		},
		{
			ID:         "F2",
			Name:       "Jindan",
			Address:    "Jln Zebra",
			Occupation: "Backend Developer",
			Reason:     "Ingin mendapatkan relasi dan mendapatkan pekerjaan yang menggunakan bahasa golang",
		},
		{
			ID:         "F3",
			Name:       "Akwan",
			Address:    "Jln Karajalemba",
			Occupation: "Software Engineer",
			Reason:     "Ingin memperdalam bahasa pemrograman golang",
		},
		{
			ID:         "F4",
			Name:       "Dharma",
			Address:    "Jln Diponegoro",
			Occupation: "Frontend Developer",
			Reason:     "Mencoba coba bahasa pemrograman golang",
		},
		{
			ID:         "F5",
			Name:       "Ryan",
			Address:    "Jln Elang",
			Occupation: "Software Engineer",
			Reason:     "Ingin membuat beberapa project menggunakan golang",
		},
	}

	var isFound bool

	for _, value := range friends {
		if id == value.ID {
			fmt.Printf("id: %s\n", value.ID)
			fmt.Printf("name: %s\n", value.Name)
			fmt.Printf("address: %s\n", value.Address)
			fmt.Printf("occupation: %s\n", value.Occupation)
			fmt.Printf("reason: %s\n", value.Reason)

			isFound = true
		}
	}

	if !isFound {
		fmt.Printf("friend with ID %s was not found\n", id)
	}
}

func main() {
	var arguments []string = os.Args

	if len(arguments) < 2 {
		fmt.Println("Please input the friend's ID")
		return
	}

	if len(arguments) > 2 {
		fmt.Println("Please input 1 ID only")
		return
	}

	SearchFriendById(arguments[1])
}
