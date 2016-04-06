package main
import "log"
import "github.com/minio/minio-go"

func main()  {
	s3Client, err := minio.NewV2(Url, AccessKey, SecretKey, true)
	if err != nil {
		log.Fatalln(err)
	}

	// This operation will only work if your bucket is empty.
	err = s3Client.RemoveBucket(MyBucket)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Success")
}
