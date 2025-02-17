package main

import "fmt"

func main() {

	var ptr *int
	fmt.Println("Value of pointer is: ", ptr)

	newNum := 70
	newPtr := &newNum

	fmt.Println("Value of new num: ", newNum)
	fmt.Println("Value of new ptr: ", newPtr)

	newPtr1 := *newPtr + 2
	//*newPtr = *newPtr + 2
	fmt.Println("New Value of ptr1: ", newPtr1)
	fmt.Println("Value of newNum after update: ", newNum)
	fmt.Println("Value of newPtr: ", *newPtr)

}
