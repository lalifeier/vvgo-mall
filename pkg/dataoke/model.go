package dataoke

type CarouseListResp struct {
	RequestID string    `json:"requestId"`
	Time      int64     `json:"time"`
	Code      int       `json:"code"`
	Msg       string    `json:"msg"`
	Data      []Carouse `json:"data"`
}

type Carouse struct {
	TopicID    int    `json:"topicId"`
	TopicName  string `json:"topicName"`
	TopicImage string `json:"topicImage"`
	Link       string `json:"link"`
	SourceType int    `json:"sourceType"`
	ActivityID string `json:"activityId"`
}

type GetSuperCategoryResp struct {
	Cache     bool         `json:"cache"`
	Code      int          `json:"code"`
	Data      []Categories `json:"data"`
	Msg       string       `json:"msg"`
	RequestID string       `json:"requestId"`
	Time      int64        `json:"time"`
}

type Subcategories struct {
	Scpic    string `json:"scpic"`
	Subcid   int    `json:"subcid"`
	Subcname string `json:"subcname"`
}

type Categories struct {
	Cid           int             `json:"cid"`
	Cname         string          `json:"cname"`
	Cpic          string          `json:"cpic"`
	Subcategories []Subcategories `json:"subcategories"`
}
