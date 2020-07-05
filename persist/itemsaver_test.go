package persist

import (
	"context"
	"encoding/json"
	"imooc.com/ccmouse/learngo/crawler_concurrent/engine"
	"testing"

	"github.com/olivere/elastic/v7"
	"imooc.com/ccmouse/learngo/crawler_concurrent/model"
)

func TestSave(t *testing.T) {

	expected := engine.Item{
		Url:  "http://www.7799520.com/user/595058.html",
		Type: "zhenai",
		Id:   "PvLVEnMBHeXOqtUrJH1P",
		Payload: model.Profile{
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
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	// Save expected item
	const index = "dating_test"
	err = Save(client, index, expected)
	if err != nil {
		panic(err)
	}

	// Fetch saved item
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	err = json.Unmarshal(resp.Source, &actual)
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	//
	//err = json.Unmarshal([]byte(resp.Source), &actual)
	//if err != nil {
	//	panic(err)
	//}

	// verifed result
	if actual != expected {
		t.Errorf("got %v != %v",
			actual, expected)
	}
}
