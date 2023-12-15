package util

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestDownloadAllJSON(t *testing.T) {
	configFilePath := "C:\\Users\\daishaokun1\\Desktop\\全国行政单位.txt"
	config := gfile.GetContents(configFilePath)

	arr := gstr.SplitAndTrim(config, "\n")
	for _, v := range arr {
		// 获取前6位
		code := v[:6]
		url := getUrl(code)
		if url == "" {
			continue
		}
		go downloadJSON(url, code)
	}
	//阻塞主线程
	select {}
}

func downloadJSON(url string, code string) {
	fmt.Println("开始下载：" + url)
	folder := "D:\\file\\geojson\\"
	fileName := code + ".json"
	filePath := folder + fileName
	if gfile.Exists(filePath) {
		return
	}

	// Send HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	// Write the body to file
	_, err = io.Copy(out, response.Body)
	if err != nil {
		_ = gfile.Remove(filePath)
		panic(err)
	}

	fmt.Println(code, "successfully")
}

func getUrl(code string) string {
	code = gstr.TrimAll(code)
	if len(code) == 0 {
		return ""
	}

	/*
		var subfix = "";
		        //如果code是市级或省级（最后两位是00），则subfix改成"_full"
		        //截取最后两位
		        var lastTwo = code.substring(code.length - 2);
		        if (lastTwo === "00") {
		          subfix = "_full";
		        }
	*/

	subfix := ""
	if code[len(code)-2:] == "00" {
		subfix = "_full"
	}
	return "https://geo.datav.aliyun.com/areas_v3/bound/" + code + subfix + ".json"
}
