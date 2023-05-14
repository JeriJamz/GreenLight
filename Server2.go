//This is some bookwork but I wanna add the Server.go wit this
//Its for TCP Clients to get header info
package main

import("fmt"
       "os"
       "io/ioutil"
       "log"
       "net"
       "net/http"
        "sync")

var mu sync.Mutex
var count int

func main(){

    if len(os.Args) != 2{

	fmt.Fprint(os.Stderr, "Usage: %s Host:Port ", os.Args[0])
	os.Exit(1)

    }

    service := os.Args[1]
    tcpAddr, err := net.ResolveTCPAddr("tcp", service)
    checkError(err)

    conn,err := net.DialTCP("tcp6", nil, tcpAddr)
    checkError(err)

    result, err := ioutil.ReadAll(conn)
    checkError(err)

    fmt.Println(string(result))
    os.Exit(0)

}


func checkError(err error){

    if err != nil{

	fmt.Fprintf(os.Stderr,"Error : %s", err.Error())
	os.Exit(1)

    }

}

func Ear(){

    http.HandleFunc("/",handler)
    http.HandleFunc("/count", counter)
    log.Fatal(http.ListenAndServe("Hal:9000",nil))//A space adventure

}

func handler(w http.ResponseWriter, r *http.Request){

    mu.Lock()
    count++
    mu.Unlock()

}

func counter(w http.ResponseWriter, r *http.Request){

    mu.Lock()//I need to look up mu and read the docs on Go Socket stuff. Crazy package work
    fmt.Fprintf(w, "The count is: %d\n", count)

}

