package parser

import (
	"imooc.com/learn-go/crawler/model"
	"io/ioutil"
	"testing"
	"fmt"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "")

	if len(result.Items) != 1 {
		t.Errorf("Items shoud contain 1 element , but was %v", result.Items)
	}

	profile:= result.Items[0].(model.Profile)
	expected := model.Profile{
		Age:        34,
		Height:     162,
		Weight:     57,
		Income:     "",
		Gender:     "",
		Name:       "",
		Xinzuo:     "",
		Occupation: "",
		Marriage:   "",
		Hourse:     "",
		Car:        "",
		Education:  "",
	}

	if profile != expected {
		fmt.Println(profile)
		fmt.Println(expected)
	}

}
