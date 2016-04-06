package main
import (
	"os"
	"log"
	"github.com/minio/minio-go"
//	"io"
)


func main()  {
//	arrStrings := []string {"abc1", "abc2", "abc3", "abc4", "ab5", "abc6", "abc7"}

	svc, err := minio.NewV2(Url, AccessKey, SecretKey, true)
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

	n, err := svc.PutObject(MyBucket, MyObject, object, "application/octet-stream")

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Uploaded", MyObject, " of size: ", n, "Successfully.")
}

