import("fmt"
       "os"
       "net"
       "net/http"
       "sync"
       "log"
       "encoding/json"
       "encoding/gob")



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
    else(){

	goto Login

    }

    login{

	fmt.Println("Please enter your. User Name or Email")
	UN = scanln()
	fmt.Println("Please Enter Your Password.")
	PW = scanln()	

    }

}


func NewUsrLogin(){

    NewUser := Usr{UsrName:UserName{usrName:"",Email:""},
		   {Password:""}}

    WriteNewUsrLogin(fileName, key interface){

        fmt.Println("Hello New User. Please Type in your User Name")
    
        outFile, err := os.Create(fileName)
        checkError(err)
        encoder := gob.NewEncoder(outFile)
        err = encoder.Encode(key)
        checkError(err)
        outFile.Close()

    }
}

func checkError(err error){

    if err != nil{

	fmt.Println("Something went wrong")
	os.Exit(1)

    }

}

