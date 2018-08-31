package persist

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"fmt"
	"imooc.com/learn-go/crawler/engine"
	"github.com/pkg/errors"
)

func ItemSaver(index string) (item chan engine.Item, err error) {
	out := make(chan engine.Item)
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: Got item #%d, %+v", itemCount, item)
			itemCount++
			err := save(index, client, item)
			if err != nil {
				fmt.Printf("Item Saver: error saving item: %+v: %v ", item,
					err)
			}
		}
	}()
	return out, nil
}

func save(index string, client *elastic.Client, item engine.Item) error {
	if item.Type == "" {
		return errors.New("Must Supply Type")
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
