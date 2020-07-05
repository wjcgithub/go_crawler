package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun") http://www.7799520.com/jiaou
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents, "")

	const resultSize = 389
	expectedUrls := []string{
		"http://www.7799520.com/jiaou/anhui", "http://www.7799520.com/jiaou/aomen", "http://www.7799520.com/jiaou/anqing",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d"+
			"reuests; but had %d",
			resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but "+
				"was %s",
				i, url, result.Requests[i].Url)
		}
	}
}
