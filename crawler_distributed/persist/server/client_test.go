package main

import (
	"fmt"
	"go_crawler/crawler_concurrent/engine"
	"go_crawler/crawler_concurrent/model"
	config2 "go_crawler/crawler_distributed/config"
	rpcsupport2 "go_crawler/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	// start ItemSaverServer
	const host = ":1234"
	go serveRpc(host, "test1")
	time.Sleep(time.Second)
	// start ItemSaverClient
	client, err := rpcsupport2.NewClient(host)
	if err != nil {
		panic(err)
	}
	// call save
	item := engine.Item{
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

	result := ""
	err = client.Call(config2.ItemSaverRpc, item, result)
	fmt.Printf("%v", result)
	//if err != nil || result != "ok" {
	//	t.Errorf("result: %+v; err: %+v",
	//		result, err)
	//}
}
