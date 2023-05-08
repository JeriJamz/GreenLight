package main //I jus wanna print hello world.

import(
	"fmt"
      )

func main(name string) string{

	message := fmt.Sprintf("Hi, %v. Welcome!", name)//field spec(%v)?
	return message	

}



