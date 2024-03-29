//ASN1 Daytime Server

package main

import("encoding/asn1"
       "fmt"
       "net"
       "os"
       "time")

func main(){

    service := ":1200"
    tcpAddr, err := net.ResolveTCPAddr("tcp", service)
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    for{

	conn, err := listener.Accept()
	if err != nil{

	    continue

	}
	daytime := time.Now()
	//Ignore Return Network errors
	mdata, _ := asn1.Marshal(daytime)
	conn.Write(mdata)
	conn.Close()//finished

    }

}

func checkError(err error){

    if err != nil{

	fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())

    }

}

