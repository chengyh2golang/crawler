package parser

import (
	"regexp"
	"crawler/engine"
)
const  cityListRe = `<a href="http://m.(zhenai.com/zhenghun/[a-z0-9]+)"[^>]+>([^<]+)</a>`
//const  cityListRe  = `<a href="//m.(zhenai.com/zhenghun/[a-z0-9]+)">([^<]+)</a>`


func ParserCityList(contens []byte)  engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	//matches := re.FindAll(contens,-1)

	//-1，找到所有的匹配个数；FindAllSubmatch的返回值[][][]byte 相当于[][]string
	matches := re.FindAllSubmatch(contens, -1)

	result := engine.ParserResult{}
	for _,m := range matches {

		result.Items = append(result.Items,string(m[2]))
		//fmt.Println("http://www."+string(m[1]))
		result.Requests = append(result.Requests, engine.Request{Url:"http://m." + string(m[1]),ParserFunc:ParserCity})

	}
	//fmt.Println(len(matches))
	return result
}
