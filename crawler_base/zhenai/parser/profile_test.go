package parser

import (
	model2 "go_crawler/crawler_base/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "安静的雪")
	if len(result.Items) != 1 {
		t.Errorf("Items uld containe 1 element."+
			"element; but was %v", result.Items)

		profile := result.Items[0].(model2.Profile)

		expected := model2.Profile{
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			House:      "已购房",
			Hokou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		}

		if profile != expected {
			t.Errorf("expected %v; but was %v",
				expected, profile)
		}
	}
}
