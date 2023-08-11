package main
import("fmt"
       "os"
       "strings")


func main(){

    var Ntry string
   
    fmt.Println("Welcome to watever this is. Do you want to chat:\n <Yes> or <No>")
    fmt.Scanln(&Ntry)

    if Ntry == strings.ToLower("Yes"){

	fmt.Println("Welcome to Chat")
	var Chatput string

	for true{

	    fmt.Println("Enter <Exit> to terminate.\nChat:")
	    fmt.Scanln(&Chatput)

	    if Chatput == "<Exit>"{//this works

        	    os.Exit(1)

	    }


	}

    }	

}


