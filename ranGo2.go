//im still workin on takin input. I dont think it takes this much.
package main

import("fmt"
       "os"
       "bufio"
)

func main(){

  input :=  bufio.NewScanner(os.Stdin)
  Usr := input.Text()
  //var s string
    fmt.Println("Im trying to take input and print it out")
	Usr//The problem here. Idk how to get the input yet.(User) Utility?
	fmt.Println(Usr)


}


