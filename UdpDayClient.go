//Ok im starting to get a better grip on Servers. This is a client daytime server.
//BOOKWORK
package main

import("fmt"
       "os"
       "net")

func main(){

    if len(os.Args) != 2{

	fmt.Fprintf(os.Stderr, "Error, Usage %s: host:port", os.Args[0])
	os.Exit(1)

    }

    service := os.Args[1]

    updAddr, err := net.ResolveUDPAddr("udp", service)
    checkError(err)

    conn, err := net.DialUDP("udp", nil, updAddr)
    checkError(err)

    _, err = conn.Write([]byte("anything"))
    checkError(err)

    var buff [512]byte
    n, err := conn.Read(buff[0:])
    checkError(err)

    fmt.Println(string(buff[0:n]))

    os.Exit(0)

}

func checkError(err error){

    if err != nil{

	fmt.Fprintf(os.Stderr, "Error: ", err.Error())
	os.Exit(1)

    }

}


