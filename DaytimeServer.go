//Ya already kno still goin over book work but I think this is gonna show me how to set ports
package main

import("fmt"
       "net"
       "os"
       "time")

func main(){

    service := ":1200"//Think this is the port
    tcpAddr, err := net.ResolveTCPAddr("tcp", service)
    checkError(err)

    ear,err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)//always check for errors

    for{

	conn, err := ear.Accept()
	if err != nil{

	    continue

	}

    daytime := time.Now().String()
    conn.Write([]byte(daytime))
    conn.Close()

    }

}

func checkError(err error){

    if err != nil{

	fmt.Fprintf(os.Stderr, "There was a lil error: %s ", err.Error())
	os.Exit(1)

    }

}

func (c *TCPConn) SetDeadline(t time.Time) error