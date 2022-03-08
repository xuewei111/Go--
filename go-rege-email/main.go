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
)

func HandleErr(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

func main() {
	html := GetHtml("https://www.xyforex.com.cn/hot/431911")
	// fmt.Println(html)

	// re := regexp.MustCompile(rePhone)
	re := regexp.MustCompile(reEmail)
	// allString := re.FindAllString(html, -1)
	allString := re.FindAllStringSubmatch(html, -1)
	// fmt.Println(allString)
	for _, x := range allString {
		fmt.Println(x[0])
	}

}

func GetHtml(url string) string {
	resp, err := http.Get(url)

	HandleErr(err, `http.Get`)

	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)
	return html
}
