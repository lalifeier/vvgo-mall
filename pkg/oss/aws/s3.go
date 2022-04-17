package aws

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// https://aws.github.io/aws-sdk-go-v2/docs/

type (
	Config struct {
		AccessKeyId     string
		AccessKeySecret string
		SessionToken    string
		Region          string
		Bucket          string
	}

	Client struct {
		S3     *s3.Client
		Config *Config
	}
)

func New(c *Config) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(c.Region),
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

func (c *Client) CopyObject(destinationBucket, objectKey string) error {
	input := &s3.CopyObjectInput{
		Bucket:     aws.String(c.Config.Bucket),
		CopySource: aws.String(destinationBucket),
		Key:        aws.String(objectKey),
	}
	_, err := c.S3.CopyObject(context.TODO(), input)
	return err
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
		Bucket: aws.String(c.Config.Bucket),
	}
	_, err := c.S3.DeleteBucket(context.TODO(), input)
	return err
}

func (c *Client) PutObject(objectKey string, reader io.Reader) (err error) {
	input := &s3.PutObjectInput{
		Bucket: aws.String(c.Config.Bucket),
		Key:    aws.String(objectKey),
		Body:   reader,
	}

	_, err = c.S3.PutObject(context.TODO(), input)
	return err
}

func (c *Client) DeleteObject(objectKey string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(c.Config.Bucket),
		Key:    aws.String(objectKey),
	}
	_, err := c.S3.DeleteObject(context.TODO(), input)
	return err
}

func (c *Client) GetObject(objectKey string, expires int64) (io.ReadCloser, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(c.Config.Bucket),
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
		Bucket: aws.String(c.Config.Bucket),
		Key:    aws.String(objectKey),
	}

	psClient := s3.NewPresignClient(c.S3, s3.WithPresignExpires(time.Duration(expires)))

	resp, err := psClient.PresignGetObject(context.TODO(), input)
	if err != nil {
		return "", err
	}
	return resp.URL, nil
}

func (c *Client) ListObjects() error {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(c.Config.Bucket),
	}

	result, err := c.S3.ListObjectsV2(context.TODO(), input)
	if err != nil {
		return err
	}

	for _, item := range result.Contents {
		fmt.Println("Name:          ", *item.Key)
		fmt.Println("Last modified: ", *item.LastModified)
		fmt.Println("Size:          ", item.Size)
		fmt.Println("Storage class: ", item.StorageClass)
	}
	return nil
}
