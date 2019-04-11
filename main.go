package main

import (
	"fmt"
	"io/ioutil"
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

	all, err := ioutil.ReadAll(response.Body)
	if err != nil  {
		panic(err)
	}
	fmt.Printf("%s\n", all)

}
