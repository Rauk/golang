package main

import (
	"log"

	"github.com/minio/minio-go"
)

func main() {
	accessKey := "439SQDG76BGBAM8ILSKR"
	secretKey := "zu7wZxwJYKUHMf7KJISKFSbvUC546Ge3KO3qVXbT"
	url := "10.47.2.3"
	object := "2016-03-16.135453.test-0000000102"
	myBucket := "raunaktestbucket"

	s3Client, err := minio.NewV2(url, accessKey, secretKey, true)
	if err != nil {
		log.Fatalln(err)
	}
	err = s3Client.RemoveObject(myBucket, object)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Success")
}
