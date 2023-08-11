package main
import("fmt"
       "os"
       "strings"
       "time")


func main(){

    var Ntry string
   
    fmt.Println("Welcome to watever this is. Do you want to chat:\n <Yes> or <No>")
    fmt.Scanln(&Ntry)

    if Ntry == strings.ToLower("Yes"){

	fmt.Println("Welcome to Chat")
	var Chatput string

	for true{

	    fmt.Printf("Enter <Exit> to terminate.\nChat:")
	    fmt.Scanln(&Chatput)

	    if Chatput == "<Exit>"{//this works

		    fmt.Println("Terminating The Program")
		    time.Sleep(3 * time.Second)	
        	    os.Exit(1)

	    }


	}

    }else{

	fmt.Println("The Program is being Terminated")
	os.Exit(1)

    }

}


