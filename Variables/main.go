package main

import "fmt"

const TestVar string = "mehar" //Public

func main() {
	fmt.Println("Variables")

	var username string = "Mehar"
	fmt.Printf("Variable type is: %T \n", username)

	var isLoggedIn bool = true
	fmt.Printf("Variable type is: %T \n", isLoggedIn)
	fmt.Printf("Variable value is: %v \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable type is: %T \n", smallVal)

	var smallFloat float64 = 255.234
	fmt.Println(smallFloat)
	fmt.Printf("Variable type is: %T \n", smallFloat)

	//default values and some aliases
	var newVar int
	fmt.Println(newVar)

	var newVar1 string
	fmt.Println(newVar1)

	//implicit type
	var website = "mehar@gmail.com"
	fmt.Println(website)

	//no var style (not allowed outside a function)
	numberOfUsers := 12345
	fmt.Println(numberOfUsers)
}
