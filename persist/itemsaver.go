package persist

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
	"imooc.com/ccmouse/learngo/crawler_concurrent/engine"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		//Must turn off sniff in docker
		elastic.SetSniff(false),
	)

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
			err := save(client, index, item)
			if err != nil {
				fmt.Printf("Item Saver： error "+
					" saveing item %v: %v", item, err)
			}
		}
	}()

	return out, nil
}

func save(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.
		Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
