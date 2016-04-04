package main

import (
	"log"
	"github.com/minio/minio-go"
	"fmt"
	"sync"
)

func main() {
	accessKey := "439SQDG76BGBAM8ILSKR"
	secretKey := "zu7wZxwJYKUHMf7KJISKFSbvUC546Ge3KO3qVXbT"
	url := "10.47.2.3"
	myBucket := "minio-bucket1"
//	myBucket := "raunaktestbucket"

	s3Client, err := minio.NewV2(url, accessKey, secretKey, true)
	if err != nil {
		log.Fatalln(err)
	}
	// Create a done channel to control 'ListObjects' go routine.
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)
	var arr []minio.ObjectInfo

	// List all objects from a bucket-name with a matching prefix.
	for object := range s3Client.ListObjects(myBucket, "", false, doneCh) {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		arr = append(arr, object)
	}
	var wg sync.WaitGroup
	for j, _ := range arr {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			err = s3Client.RemoveObject(myBucket, arr[j].Key)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("Success", arr[j].Key)
		} (j)
	}
	log.Println("Count", len(arr))
	wg.Wait()
	return
}
