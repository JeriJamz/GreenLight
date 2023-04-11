/*
This Is supposed to find the Network Mask
*/

package main

import(
	"fmt"
	"net"
	"os"
	"strconv"
)

func main(){

	if len(os.Args) != 4{

		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr ones bits\n", os.Args[0])
		os.Exit(1)

	}
	dotAddr := os.Args[1]
	ones, _ := strconv.Atoi(os.Args[2])//this is a str conversion I think. 
	bits, _ := strconv.Atoi(os.Args[3])//Its the reason its gone read hex

	addr := net.ParseIP(dotAddr)
	if addr == nil{

		fmt.Println("Invalid Address")
		os.Exit(1)

	}
	mask := net.CIDRMask(ones, bits)//CIDR checks for iPv4 
	network := addr.Mask(mask)// 
	fmt.Println("Address is ", addr.String(),
				"\nMask is Length is ", bits,
				"\nLeading ones count is ",ones,
				"\nMask is (hex) ", mask.String(),
				"\nNetwork is ", network.String())
	os.Exit(0)

}


