//this is sum ole book work still trynna learn network programming
//JSON is weird Save my rant for another program
package main

import("fmt"
       "encoding/json"
       "os"
       "net"
       "bytes"
       "io")

type Person struct{

    Name Name
    Email []Email //apparently when using JSON from Go.Makes this a base64 string. This lil array
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
    for _, v := range p.Email{

	s += "\n" + v.Type + " : " + v.Address

    }
    d := p.Age.Birthday + ". You are: " + p.Age.Candles + " Years Old"

    return s
    return d

}

func main(){

    person := Person{

	Name: Name{Family:"Hunt", First:"Jerimaine"},
        Age: Age{Birthday:"Jan, 9th 2000.", Candles:"23"},
	Email:[]Email{Email{Type:"Home", Address:"jerimainehuntjr@gmail.com"}}}
    
    
    if len(os.Args) != 3{

	fmt.Println("Usage: ", os.Args[0],"host:port")
	os.Exit(1)

    }
    service := os.Args[1]

    conn,err := net.Dial("tcp", service)
    checkError(err)
    encoder := json.NewEncoder(conn)
    decoder := json.NewDecoder(conn)

    for n := 0; n < 10; n++{//this gonna repeat the message 10 times

	encoder.Encode(person)
	var newPerson Person
	decoder.Decode(&newPerson)
	fmt.Println(newPerson.String())

    }
    os.Exit(0)
}

func checkError(err error){

    if err != nil{

	fmt.Println("Fatal Error ", err.Error())
        os.Exit(1)

    }

}

func readFully(conn net.Conn)([]byte, error){

    defer conn.Close()
    result := bytes.NewBuffer(nil)
    var buf [512]byte
    for{

	n, err := conn.Read(buf[0:])
	result.Write(buf[0:n])
	if err != nil{

	    if err == io.EOF{

		break

	    }
	    return nil, err
        }

    }
    return result.Bytes(),nil
}



