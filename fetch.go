// Fetch выводит ответ на запрос по заданному URL.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		pref1 := "http://"
		pref2 := "https://"
		if strings.HasPrefix(url, pref2) == false && strings.HasPrefix(url, pref1) == false {
			url = pref1 + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body) / deprecated
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
