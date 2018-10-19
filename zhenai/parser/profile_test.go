package parser

import (
	"testing"
	"io/ioutil"
	"fmt"
)

func TestParseProfile(t *testing.T) {
	contens, err := ioutil.ReadFile("testprofile.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contens,"醉梦人")
	fmt.Println(result)

}
