package aliyun

import (
	"io"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	gooss "github.com/lalifeier/vvgo-mall/pkg/oss"
	"github.com/lalifeier/vvgo-mall/pkg/util/convert"
)

// https://help.aliyun.com/product/31815.html

var _ gooss.IOSS = (*Client)(nil)

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

func (c Client) GetBucket() string {
	return c.Config.Bucket
}

func (c Client) GetEndpoint() string {
	if c.Config.Endpoint != "" {
		if strings.HasSuffix(c.Config.Endpoint, "aliyuncs.com") {
			return c.GetBucket() + "." + c.Config.Endpoint
		}
		return c.Config.Endpoint
	}

	endpoint := c.Bucket.Client.Config.Endpoint
	for _, prefix := range []string{"https://", "http://"} {
		endpoint = strings.TrimPrefix(endpoint, prefix)
	}

	return c.GetBucket() + "." + endpoint
}

func New(config *Config) (*Client, error) {
	if config.Endpoint == "" {
		config.Endpoint = "http://oss-cn-hangzhou.aliyuncs.com"
	}

	if config.UseCname {
		config.ClientOptions = append(config.ClientOptions, oss.UseCname(config.UseCname))
	}

	client, err := oss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret, config.ClientOptions...)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(config.Bucket)
	if err != nil {
		return nil, err
	}

	return &Client{
		Bucket: bucket,
		Config: config,
	}, err
}

func (c *Client) PutObject(objectName string, filePath string) (err error) {
	reader, err := convert.FileToReader(filePath)
	if err != nil {
		return err
	}
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

func (c *Client) GetObjectToFile(objectKey, filePath string) error {
	return c.Bucket.GetObjectToFile(objectKey, filePath)
}

func (c *Client) ListObjects(prefix string) ([]*gooss.Object, error) {
	results, err := c.Bucket.ListObjects(oss.Prefix(prefix))
	if err != nil {
		return nil, err
	}

	var objects []*gooss.Object
	for _, result := range results.Objects {
		objects = append(objects, &gooss.Object{
			Key:          result.Key,
			Type:         result.Type,
			Size:         result.Size,
			LastModified: &result.LastModified,
		})
	}

	return objects, err
}

func (c *Client) IsObjectExist(objectKey string) (isExist bool, err error) {
	return c.Bucket.IsObjectExist(objectKey)
}
