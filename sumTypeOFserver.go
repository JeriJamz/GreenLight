import("fmt"
       "os"
       "net"
       "net/http"
       "sync"
       "log"
       "encoding/json"
       "encoding/gob")



type Usr struct{

    UsrName UsrName
    Password Password
    Email []Email

}


type UsrName struct{

    usrName string

}

type Password struct{

    Password string

}

type Email struct{

    Address string

}

func main(){

    fmt.Println("Welcome\nPlease Enter your User Name and Password")
    response = scanln()

    

    if("Yes"){

	goto NewUsrLogin()
	return 0
    }
    else(){

	

    }
}

func NewUsrLogin(fileName, key interface){

    fmt.Println("Hello New User. Please Type in your User Name")
    
    outFile, err := os.Create(fileName)
    checkError(err)
    encoder := gob.NewEncoder(outFile)
    err = encoder.Encode(key)
    checkError(err)
    outFile.Close()

}

func checkError(err error){

    if err != nil{

	fmt.Println("Something went wrong")
	os.Exit(1)

    }

}

