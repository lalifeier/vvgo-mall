package aliyun

import (
	"io"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// https://help.aliyun.com/product/31815.html

type (
	Config struct {
		AccessKeyId     string
		AccessKeySecret string
		SessionToken    string
		Region          string
		Bucket          string
		Endpoint        string
		UseCname        bool
		ACL             oss.ACLType
		ClientOptions   []oss.ClientOption
	}

	Client struct {
		Bucket *oss.Bucket
		Config *Config
	}
)

func New(config *Config) (*Client, error) {
	if config.Endpoint == "" {
		config.Endpoint = "http://oss-cn-hangzhou.aliyuncs.com"
	}

	if config.ACL == "" {
		config.ACL = oss.ACLPublicRead
	}

	if config.UseCname {
		config.ClientOptions = append(config.ClientOptions, oss.UseCname(config.UseCname))
	}

	client, err := oss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret, config.ClientOptions...)
	if err != nil {
		return nil, err
	}

	// err = client.CreateBucket(bucketName)
	// if err != nil {
	// 	return nil, err
	// }

	bucket, err := client.Bucket(config.Bucket)
	if err != nil {
		return nil, err
	}

	return &Client{
		Bucket: bucket,
		Config: config,
	}, err
}

func (c Client) GetRelativePath(objectName string) string {
	return strings.TrimPrefix(objectName, "/")
}

func (c *Client) PutObject(objectName string, reader io.Reader) (err error) {
	return c.Bucket.PutObject(objectName, reader)
}

func (c *Client) DeleteObject(objectKey string) error {
	return c.Bucket.DeleteObject(objectKey)
}

func (c *Client) DeleteObjects(objectKeys []string) error {
	_, err := c.Bucket.DeleteObjects(objectKeys)
	return err
}

func (c *Client) GetObject(objectKey string) (io.ReadCloser, error) {
	return c.Bucket.GetObject(objectKey)
}

func (c *Client) GetObjectURL(objectKey string, expires int64) (url string, err error) {
	url, err = c.Bucket.SignURL(objectKey, oss.HTTPGet, expires)
	return
}

func (c *Client) DownloadObject(objectKey, filePath string) error {
	return c.Bucket.GetObjectToFile(objectKey, filePath)
}

func (c *Client) ListObjects(objectName string) (err error) {
	results, err := c.Bucket.ListObjects()
	if err != nil {
		return err
	}

	// var objects []*oss.Object
	for _, result := range results.Objects {
		print(result.Key)
	}

	return nil
}

func (c *Client) IsObjectExist(objectKey string) (isExist bool, err error) {
	return c.Bucket.IsObjectExist(objectKey)
}

func (c Client) GetEndpoint() string {
	if c.Config.Endpoint != "" {
		if strings.HasSuffix(c.Config.Endpoint, "aliyuncs.com") {
			return c.Config.BucketName + "." + c.Config.Endpoint
		}
		return c.Config.Endpoint
	}

	endpoint := c.Bucket.Client.Config.Endpoint
	for _, prefix := range []string{"https://", "http://"} {
		endpoint = strings.TrimPrefix(endpoint, prefix)
	}

	return c.Config.BucketName + "." + endpoint
}
