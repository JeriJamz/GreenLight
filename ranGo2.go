//im still workin on takin input. I dont think it takes this much.
package main

import("fmt"
       "bufio"
       "os"
)

func main(){

	input := bufio.NewScanner(os.Stdin)

	line := input.Text()
	fmt.Println(line)

}


