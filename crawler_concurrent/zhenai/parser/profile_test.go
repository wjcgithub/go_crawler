package parser

import (
	engine2 "go_crawler/crawler_concurrent/engine"
	model2 "go_crawler/crawler_concurrent/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "http://www.7799520.com/user/595058.html","安静的雪")
	if len(result.Items) != 1 {
		t.Errorf("Items uld containe 1 element."+
			"element; but was %v", result.Items)

		actual := result.Items[0]

		expected := engine2.Item{
			Url:     "http://www.7799520.com/user/595058.html",
			Type:    "zhenai",
			Id:      "PvLVEnMBHeXOqtUrJH1P",
			Payload: model2.Profile{
				Age:        34,
				Height:     162,
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
			},
		}

		if actual != expected {
			t.Errorf("expected %v; but was %v",
				expected, actual)
		}
	}
}
