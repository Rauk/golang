package main

import (
	"log"

	"github.com/minio/minio-go"
)

func main() {
	s3Client, err := minio.NewV2(Url, AccessKey, SecretKey, true)
	if err != nil {
		log.Fatalln(err)
	}
	err = s3Client.RemoveObject(MyBucket, MyObject)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Success")
}
