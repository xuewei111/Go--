package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

var (
	// 电话
	rePhone = `(1[3456789]\d)(\d{4})(\d{4})`
	// 邮箱
	reEmail = `[\w\.]+@\w+\.[a-z]{2,3}(\.[a-z]{2,3})?`
	// 超链接
	reLink = `<a[\s\S]+?href="(http[\s\S]+?)"`
	// 身份证
	reID = `^[1-9]\\d{7}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}$|^[1-9]\\d{5}[1-9]\\d{3}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}([0-9]|X)$`
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
	html := GetHtml("https://www.thepaper.cn/newsDetail_forward_8981455")

	r := regexp.MustCompile(reID)

	allString := r.FindAllStringSubmatch(html, -1)
	for _, x := range allString {
		fmt.Println(x[1])
	}
}
