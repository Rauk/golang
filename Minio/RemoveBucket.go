package main
import "log"
import "github.com/minio/minio-go"

func main()  {
	accessKey := "439SQDG76BGBAM8ILSKR"
	secretKey := "zu7wZxwJYKUHMf7KJISKFSbvUC546Ge3KO3qVXbT"
	url := "10.47.2.3"
	myBucket := "raunaktestbucket"

	s3Client, err := minio.NewV2(url, accessKey, secretKey, true)
	if err != nil {
		log.Fatalln(err)
	}

	// This operation will only work if your bucket is empty.
	err = s3Client.RemoveBucket(myBucket)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Success")
}
