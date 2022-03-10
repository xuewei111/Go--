package main

import (
	"io/ioutil"
	"net/http"
)

var (
	reImg = `<img[\s\S]+?src="(http[\s\S]+?)"`
)

func GetHtml(url string) string {
	r, _ := http.Get(url)
	defer r.Body.Close()

	b, _ := ioutil.ReadAll(r.Body)
	return string(b)
}


