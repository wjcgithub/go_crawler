package parser

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	contents, err := ioutil.ReadFile("/Users/wangjichao/www/go/go-ccmouse/crawler/zhenai/parser/reg_test_data.html")
	if err != nil {
		panic(err)
	}

	var ageRe = regexp.MustCompile(
		`<span class="age s1">(\d+)岁</span>`)

	ret := ageRe.FindSubmatch(contents)
	for _, item := range ret {
		fmt.Printf("%s \n", item)
	}
}
