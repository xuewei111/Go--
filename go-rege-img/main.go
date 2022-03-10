package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func GetPageImgUrls(url string) []string {
	html := GetHtml(url)

	// fmt.Println(html)
	re := regexp.MustCompile(reImg)
	rets := re.FindAllStringSubmatch(html, -1)
	fmt.Println("捕获到图片张数:", len(rets))

	imgUrls := make([]string, 0)
	for _, ret := range rets {
		imgUrls = append(imgUrls, ret[1])
	}

	return imgUrls
}

func DownLoadImg(url string) {
	r, _ := http.Get(url)
	defer r.Body.Close()

	imgBytes, _ := ioutil.ReadAll(r.Body)
	flieName := `D:\project\Go-アイテムコレクション\go-rege-img\images\` + strconv.Itoa(int(time.Now().UnixNano())) + ".jpg"

	err := ioutil.WriteFile(flieName, imgBytes, 0644)
	if err == nil {
		fmt.Println("下载成功")
	} else {
		fmt.Println("下载失败")
	}
}

func DownLoadImgAsync(url string) {

}

func main() {
	imgurls := GetPageImgUrls("https://www.163.com/")
	for _, imgUrl := range imgurls {
		fmt.Println(imgUrl)
		DownLoadImg(imgUrl)
	}
}
