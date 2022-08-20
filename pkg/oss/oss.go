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

type IOSS interface {
	PutObject(objectKey string, filePath string) error
	DeleteObject(objectKey string) error
	GetObject(objectKey string) (io.ReadCloser, error)
	GetObjectURL(objectKey string, expires int64) (string, error)
	GetObjectToFile(objectKey, filePath string) error
	ListObjects(prefix string) ([]*Object, error)
}
