package dataoke

import "github.com/lalifeier/vvgo-mall/pkg/request"

type Client struct {
	Options *Options
}

type OptionsFunc func(*Options)

type Options struct {
	appKey    string
	appSecret string
	version   string
}

func WithAppKey(v string) OptionsFunc {
	return func(o *Options) {
		o.appKey = v
	}
}

func WithAppSecret(v string) OptionsFunc {
	return func(o *Options) {
		o.appSecret = v
	}
}

// func NewClient(opts ...Options) *Server {
// 	options := Options{}

// 	for _, opt := range opts {
// 		opt.apply(&options)
// 	}

// 	return &Server{}
// }

// appKey
// appsecret
// nonce：一个6位的随机数
// timer：毫秒级时间戳
func (c *Client) GenerateSign(params map[string]string) {
	if _, ok := params["appKey"]; !ok {
		params["appKey"] = c.Options.appKey
	}
	if _, ok := params["appSecret"]; !ok {
		params["appSecret"] = c.Options.appSecret
	}
	if _, ok := params["version"]; !ok {
		params["version"] = c.Options.version
	}

	// md5 signRan
	// appKey=xxx&timer=xxx&nonce=xxx&key=xxx
}

func (c *Client) Get(url string, params map[string]string) (string, error) {
	return request.Get(url, params, nil)
}

func (c *Client) PostJson(url string, data map[string]string) (string, error) {
	return request.PostJson(url, data, nil)
}

// 轮播图
func (c *Client) CarouseList() (string, error) {
	return c.Get("https://openapi.dataoke.com/api/goods/topic/carouse-list", nil)
}
