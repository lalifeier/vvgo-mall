package dataoke

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	appKey    = ""
	appSecret = ""
	version   = "v2.0.0"
)

type CarouseList struct {
	RequestID string `json:"requestId"`
	Time      int64  `json:"time"`
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Data      []Data `json:"data"`
}

type Data struct {
	TopicID    int    `json:"topicId"`
	TopicName  string `json:"topicName"`
	TopicImage string `json:"topicImage"`
	Link       string `json:"link"`
	SourceType int    `json:"sourceType"`
	ActivityID string `json:"activityId"`
}

func TestCarouseList(t *testing.T) {
	c := NewClient(appKey, appSecret, WithVersion(version))

	var (
		params = make(map[string]string)
	)

	data, _ := c.CarouseList(params)
	carouseList, _ := json.Marshal(data)

	fmt.Println(string(carouseList))
}

func TestGetSuperCategory(t *testing.T) {
	c := NewClient(appKey, appSecret, WithVersion(version))

	var (
		params = make(map[string]string)
	)

	data, _ := c.GetSuperCategory(params)
	_data, _ := json.Marshal(data)

	fmt.Println(string(_data))
}
