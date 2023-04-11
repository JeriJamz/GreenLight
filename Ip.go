/*This is sum book work
Some Its loooks pretty cool
Some how Im teaching myself Go and Networkin at the same time
*/

package main//need this to run and I/O Algo

import(

	"fmt"//this is the std formatting for string
	"net"//this get you to the internet
	"os"//idk but im betting os stuff

)

func main(){

	if len(os.Args) !=2{//if len(args from os) if its doesnt equal 2

		fmt.Fprintf(os.Stderr,  "Usage : %s Ip-addr\n", os.Args[0])/*%s is a string field spec os.Args is now an array that holds bytes I think
			Computers are weird. Ip address are really bytes that refers to something that can connect to the internet*/
		os.Exit(1)

	}
	
	name := os.Args[1]

	addr := net.ParseIP(name)
	if addr == nil{//nil means 0

		fmt.Println("Invalid address")

	}else{//this is pretty cool formatting

		fmt.Println("The address is ", addr.String())// prints the addr as a string value

	}
	os.Exit(0)//this how to break out the OS package

}




