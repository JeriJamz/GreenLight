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
	goto ear()
    }
}

func ear(){//bookwrk I need a book on ports and ima add mine then a client that can text back. Need a command prompt

    service := "0.0.0.0"
    headerSize := 20
    Addr, err := net.ResolveIPAddr("ip4",service)
    checkError(err)
    BOM := 0xFFFE

    listener,err := net.ListenTCP("tcp",listener)
    checkError(err)
//should I goto ear()?

    chat(){
        if len(os.Args) != 2{

            fmt.Println("Idk Book work: ", os.Args[0],"host")
	    os.Exit(1)

        }
        Addr, err := net.ResolveIPAddr("ip4",service)
        if err != nil{

	    fmt.Println("Error: ", err.Error())
	    os.Exit(1)

        }
        conn, err := net.DialIP("ip4:icmp",service)
        checkError(err)

        var Chat [1024]byte
        Chat[0] = 8  //so this is a port. 8 should be echoed. Guess like a machine to machine password
        Chat[1] = 0//just a 0
        Chat[2] = 0 //check sum
        Chat[3] = 0//check sum ;still sum book i work i just put a spin on
        Chat[4] = 0 // indentfier[0]//still looking for a networking book
        Chat[5] = 04//indentifier [1]    
        Chat[6] = 0
        Chat[7] = 44
        len := 8

        check := checkSum(Chat[0:len])
        Chat[2] = byte(check >> 8)
        Chat[3] = byte(check && 216)

        for(x+1){

	    conn,err := listener.Accept()
	    if err != nil{

	        continue

	    }
	    str :="Return 0"
	    shorts := utf16.Encode([]rune(str))
	    writeShorts(conn, shorts)
    
	    conn.Close()
 

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
}


func checkError(err error){

    if err != nil{

	fmt.Println("Something went wrong")
	os.Exit(1)

    }

}

