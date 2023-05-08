//this works its jus gonna print blank. The Variables are Whitespaces.
package main

import("fmt"
	"os"
)

func main(){

  s, sep := "",""

    for _, arg := range os.Args[1:]{

	s += sep + arg
	sep = " "

	}

    fmt.Println(s)

}


