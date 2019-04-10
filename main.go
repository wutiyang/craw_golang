package main

import (
	"log"
	"net/http"
)

const url string = "http://www.zhenai.com/zhenghun"

func main() {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	if response.StatusCode != http.StatusOK {
		panic("http respose error")
	}

	body := response.Body
	log.Printf("%s", body)

	//HTTP.get(url)
	//re
	//css选择器
	//xpath
	//
	//转码
	//return contents
	//
	//// 解析器解析
	//return parseResult{
	//request
	//parseFunc
	//item
	//}
	//
	//// 循环解析结果
	//for result in range parseResult {
	//// requests ，不断推送给fatcher
	//urls{
	//url
	//parseFunc
	//}
	//
	//// item
	//log.printf()
	//}
	//
	//// main文件：
	//url-》fatcher=〉解析-》url，开始第一步
	//request       reqeust     item
	//engine -〉 schedule -> worker  -> engine
	//各自功能：
	//engine：启动？
	//schedule：接受request，发送request？
	//worker：fatch，parse
	//如何衔接：？？？
}
