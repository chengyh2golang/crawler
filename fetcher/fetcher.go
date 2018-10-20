package fetcher

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"

	"golang.org/x/text/encoding/unicode"
	"log"
)

func DeterminEncoding(r *bufio.Reader) encoding.Encoding{

	//用peek的原因，是因为DetermineEncoding要读前面的1024字节，peek相当于偷看前面的1024个字节
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetcher error :%v",err)
		return unicode.UTF8  //如果peek不出来编码，那么打印一个log信息并返回一个默认值utf-8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return  e
}

func Fetcher(url string) ([]byte,error){

	//resp, err := http.Get(url)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil,err
	}
	request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.137 Safari/537.36 LBBROWSER")
	resp, err :=http.DefaultClient.Do(request)
	if  err!=nil{
		return nil,err
	}
	if resp.StatusCode != http.StatusOK {
		//fmt.Println("err status code.",resp.StatusCode)
		//使用errors.New()或者fmt.Errorf()来生成一个error
		return nil,
		fmt.Errorf("wrong status code：%d",resp.StatusCode)
	}
	//body, err := httputil.DumpResponse(resp, true)
	//if err != nil {
	//	panic(err)
	//}

	bodyReader := bufio.NewReader(resp.Body)
	//调用deteminEncoding，通过看前1024个字节，检测出网页的编码格式并返回
	e:= DeterminEncoding(bodyReader)
	//fmt.Println(e)

	//用自动检测到的编码格式去做跟utf-8的转换
	utf8reader := transform.NewReader(bodyReader, e.NewDecoder())

	//utf8reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())

	//读转换成utf8格式后的resp.body
	return ioutil.ReadAll(utf8reader)

}
