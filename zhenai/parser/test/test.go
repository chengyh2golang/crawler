package main

import (
	"regexp"
	"fmt"
	"strings"
)

const test  = `<a href="http://m.zhenai.com/u/1571699528#seo" class="left-item" data-v-a269c2a6><div class="right-item" data-v-a269c2a6><div class="f-nickname" data-v-a269c2a6>
小云
<span class="u-gender0" data-v-a269c2a6></span> <img src="//i.zhenai.com/m/zhenghun/zhDetail/static/images/icon-women.ba9d641.png" alt data-v-a269c2a6></div> <div class="t-info" data-v-a269c2a6>36岁 | 北京大兴区 | 165cm</div> <div class="c-tag" data-v-a269c2a6><div class="tag" data-v-7a8234e4 data-v-a269c2a6>
离异
</div>`
const  cityRe = `<a href="http://m.(zhenai.com/u/[0-9]+)[^d]*" class="left-item"[^>]*><div class="right-item"[^>]*><div class="f-nickname"[^>]*>([^<]+)<span class=`
func main() {

	re := regexp.MustCompile(cityRe)
	matches := re.FindAllStringSubmatch(test,-1)
	for _, match := range matches {
		fmt.Println(strings.TrimSpace(match[2]))
		fmt.Println(match[1])
	}


}
