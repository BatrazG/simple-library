package main

import "fmt"

func main() {
	user1 := Reader{
		ID:        1,
		FirstName: "Agunda",
		LastName:  "Kokoyti",
		IsActive:  true,
	}

	user1.Deactivate()

	fmt.Println(user1)
}
