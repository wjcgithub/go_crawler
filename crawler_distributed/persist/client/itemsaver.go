package client

import (
	"fmt"
	"go_crawler/crawler_concurrent/engine"
	config2 "go_crawler/crawler_distributed/config"
	rpcsupport2 "go_crawler/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport2.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	profileCount := 0
	go func() {
		for {
			item := <-out
			log.Printf("Item saver got profile: #%d %v", profileCount, item)
			profileCount++
			result := ""
			err = client.Call(config2.ItemSaverRpc, item, &result)
			if err != nil {
				fmt.Printf("Item Saver： error "+
					" saveing item %v: %v", item, err)
			}
		}
	}()

	return out, nil
}
