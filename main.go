package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

var (
	rePhone = `(1[3456789]\d)(\d{4})(\d{4})`
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
	// fmt.Println(html)

	re := regexp.MustCompile(rePhone)
	// allString := re.FindAllString(html, -1)
	allString := re.FindAllStringSubmatch(html, -1)
	// fmt.Println(allString)
	for _, x := range allString {
		fmt.Println(x)
	}

}
