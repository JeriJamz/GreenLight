package main

import("fmt")

func main(){

	var x string
	x = "A"
	fmt.Println(x)
	x = "B"
	fmt.Println(x)
	x = x + " A"
	fmt.Println(x)
	z := "C"
	fmt.Println(z)
	A := x + z
	fmt.Println(A)
	var B = "This is B"
	fmt.Println(B)
	var C string = B +"," + A
	fmt.Println(C)
}





