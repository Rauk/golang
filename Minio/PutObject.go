package main
import (
	"os"
	"log"
	"github.com/minio/minio-go"
//	"io"
)


func main()  {

	accessKey := "439SQDG76BGBAM8ILSKR"
	secretKey := "zu7wZxwJYKUHMf7KJISKFSbvUC546Ge3KO3qVXbT"
	url := "10.47.2.3"
	myBucket := "minio-bucket1"
//	myBucket := "raunaktestbucket"
	myObject := "object1"
//	arrStrings := []string {"abc1", "abc2", "abc3", "abc4", "ab5", "abc6", "abc7"}

	svc, err := minio.NewV2(url, accessKey, secretKey, true)
	if err != nil {
		log.Fatalln(err)
	}
	object, err := os.Open("/Users/raunak.agarwal/Desktop/test.txt")
//	for _, arrString := range arrStrings {
//		go func(str string) {
//			if err != nil {
//				log.Fatalln(err)
//			}
//
//			_, err = svc.PutObject(myBucket, str, object, "application/octet-stream")
//
//			if err != nil {
//				log.Fatalln(err)
//			}
//		} (arrString)
//	}
	defer object.Close()

	//	object, err := os.Open("/Users/raunak.agarwal/Desktop/text2.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer object.Close()

	n, err := svc.PutObject(myBucket, myObject, object, "application/octet-stream")

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Uploaded", myObject, " of size: ", n, "Successfully.")
}

