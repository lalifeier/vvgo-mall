package tencent

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

// https://aws.github.io/aws-sdk-go-v2/docs/

type (
	Config struct {
		AccessKeyId     string // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		AccessKeySecret string // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		SessionToken    string
		Region          string
		Bucket          string
	}

	Client struct {
		Client *cos.Client
		Config *Config
	}
)

func getBucketURL(bucket, region string) string {
	return fmt.Sprintf("https://%s.cos.%s.myqcloud.com", bucket, region)
}

func getServiceURL(region string) string {
	return fmt.Sprintf("https://cos.%s.myqcloud.com", region)
}

func New(c *Config) (*Client, error) {
	bucketURL := getBucketURL(c.Bucket, c.Region)
	u, _ := url.Parse(bucketURL)

	serviceURL := getServiceURL(c.Region)
	su, _ := url.Parse(serviceURL)
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:     c.AccessKeyId,
			SecretKey:    c.AccessKeySecret,
			SessionToken: c.SessionToken,
		},
	})

	return &Client{
		Client: client,
		Config: c,
	}, nil
}

func (c *Client) PutObject(objectKey string, filepath string) (err error) {
	_, _, err = c.Client.Object.Upload(context.Background(), objectKey, filepath, nil)
	return err
}

func (c *Client) DeleteObject(objectKey string) error {
	_, err := c.Client.Object.Delete(context.Background(), objectKey)
	return err
}

func (c *Client) DeleteObjects(objectKeys []string) error {
	obs := []cos.Object{}
	for _, objectKey := range objectKeys {
		obs = append(obs, cos.Object{Key: objectKey})
	}
	opt := &cos.ObjectDeleteMultiOptions{
		Objects: obs,
		// 布尔值，这个值决定了是否启动 Quiet 模式
		// 值为 true 启动 Quiet 模式，值为 false 则启动 Verbose 模式，默认值为 false
		// Quiet: true,
	}
	_, _, err := c.Client.Object.DeleteMulti(context.Background(), opt)
	return err
}

func (c *Client) GetObject(objectKey string, expires int64) (io.ReadCloser, error) {
	resp, err := c.Client.Object.Get(context.Background(), objectKey, nil)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (c *Client) GetObjectURL(objectKey string, expires int64) (string, error) {
	url := c.Client.Object.GetObjectURL(objectKey)
	return url.Path, nil
}

func (c *Client) DownloadObject(objectKey string, filepath string) (io.ReadCloser, error) {
	opt := &cos.MultiDownloadOptions{
		ThreadPoolSize: 5,
	}

	result, err := c.Client.Object.Download(context.Background(), objectKey, filepath, opt)
	if err != nil {
		return nil, err
	}
	return result.Body, nil
}

func (c *Client) ListObjects(objectKey string) error {
	var marker string
	opt := &cos.BucketGetOptions{
		Prefix:    "folder/", // prefix表示要查询的文件夹
		Delimiter: "/",       // deliter表示分隔符, 设置为/表示列出当前目录下的object, 设置为空表示列出所有的object
		MaxKeys:   1000,      // 设置最大遍历出多少个对象, 一次listobject最大支持1000
	}
	isTruncated := true
	for isTruncated {
		opt.Marker = marker
		v, _, err := c.Client.Bucket.Get(context.Background(), opt)
		if err != nil {
			fmt.Println(err)
			break
		}
		for _, content := range v.Contents {
			fmt.Printf("Object: %v\n", content.Key)
		}
		// common prefix表示表示被delimiter截断的路径, 如delimter设置为/, common prefix则表示所有子目录的路径
		for _, commonPrefix := range v.CommonPrefixes {
			fmt.Printf("CommonPrefixes: %v\n", commonPrefix)
		}
		isTruncated = v.IsTruncated // 是否还有数据
		marker = v.NextMarker       // 设置下次请求的起始 key
	}
	return nil
}

func (c *Client) IsObjectExist(objectKey string) (bool, error) {
	return c.Client.Object.IsExist(context.Background(), objectKey)
}
