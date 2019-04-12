package fetcher

import (
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	/*
	resp, err := http.Get(url)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	*/
	client := &http.Client{}

	// 提交请求
	request, err := http.NewRequest("GET", url, nil)

	// 增加header选项
	//  request.Header.Add("Accept", "text/html")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0;Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	//	reqest.Header.Add("Accept-Encoding", "gzip, deflate")
	if err != nil {
		panic(err)
	}

	// 处理返回结果
	resp, _ := client.Do(request)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	e := determineEncoding(resp.Body)
	// 转码
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}
