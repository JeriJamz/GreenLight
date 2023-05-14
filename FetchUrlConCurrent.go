//Go is starting to get easier. This is gonna work on time and network programmin
//still book work
//I think this work but the "Fetch" func. Wont accepct any names
package main

import("fmt"
       "io"
       "io/ioutil"
       "net/http"
       "os"
       "time")

func main(){

    start := time.Now()//wen this is called. It started a stopwatch
    ch := make(chan string)//this is a string. Chan means channel
    for _, url := range os.Args[1:]{

	go fetch(url, ch)//a goRoutine. I think it grabs the url, and the channel(Port #?)

    }

    for range os.Args[1:]{

	fmt.Println(<-ch)//prints from the channel

    }

    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())//prints the time elapsed in seconds

    func fetch(url string, ch chan<- string){//makes the URL a string and grabs the channel and converts it into a string

	start := time.Now()//This is the Fetch Func stopwatch
	resp, err := http.Get(url)
	if err != nil{

	    ch <- fmt.Sprintf("While reading %s: %v", url, err)
	    return

	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()//Close the Convo
	if err != nil{

	    ch <- fmt.Sprintf("while reading %s: %v", url,err)
	    return

	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", seconds, nbytes, url)

    }

}
