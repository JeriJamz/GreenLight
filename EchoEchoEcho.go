package main

import("fmt"
	   "os")

func main(){

	var s, sep string//var decla
	for i :=1 ; i < len(os.Args) ; i++{ //for loop condi

		s += sep + os.Args[i]// S is decla with instruct
		sep = " "//sep is decla wit instruct. It is time to print

	}

	fmt.Println(s)

}




