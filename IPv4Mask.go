//This is sum mo book work
//IPv4 Mask
package main

import(//Gone need these
  "fmt"
  "os"//I/O from the whole System. I think
  "net"//how to deal with Socket stuff
)

func main(){

    if len(os.Args) != 2{//IF Os packages are less than 2

	fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n", os.Args[0])//this gone print the error
	os.Exit(1)//exit if false?

    }
    dotAddr := os.Args[1]//Start the addr at slot 2[1]
    addr := net.ParseIP(dotAddr)//Parse the IP. Store the bite as an array in dotaddr. Then Put that array in adder
    if addr == nil{//If the addr pull up an error. Not it in the right formatt

	fmt.Println("Invalid address")
	os.Exit(1)

    }

    mask := addr.DefaultMask()//Wateva Default mask is its in mask.
    network := addr.Mask(mask)//Network is a vari for the .Mask func thats holding a hand writtten mask func
//The Abstraction OOOOUUUUUUUU
    ones, bits := mask.Size()//Ima assume this is a map of some kind
    fmt.Println("Address is, ", addr.String(),
		"\nDefault mask length is, ", bits,
		"\nLeading ones count is, ", ones,
		"\nMask, in hex, is, ", mask.String(),
		"\nNetwork is, ", network.String())
    os.Exit(0)//if true exit

}


