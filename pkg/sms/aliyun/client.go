package sms

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

var (
	scheme string = "https"
)

type Client struct {
	client       *dysmsapi.Client
	SignName     string
	TemplateCode string
}

func NewClient(regionId, accessKeyId, accessKeySecret, signName, templateCode string) (*Client, error) {
	client, err := dysmsapi.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}

	return &Client{
		SignName:     signName,
		TemplateCode: templateCode,
		client:       client,
	}, nil
}

func Code() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(899999) + 100000
	res := strconv.Itoa(code)
	return res
}

func (c *Client) SendCode(mobile, code string) (err error) {
	request := dysmsapi.CreateSendSmsRequest()

	request.Scheme = scheme
	request.PhoneNumbers = mobile
	request.SignName = c.SignName
	request.TemplateCode = c.TemplateCode
	request.TemplateParam = fmt.Sprintf("{code:%s}", code)

	response, err := c.client.SendSms(request)
	if err != nil {
		return errors.New("阿里云短信发送失败" + err.Error())
	}
	if response.Code != "OK" {
		if response.Code == "isv.BUSINESS_LIMIT_CONTROL" {
			return errors.New("获取短信过于频繁，请稍后再试")
		}
		return errors.New("阿里云短信发送失败" + response.Message)
	}
	return nil
}
