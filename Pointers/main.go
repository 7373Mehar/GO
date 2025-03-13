package main

import "fmt"

func main() {

	var ptr *int
	//NIL Value as it is only initialized here.
	fmt.Println("Value of pointer is: ", ptr)

	newNum := 70
	newPtr := &newNum
	//Passing memory address.

	fmt.Println("Value of new num: ", newNum)                  //Should print 70
	fmt.Println("Value of new ptr: ", newPtr)                  // Will print memory address
	fmt.Println("Actual value at the memory address", *newPtr) //70

	//newPtr1 := *newPtr + 2 // Will add 2 to the initial value.
	//fmt.Println("New Value of ptr1: ", newPtr1)           //72

	*newPtr = *newPtr + 2
	fmt.Println("Value of newNum after update: ", newNum) //70
	fmt.Println("Value of newPtr: ", *newPtr)             //72

}
