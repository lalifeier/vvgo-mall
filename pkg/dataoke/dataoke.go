package dataoke

import (
	"fmt"
	"strings"
	"time"

	"github.com/lalifeier/vvgo-mall/pkg/crypto/md5"
	"github.com/lalifeier/vvgo-mall/pkg/random"
	"github.com/lalifeier/vvgo-mall/pkg/request"
	"github.com/lalifeier/vvgo-mall/pkg/util/convert"
)

var defaultClientOptions = ClientOptions{
	Version: "v2.0.0",
}

type ClientOption func(*ClientOptions)

type ClientOptions struct {
	Version string
}

func WithVersion(v string) ClientOption {
	return func(o *ClientOptions) {
		o.Version = v
	}
}

type Client interface {
	GenerateSign(params map[string]string) map[string]string
	Get(url string, params map[string]string) (string, error)

	CarouseList(params map[string]string) (string, error)
}

type client struct {
	appKey    string
	appSecret string
	version   string
}

func NewClient(appKey, appSecret string, opts ...ClientOption) Client {
	options := defaultClientOptions

	for _, o := range opts {
		o(&options)
	}

	return &client{
		appKey:    appKey,
		appSecret: appSecret,
		version:   options.Version,
	}
}

// appKey
// appsecret
// nonce：一个6位的随机数
// timer：毫秒级时间戳
func (c *client) GenerateSign(params map[string]string) map[string]string {
	if _, ok := params["appKey"]; !ok {
		params["appKey"] = c.appKey
	}
	if _, ok := params["appSecret"]; !ok {
		params["key"] = c.appSecret
	}
	if _, ok := params["version"]; !ok {
		params["version"] = c.version
	}

	params["nonce"] = random.GenerateNumber(6)
	params["timer"] = convert.Int64ToString(time.Now().Unix() * 1000)

	data := fmt.Sprintf("appKey=%s&timer=%s&nonce=%s&key=%s", params["appKey"], params["timer"], params["nonce"], params["key"])
	params["signRan"] = strings.ToUpper(md5.Encode(data))

	return params
}

func (c *client) Get(url string, params map[string]string) (string, error) {
	params = c.GenerateSign(params)
	return request.Get(url, params, nil)
}

// func (c *Client) PostJson(url string, data map[string]string) (string, error) {
// 	return request.PostJson(url, data, nil)
// }

// 轮播图
func (c *client) CarouseList(params map[string]string) (string, error) {
	return c.Get("https://openapi.dataoke.com/api/goods/topic/carouse-list", params)
}
