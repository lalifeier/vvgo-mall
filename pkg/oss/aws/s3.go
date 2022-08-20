package aws

import (
	"context"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	gooss "github.com/lalifeier/vvgo-mall/pkg/oss"
	"github.com/lalifeier/vvgo-mall/pkg/util"
	"github.com/lalifeier/vvgo-mall/pkg/util/convert"
)

// https://aws.github.io/aws-sdk-go-v2/docs/

var _ gooss.IOSS = (*Client)(nil)

type (
	Config struct {
		AccessKeyId     string
		AccessKeySecret string
		SessionToken    string
		Region          string
		Bucket          string
		Endpoint        string
	}

	Client struct {
		S3     *s3.Client
		Config *Config
	}
)

func (c Client) GetEndpoint() string {
	return c.Config.Endpoint
}

func (c Client) GetBucket() string {
	return c.Config.Bucket
}

func New(c *Config) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(c.Region),
		// config.WithEndpointResolverWithOptions(
		// 	aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		// 		return aws.Endpoint{
		// 			SigningRegion: "",
		// 			URL:           "",
		// 		}, nil
		// 	}),
		// ),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(c.AccessKeyId, c.AccessKeySecret, c.SessionToken)),
	)

	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return &Client{
		S3: client,
	}, nil
}

func (c *Client) CreateBucket(bucket string) error {
	input := &s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	}
	_, err := c.S3.CreateBucket(context.TODO(), input)
	return err
}

func (c *Client) DeleteBucket() error {
	input := &s3.DeleteBucketInput{
		Bucket: aws.String(c.GetBucket()),
	}
	_, err := c.S3.DeleteBucket(context.TODO(), input)
	return err
}

func (c *Client) CopyObject(destinationBucket, objectKey string) error {
	input := &s3.CopyObjectInput{
		Bucket:     aws.String(c.GetBucket()),
		CopySource: aws.String(destinationBucket),
		Key:        aws.String(objectKey),
	}
	_, err := c.S3.CopyObject(context.TODO(), input)
	return err
}

func (c *Client) PutObject(objectKey string, filePath string) (err error) {
	reader, err := convert.FileToReader(filePath)
	if err != nil {
		return err
	}

	input := &s3.PutObjectInput{
		Bucket: aws.String(c.GetBucket()),
		Key:    aws.String(objectKey),
		Body:   reader,
	}

	_, err = c.S3.PutObject(context.TODO(), input)
	return
}

func (c *Client) DeleteObject(objectKey string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(c.GetBucket()),
		Key:    aws.String(objectKey),
	}
	_, err := c.S3.DeleteObject(context.TODO(), input)
	return err
}

func (c *Client) GetObject(objectKey string) (io.ReadCloser, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(c.GetBucket()),
		Key:    aws.String(objectKey),
	}

	result, err := c.S3.GetObject(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return result.Body, nil
}

func (c *Client) GetObjectURL(objectKey string, expires int64) (string, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(c.GetBucket()),
		Key:    aws.String(objectKey),
	}

	psClient := s3.NewPresignClient(c.S3, s3.WithPresignExpires(time.Duration(expires)))

	resp, err := psClient.PresignGetObject(context.TODO(), input)
	if err != nil {
		return "", err
	}
	return resp.URL, nil
}

func (c *Client) GetObjectToFile(objectKey string, filePath string) (err error) {
	url, err := c.GetObjectURL(objectKey, 3600)
	if err != nil {
		return err
	}

	return util.WriteFile(url, filePath)
}

func (c *Client) ListObjects(prefix string) ([]*gooss.Object, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(c.GetBucket()),
		Prefix: aws.String(prefix),
	}

	result, err := c.S3.ListObjectsV2(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	var objects []*gooss.Object
	for _, item := range result.Contents {
		objects = append(objects, &gooss.Object{
			Key:          *item.Key,
			LastModified: item.LastModified,
			Size:         item.Size,
		})
	}
	return objects, err
}
