package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	_ "io/ioutil"
	"log"
	"net/http"
)

//const url string = "http://www.zhenai.com/zhenghun"
func main() {
	response, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 包含头部信息
	//httputil.DumpResponse()
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", response.StatusCode)
		return
	}

	//all, err := ioutil.ReadAll(response.Body)
	//if err != nil  {
	//	panic(err)
	//}
	//fmt.Printf("%s\n", all)

	e := determineEncoding(response.Body)
	reader := transform.NewReader(response.Body, e.NewDecoder())
	if response.StatusCode != http.StatusOK {
		log.Println("Error response ", response.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", bodyBytes)
}

func  determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}