package main

import (
	"log"
	"github.com/minio/minio-go"
	"fmt"
)

func main() {
	s3Client, err := minio.NewV2(Url, AccessKey, SecretKey, true)
	if err != nil {
		log.Fatalln(err)
	}
	// Create a done channel to control 'ListObjects' go routine.
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)
	count := 0
	// List all objects from a bucket-name with a matching prefix.
	for object := range s3Client.ListObjects(MyBucket, "", false, doneCh) {
		count++
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
//		fmt.Println(1, object.ContentType)
//		fmt.Println(2, object.Err)
//		fmt.Println(3, object.ETag)
//		fmt.Println(4, object.Key)
//		fmt.Println(5, object.LastModified)
//		fmt.Println(6, object.Owner)
//		fmt.Println(7, object.Size)
//		fmt.Println(8, object.StorageClass)
		fmt.Println(object.Key)
	}
	fmt.Println("Count", count)
	return
}
