package parser

import (
	engine2 "go_crawler/crawler_concurrent/engine"
	model2 "go_crawler/crawler_concurrent/model"
	"regexp"
	"strconv"
	"strings"
)

var ageRe = regexp.MustCompile(
	`<span class="age s1">(\d+)岁</span>`)
var heightRe = regexp.MustCompile(
	`<span class="height">(\d+)cm</span>`)
var genderRe = regexp.MustCompile(`<span>([^<]+)</span>`)
var incomeRe = regexp.MustCompile(
	`<li>收入：<span>([^<]+)</span></li>`)
var xinzuoRe = regexp.MustCompile(
	`<li>星座：<span>([^<]+)</span></li>`)
var marriageRe = regexp.MustCompile(
	`<span class="marrystatus">([^<]+)</span>`)
var educationRe = regexp.MustCompile(
	`<span class="education">([^<]+)</span>`)
var occupationRe = regexp.MustCompile(
	`<li>职业：<span>([^<]+)</span></li>`)
var hokouRe = regexp.MustCompile(
	`<li>现居：<span>([^<]+)</span></li>`)
var imgRe = regexp.MustCompile(`<li class="" data-uid="[0-9]+"><img src="(http://img.7799520.com/.+\.png)" alt=""></li>`)

var idUrlRe = regexp.MustCompile(`http://www.7799520.com/user/([\d]+)\.html`)

func ParseProfile(contents []byte, url string, name string) engine2.ParseResult {
	profile := model2.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(
		extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(
		extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	profile.Income = extractString(
		contents, incomeRe)
	profile.Gender = extractString(
		contents, genderRe)
	profile.Car = extractString(
		contents, imgRe)
	profile.Education = extractString(
		contents, educationRe)
	profile.Hokou = extractString(
		contents, hokouRe)
	//profile.House = extractString(
	//	contents, houseRe)
	profile.Marriage = extractString(
		contents, marriageRe)
	profile.Occupation = extractString(
		contents, occupationRe)
	profile.Xinzuo = extractString(
		contents, xinzuoRe)

	result := engine2.ParseResult{
		Items: []engine2.Item{{
			Url:     url,
			Type:    "zhenai",
			Id:      extractString([]byte(url), idUrlRe),
			Payload: profile,
		}},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {

	}
	if match != nil {
		return strings.TrimSpace(string(match[1]))
	} else {
		return ""
	}
}

func ProfileParser(name string) engine2.ParserFunc {
	return func(c []byte, url string) engine2.ParseResult {
		return ParseProfile(c, url, name)
	}
}