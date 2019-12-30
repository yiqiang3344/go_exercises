// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	//"io"
	"strings"
	//"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url,"https://") {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//out, err := io.Copy(os.Stdout, resp.Body);
		//b, err := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s", resp.Status)
		resp.Body.Close()
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		//	os.Exit(1)
		//}
		//fmt.Printf("%s", out)
	}
}