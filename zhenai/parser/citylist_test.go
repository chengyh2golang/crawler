package parser

import (
	"testing"
	"io/ioutil"
)

func TestParserCityList(t *testing.T) {
	//contents, err := fetcher.Fetcher("http://m.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("testdata.html")
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n",contents)
	const  resultSize  =  481
	expectedUrl := []string {
		"http://www.zhenai.com/zhenghun/beijing",
		"http://www.zhenai.com/zhenghun/shanghai",
		"http://www.zhenai.com/zhenghun/tianjin",
	}

	expectedItem := []string {
		"北京",
		"上海",
		"天津",
	}
	result := ParserCityList(contents)
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d but had %d",resultSize,len(result.Requests))
	}

	for i,url := range expectedUrl {
		if result.Requests[i].Url != url {
			t.Errorf("expected #%d url:%s, got url: %s",i,url,result.Requests[i].Url)
		}
	}


	if len(result.Items) != resultSize {
		t.Errorf("result should have %d but had %d",resultSize,len(result.Items))
	}

	for i,item := range expectedItem {
		if result.Items[i].(string) != item {
			t.Errorf("expected #%d item: %s, got item: %s",i,item,result.Items[i])
		}
	}

}
