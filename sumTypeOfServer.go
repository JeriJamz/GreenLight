//some bookwork in here
import("fmt"
       "os"
       "net"
       "net/http"
       "sync"
       "log"
       "encoding/json"
       "encoding/gob"
       "bufio")



type Usr struct{

    UsrName UserName
    Password Password
    Email []Email

}


type UsrName struct{

    usrName string
    Email []Email

}

type Password struct{

    Password string

}

type Email struct{

    Address string

}

func main(){

    fmt.Println("Welcome\nAre you a new user?")
    response = scanln()

    if("Yes"){

	goto NewUsrLogin()
	return 0
    }
    else if(){

	goto Login()

    }
    else if(){
    
	goto ear()

    }
}


func NewUsrLogin(){

    NewUser := Usr{UsrName:UserName{usrName:"",Email:""},
		   {Password:""}}

    WriteNewUsrLogin(fileName, key interface){

        fmt.Println("Hello New User. Please Type in your User Name")
    
        outFile, err := os.Write("Login.txt")
        checkError(err)
        encoder := gob.NewEncoder(outFile)
        err = encoder.Encode(key)
        checkError(err)
        fmt.Println("Please Enter Your UserName:")
	newUsrName := scanln()
	outFile := newUsrName
	fmt.Println("Please Enter Your Password")
	newPassword := scanln()
	outfile := newPassword
	outFile.Close()

    }
}

login{

    fmt.Println("Please enter your. User Name or Email")
    UN := scanln()
    fmt.Println("Please Enter Your Password.")
    PW := scanln()	

    //need to add a file write function right here
    Logn, err := os.Open("Login.txt")

    if(UN != Logn){

	fmt.Println("Your Username is not in here\n"
		    "Try again")
	else if(UN == Logn){//can I do this?

	    fmt.Println("Ok, Welcome {UN}")

	}    
	goto login

	}
    else if (PN != Logn){

	fmt.Println("Thats not the secrect passphrase")	    	
	goto login
    }
    else(){

	fmt.Println("Welcome...")

    }
}

func ear(){

    service := "0.0.0.0:1244"
    tcpAddr, err := net.ResolveTCP("tcp",service)
    checkError(err)

    lister,err := net.ListenTCP("tcp",listener)
    checkError(err)

    for{

	conn,err := listenr.Accept()
	if err != nil{

	    continue

	}
	str :="this should get encoded"
	shorts := utf16.Encode([]rune(str))
	writeShorts(conn, shorts)

	conn.Close()

    }
    func writeShorts(conn net.Conn, shorts []uint16){

	var bytes [2]byte

	bytes[0] = BOM >> 8
	bytes[1] = BOM & 255
	_,err:= conn.Write(bytes[0:])
	if err != nil{

	    return

	}

    }
}


func checkError(err error){

    if err != nil{

	fmt.Println("Something went wrong")
	os.Exit(1)

    }

}

