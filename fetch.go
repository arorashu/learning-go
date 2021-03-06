// 1.5 Fetching a URL
// fetches the contents of each URL and prints them uninterpreted
// inspired by curl

package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {

    for _, url := range os.Args[1:] {
        start := time.Now()
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        b, err := ioutil.ReadAll(resp.Body)
        fmt.Printf("Time to load url: %s: %v seconds\n", url, time.Since(start).Seconds())
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)
        }
        fmt.Printf("%s", b)
    }

}
