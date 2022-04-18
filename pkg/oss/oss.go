package oss

import (
	"io"
	"time"
)

type Object struct {
	Key          string
	Type         string
	Size         int64
	LastModified *time.Time
}

type StorageInterface interface {
	PutObject(objectKey string, filepath string) error
	DeleteObject(objectKey string) error
	GetObject(objectKey string) (io.ReadCloser, error)
	GetObjectURL(objectKey string, expires int64) (string, error)
	ListObjects(objectKey string) ([]*Object, error)
	DownloadObject(objectKey, filePath string) error
}
