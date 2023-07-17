//mo book work until i figure how to make beefyservers FTP
//let me figure a .lower() and ima update this
import("fmt"
       "os"
        "net")

const(DIR = "DIR
      CD = "CD"
      PWD = "PWD")

func main(){

    service = "0.0.0.0:1201"//IPaddress
    tcpAddr, err := net.ResolveTCP("tcp",service)//a port?
    checkError(err)

    ear,err := net.ListenTCP("tcp",tcpAddr)//this is listening for request coming from the Addr
    checkError(err)

    for{

	conn,err := listener.Accept()
	if err != nil{

	    continue

	}

        go handleClinet(conn)

    }

}

func handleClient(conn net.Conn){

    defer conn.Close()

    var buf [512]byte
    for{

	n,err := conn.Read(buff[0:])
	if err != nil{

	    conn.Close()
	    return

	}

    s := string(buf[0:n])//decodes request
    if s[0:2] == CD{//the r

	chdir(conn,s[3:])

    }
    else if s[0:3] == DIR{

	dirList(conn)//so conn is the files ima assume

    }
    else if s[0:3] == PWD{

	pwd(conn)


    }

}

func chdir(conn net.Conn, s string){

    if os.Chdir(s) == nil{

	conn.Write([]byte("OK"))

    }else{
	conn.Write([]byte("ERROR"))
     }

}

func pwd(conn net.Conn){

    s, err := os.Getwd()
    if err := nil{

	conn.Write([]byte(""))
	return

    }
    conn.Write([]byte(s))
}

func dir_list(){
//this sends a blank line to terminate the function
    defer conn.Write([]byte("\r\n"))

    dir,err := os.Open(".")
    if err != nil{

	return

    }

    name,err := dir.Readdirnames(-1)
    if err != nil{

	return

    }

    for _,nm:= range names{

	conn.Write([]byte(nm + "\r\n"))

    }

}

func checkError(err, error){

    if err != nil{
	
	fmt.println("i didnt compile this so it might not be you", err.Error())
	os.Exit(1)

    }

}
