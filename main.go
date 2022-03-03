package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func HandleErr(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

func main() {
	resp, err := http.Get("https://www.gouhaowang.com/")

	HandleErr(err, `http.Get("https://www.gouhaowang.com/")`)

	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)
	fmt.Println(html)
}
