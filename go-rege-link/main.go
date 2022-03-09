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
	reEmail = `[\w\.]+@\w+\.[a-z]{2,3}(\.[a-z]{2,3})?`
	reLink  = `<a[\s\S]+?href="(http[\s\S]+?)"`
)

func HandleErr(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

func GetHtml(url string) string {
	resp, err := http.Get(url)

	HandleErr(err, `http.Get`)

	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)
	return html
}

func main() {
	html := GetHtml("http://www.abcdao.com/")

	r := regexp.MustCompile(reLink)

	allString := r.FindAllStringSubmatch(html, -1)
	for _, x := range allString {
		fmt.Println(x[1])
	}
}
