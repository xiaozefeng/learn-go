package persist

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"imooc.com/learn-go/crawler/model"
	"testing"
	"encoding/json"
	"imooc.com/learn-go/crawler/engine"
)

func TestSave(t *testing.T) {

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/105992496",
		Id:   "105992496",
		Type: "zhenai",
		Payload: model.Profile{
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			Hourse:     "已购房",
			Hokou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}
	err := save(expected)
	if err != nil {
		panic(err)
	}
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index("dating_profile").Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", *resp.Source)

	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if expected != actual {
		t.Errorf("Got %+v; expected %+v", actual, expected)
	}

}
