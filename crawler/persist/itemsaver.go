package persist

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"fmt"
	"imooc.com/learn-go/crawler/engine"
	"github.com/pkg/errors"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: Got item #%d, %+v", itemCount, item)
			itemCount++
			err := save(item)
			if err != nil {
				fmt.Printf("Item Saver: error saving item: %+v: %v ", item,
					err)
			}
		}
	}()
	return out
}

func save(item engine.Item) error {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}
	if item.Type == "" {
		return errors.New("Must Supply Type")
	}

	indexService := client.Index().
		Index("dating_profile").
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
