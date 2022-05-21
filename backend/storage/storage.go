package storage

import "github.com/gofiber/storage/s3"

var S3 *s3.Storage

func InitStorage() {
	S3 = s3.New(s3.Config{
		Bucket:   "images",
		Endpoint: "http://localhost:9000",
		Region:   "us-east-1",
		Reset:    false,
	})
}
