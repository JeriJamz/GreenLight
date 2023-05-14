//This is some mo bookwork on fetchin a URL. Lil more Network Programmin
package main

import("fmt"
       "io/ioutil"
       "net/http"
       "os")

func main(){

    for _, url := range os.Args[1:]{//Range means URL is the Limit and can b extended I think

	resp, err := http.Get(url)
	//if these not an URL print the error message
	if err != nil {

	    fmt.Fprintf(os.Stderr," Fetch: %v\n", err)
	    os.Exit(1)

	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err !=nil{

	    fmt.Fprintf(os.Stderr," Fetch: Reading %s: %v\n",url,err)
	    os.Exit(1)

	}

	fmt.Printf("%s", b)

    }

}


