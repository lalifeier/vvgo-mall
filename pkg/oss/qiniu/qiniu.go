package qiniu

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	gooss "github.com/lalifeier/vvgo-mall/pkg/oss"
	"github.com/lalifeier/vvgo-mall/pkg/util"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// https://developer.qiniu.com/kodo/1238/go

var _ gooss.IOSS = (*Client)(nil)

type (
	Config struct {
		AccessKeyId     string // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		AccessKeySecret string // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		SessionToken    string
		Region          string
		Bucket          string
		Endpoint        string
		PrivateURL      bool
	}

	Client struct {
		Config        *Config
		StorageCfg    storage.Config
		Mac           *qbox.Mac
		BucketManager *storage.BucketManager
	}
)

func (c Client) GetEndpoint() string {
	return c.Config.Endpoint
}

func (c Client) GetBucket() string {
	return c.Config.Bucket
}

var zonedata = map[string]*storage.Zone{
	"huadong":    &storage.ZoneHuadong,
	"huabei":     &storage.ZoneHuabei,
	"huanan":     &storage.ZoneHuanan,
	"beimei":     &storage.ZoneBeimei,
	"xinjiapo":   &storage.ZoneXinjiapo,
	"fogcneast1": &storage.ZoneFogCnEast1,
}

func New(c *Config) (*Client, error) {
	mac := qbox.NewMac(c.AccessKeyId, c.AccessKeySecret)
	cfg := storage.Config{
		UseHTTPS:      true,
		UseCdnDomains: true,
		Zone:          &storage.ZoneHuabei,
	}
	if z, ok := zonedata[c.Region]; ok {
		cfg.Zone = z
	} else {
		return nil, fmt.Errorf("region %s not found", c.Region)
	}

	bucketManager := storage.NewBucketManager(mac, &cfg)

	return &Client{
		Config:        c,
		StorageCfg:    cfg,
		Mac:           mac,
		BucketManager: bucketManager,
	}, nil
}

func (c *Client) PutObject(objectKey string, filePath string) (err error) {
	putPolicy := storage.PutPolicy{
		Scope: c.GetBucket(),
	}

	upToken := putPolicy.UploadToken(c.Mac)

	cfg := storage.Config{}
	resumeUploader := storage.NewResumeUploaderV2(&cfg)

	ret := storage.PutRet{}
	recorder, err := storage.NewFileRecorder(os.TempDir())
	if err != nil {
		fmt.Println(err)
		return
	}

	putExtra := storage.RputV2Extra{
		Recorder: recorder,
	}
	err = resumeUploader.PutFile(context.Background(), &ret, upToken, objectKey, filePath, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ret.Key, ret.Hash)
	return err
}

func (c *Client) DeleteObject(objectKey string) error {
	return c.BucketManager.Delete(c.GetBucket(), objectKey)
}

func (c *Client) DeleteObjects(objectKeys []string) error {
	deleteOps := make([]string, 0, len(objectKeys))
	for _, key := range objectKeys {
		deleteOps = append(deleteOps, storage.URIDelete(c.GetBucket(), key))
	}
	_, err := c.BucketManager.Batch(deleteOps)
	return err
}

func (c *Client) GetObject(objectKey string) (io.ReadCloser, error) {
	url, err := c.GetObjectURL(objectKey, 0)
	if err != nil {
		return nil, err
	}

	result, err := http.Get(url)
	if err == nil && result.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("file %s not found", objectKey)
	}
	return result.Body, nil
}

func (c *Client) GetObjectURL(objectKey string, expires int64) (url string, err error) {
	if c.Config.PrivateURL {
		deadline := time.Now().Unix() + expires
		url = storage.MakePrivateURL(c.Mac, c.GetEndpoint(), objectKey, deadline)
		return
	}

	url = storage.MakePublicURL(c.GetEndpoint(), objectKey)
	return
}

func (c *Client) GetObjectToFile(objectKey string, filePath string) error {
	url, err := c.GetObjectURL(objectKey, 3600)
	if err != nil {
		return err
	}

	return util.WriteFile(url, filePath)
}

func (c *Client) ListObjects(prefix string) ([]*gooss.Object, error) {
	var objects []*gooss.Object

	marker := ""
	for {
		entries, _, nextMarker, hasNext, err := c.BucketManager.ListFiles(c.GetBucket(), prefix, "", marker, 1000)
		if err != nil {
			fmt.Println("list error,", err)
			break
		}
		//print entries
		for _, entry := range entries {
			t := time.Unix(entry.PutTime, 0)
			objects = append(objects, &gooss.Object{
				Key:          entry.Key,
				Size:         entry.Fsize,
				LastModified: &t,
			})
		}
		if hasNext {
			marker = nextMarker
		} else {
			//list end
			break
		}
	}
	return objects, nil
}
