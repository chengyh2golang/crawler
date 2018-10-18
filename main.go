package main

import (
	"awesomeProject/crawler2/engine"
	"awesomeProject/crawler2/zhenai/parser"
)

func main() {

	//"http://m.zhenai.com/zhenghun"

	engine.Run(engine.Request{
		Url:"http://m.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})

}
