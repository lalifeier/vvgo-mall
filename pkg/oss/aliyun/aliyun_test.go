package aliyun

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

var config = Config{
	AccessKeyId:     "",
	AccessKeySecret: "",
	Bucket:          "",
}

func TestPutObject(t *testing.T) {
	c, err := New(&config)
	if err != nil {
		t.Fatal(err)
	}

	err = c.PutObject(
		"",
		"")
	if err != nil {
		t.Fatal(err)
	}
}

func TestListObjects(t *testing.T) {

	c, err := New(&config)
	if err != nil {
		t.Fatal(err)
	}
	marker := ""
	prefix := "file/202211"

	for {

		lsRes, err := c.Bucket.ListObjects(oss.Prefix(prefix), oss.Marker(marker))
		if err != nil {
			t.Fatal(err)
		}
		// 打印列举结果。默认情况下，一次返回100条记录。
		for _, object := range lsRes.Objects {
			fmt.Println("Bucket: ", object.Key)
			objectKey := object.Key

			filepath := "/tmp/oss/aliyun/" + objectKey

			if path.Ext(filepath) != ".jpg" {
				continue
			}

			dirPath := path.Dir(filepath)

			if !isExist(dirPath) {
				os.Mkdir(dirPath, os.ModePerm)
			}

			t.Log(filepath)

			c.GetObjectToFile(objectKey, "/tmp/oss/aliyun/"+objectKey)
		}
		if !lsRes.IsTruncated {
			break
		}

		marker = lsRes.NextMarker
	}
}
