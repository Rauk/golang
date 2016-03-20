package main
import (
	"os"
	"log"
	"github.com/minio/minio-go"
	"io"
)


func main()  {

	accessKey := "439SQDG76BGBAM8ILSKR"
	secretKey := "zu7wZxwJYKUHMf7KJISKFSbvUC546Ge3KO3qVXbT"
	url := "10.47.2.3"
	myBucket := "raunaktestbucket"
	myObject := "3"

	svc, err := minio.NewV2(url, accessKey, secretKey, true)
	if err != nil {
		log.Fatalln(err)
	}

	reader, err := svc.GetObject(myBucket, myObject)
	if err != nil {
		log.Fatalln(err)
	}
	defer reader.Close()


	localFile, err := os.Create("/Users/raunak.agarwal/Desktop/text3.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer localFile.Close()

	stat, err := reader.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := io.CopyN(localFile, reader, stat.Size); err != nil {
		log.Fatalln(err)
	}

}

