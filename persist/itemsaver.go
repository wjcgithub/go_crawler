package persist

import (
	"imooc.com/ccmouse/learngo/crawler_concurrent/model"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	profileCount := 0
	go func() {
		for {
			item := <-out
			if _, ok := item.(model.Profile); ok {
				log.Printf("Item saver got profile: #%d %v", profileCount, item)
				profileCount++
			}
		}
	}()

	return out
}
