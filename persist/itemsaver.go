package persist

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
	"imooc.com/ccmouse/learngo/crawler_concurrent/engine"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	profileCount := 0
	go func() {
		for {
			item := <-out
			log.Printf("Item saver got profile: #%d %v", profileCount, item)
			profileCount++
			err := save(item)
			if err != nil {
				fmt.Printf("Item Saver： error "+
					" saveing item %v: %v", item, err)
			}
		}
	}()

	return out
}

func save(item engine.Item) error {
	client, err := elastic.NewClient(
		//Must turn off sniff in docker
		elastic.SetSniff(false),
	)

	if err != nil {
		return err
	}

	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().
		Index("dating_profile_db_1").
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.
		Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
