package persist

import (
	"context"
	"errors"
	"fmt"
	engine2 "go_crawler/crawler_concurrent/engine"
	"log"

	"github.com/olivere/elastic/v7"
)

func ItemSaver(index string) (chan engine2.Item, error) {
	client, err := elastic.NewClient(
		//Must turn off sniff in docker
		elastic.SetSniff(false),
	)

	if err != nil {
		return nil, err
	}

	out := make(chan engine2.Item)
	profileCount := 0
	go func() {
		for {
			item := <-out
			log.Printf("Item saver got profile: #%d %v", profileCount, item)
			profileCount++
			err := Save(client, index, item)
			if err != nil {
				fmt.Printf("Item Saver： error "+
					" saveing item %v: %v", item, err)
			}
		}
	}()

	return out, nil
}

func Save(client *elastic.Client, index string, item engine2.Item) error {
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
