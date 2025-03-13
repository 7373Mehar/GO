package main

import (
	"fmt"
	"strings"
)

func main() {

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Watermelon"
	fruitList[2] = "Banana"
	fruitList[3] = "Orange"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("length of the list is: ", len(fruitList))

	fruitStr := strings.Join(fruitList[:], ", ")
	fmt.Println(fruitStr)

	var vegList = [3]string{"Potato", "Ladyfinger", "Beans"}
	fmt.Println("Veg List is: ", vegList)
}
