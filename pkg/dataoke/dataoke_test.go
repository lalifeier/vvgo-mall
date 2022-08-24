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

func Test(t *testing.T) {
	c := NewClient(appKey, appSecret, WithVersion(version))

	var (
		params = make(map[string]string)
	)

	// params[""] =

	resp, _ := c.CarouseList(params)

	var data CarouseList
	if err := json.Unmarshal([]byte(resp), &data); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data.Code)
	fmt.Println(data.Data[0].TopicID)
}
