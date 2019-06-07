package main

import(
	"fmt"
)

func main() {
	var name string
	var height,weight float64
	
	var age int = 30               // set the value directly
	company := "Pivotal"	   	   // shortcut to set value directly

	name = "Matt Song"
	// age = 18
	height = 180+20
	weight = 65 * 2

	fmt.Println("Name is", name);
	fmt.Println("Age is", age);
	fmt.Println("Height is", height);	
	fmt.Println("Weight is", weight);
	fmt.Println("I am working at", company);
}
