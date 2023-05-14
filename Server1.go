//Now were gettin somewhere. BOOK WORK. Time to write out the first server
package main

import ("fmt"
        "log"
        "net/http")

func main(){

    http.HandleFunc("/", handler)//each request is gonna call the handlers
    log.Fatal(http.ListenAndServe("localhost:8000",nil))

}
//handler echoes the path component of the request url
func handler(w http.ResponseWriter, r *http.Request){

    fmt.Fprintf(w,"Url.Path = %q\n", r.URL.Path)

}
