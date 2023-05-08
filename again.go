package main

import(
	"fmt"
	"os"
	)

func main(){

  var s, sep string
  name := "Jeri"

    fmt.Println("Hello World\n",
		"1 + 1 = ",1 +1,// see this syntax
		"\nAlso greetings from, ", name)//I need to gts
				//that math expression is a go
	fmt.Println(len("\nWat are Jooooeeees"),
			"\nCan I ball " + "Can I chill?")
	fmt.Println("On my Baby"[1])//idk ill debug when I wake`
	for i := 1; i < len(os.Args); i++{

		s += sep + os.Args[i]	
		sep = " "

	}

	fmt.Println(s)
}


