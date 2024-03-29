//this is the second server
package main

import("fmt"
       "log"
       "net/http"
       "sync")

var mu sync.Mutex//mu is now this
var count int

func main(){

    http.HandleFunc("/",handler)
    http.HandleFunc("\count", counter)
    log.Fatal(http.ListenAndServe("localhost:8000",nil))

}

//handler echos the path of the resqusted URL
func handler(w http.ResponseWriter, r *http.Request){

    mu.Lock()
    count++
    mu.Unlock()
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

}
//counters echos the number of calls so far
func counter(w http.ResponseWriter, r. *http.Request){

    mu.Lock()
    fmt.Fprintf(w, "Count %d\n", count)

}

