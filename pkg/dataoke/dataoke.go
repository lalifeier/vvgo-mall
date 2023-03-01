package dataoke

import (
	"fmt"
	"strings"
	"time"

	"github.com/lalifeier/vvgo-mall/pkg/util/convert"
	"github.com/lalifeier/vvgo-mall/pkg/util/crypto/md5"
	"github.com/lalifeier/vvgo-mall/pkg/util/random"
	"github.com/lalifeier/vvgo-mall/pkg/util/request"
)

var defaultClientOptions = clientOptions{
	Version: "v2.0.0",
	BaseUrl: "https://openapi.dataoke.com/api",
}

type ClientOption func(*clientOptions)

type clientOptions struct {
	Version string
	BaseUrl string
}

func WithVersion(v string) ClientOption {
	return func(o *clientOptions) {
		o.Version = v
	}
}

// type Client interface {
// 	GenerateSign(params map[string]string) map[string]string
// 	Get(url string, params map[string]string) (string, error)

// 	CarouseList(params map[string]string) (string, error)
// }

type Client struct {
	appKey    string
	appSecret string
	version   string
	baseUrl   string
}

func NewClient(appKey, appSecret string, opts ...ClientOption) *Client {
	options := defaultClientOptions

	for _, o := range opts {
		o(&options)
	}

	return &Client{
		appKey:    appKey,
		appSecret: appSecret,
		version:   options.Version,
		baseUrl:   options.BaseUrl,
	}
}

// appKey
// appsecret
// nonce：一个6位的随机数
// timer：毫秒级时间戳
func (c *Client) GenerateSign(params map[string]string) map[string]string {
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

func (c *Client) Get(url string, params map[string]string) (string, error) {
	params = c.GenerateSign(params)
	url = c.baseUrl + url
	return request.Get(url, params, nil)
}
