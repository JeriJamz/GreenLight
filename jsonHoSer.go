//this is an JSON echo server
package main

import("fmt"
       "encoding/json"
       "net"
       "os")

type Person struct{

    Name Name
    Email []Email
    Age Age

}

type Name struct{

    Family string
    First string
    NickName string

}

type Email struct{

    Type string
    Address string

}

type Age struct{

    Candles string
    Birthday string

}

func(p Person) String() string{

    s := p.Name.First + " " + p.Name.Family + " " + p.Name.NickName
    for _,v := range p.Email{

	s +="\n" + v.Type + " : " + v.Address

    }
    d := p.Age.Birthday + ". You are: " + p.Age.Candles + " Years old."

    return s
    return d

}

func main(){

    service := "0.0.0.0:1200"
    tcpAddr,err := net.ResolveTCPAddr("tcp", service)
    checkError(err)
   
    listener, err := net.ListenTCP("tcp",tcpAddr)
    checkError(err)

    for{

	conn,err := listener.Accept()
	if err != nil{

	    continue

	}
	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)    

	for n := 0; n < 10; n++{

	    var person Person
	    decoder.Decode(&person)
	    fmt.Println(person.String())
	    encoder.Encode(person)

	}
	conn.Close()
    }

}


func checkError(err error){

    if err != nil{

	fmt.Println("Fatal Error ", err.Error())
	os.Exit(1)

    }

}

