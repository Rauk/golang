package main

import (
	"log"
	"github.com/minio/minio-go"
//	"fmt"
)

func main() {
	// Requests are always secure (HTTPS) by default. Set insecure=true to enable insecure (HTTP) access.
	// This boolean value is the last argument for New().

	// New returns an Amazon S3 compatible client object. API copatibality (v2 or v4) is automatically
	// determined based on the Endpoint value.

	//	accessKey := "AKIAJPGGJRNANVFRLCRA"
	//	secretKey := "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"
	//	url := "s3.amazonaws.com"

	accessKey := "439SQDG76BGBAM8ILSKR"
	secretKey := "zu7wZxwJYKUHMf7KJISKFSbvUC546Ge3KO3qVXbT"
	url := "10.47.2.3"

	svc, err := minio.NewV2(url, accessKey, secretKey, true)
	if err != nil {
		log.Fatalln(err)
	}
	err = svc.MakeBucket("test_keyspace-0", "", "")

	if err != nil {
		log.Fatalln(err)
	}

	buckets, err := svc.ListBuckets()
	if err != nil {
		log.Fatalln(err)
	}
	for _, bucket := range buckets {
		log.Println(bucket)
	}
}


