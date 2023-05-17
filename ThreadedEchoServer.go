//This is over a Threaded Echo Server
package main

import("fmt"
       "os"
       "net")

func main(){

    service := "1201"
    tcpAddr,err := net.ResolveTCPAddr("tcp",service)
    checkError(err)

    ear,err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    for{

	conn, err  := ear.Accept()
	if err != nil{

	    continue

	}

    go handleClient(conn)//go routine

    }

}

func handleClient(conn net.Conn){

    defer conn.Close()//this closes

    var buff [512]byte//buff is now 512 bytes
    for{

	n, err := conn.Read(buff[0:])//this reads from 0 to 511(its 512 bytes 0 is the start or 1)
	if err != nil{

	    return

	}

	fmt.Println(string(buff[0:]))
	_, err2 := conn.Write(buff[0:n])//this will read to n bytes or what ever bytes
	if err2 != nil{

	    return

	}

    }

}

func checkError(err error){

    if err != nil{

	fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	os.Exit(1)

    }

}




