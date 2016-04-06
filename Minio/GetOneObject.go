package main
import (
	"log"
	"os"
	"github.com/minio/minio-go"
	"io"
)

func main()  {
	svc, err := minio.NewV2(Url, AccessKey, SecretKey, true)
	if err != nil {
		log.Fatalln(err)
	}

	reader, err := svc.GetObject(MyBucket, MyObject)
	if err != nil {
		log.Fatalln(err)
	}
	defer reader.Close()

	localFile, err := os.Create("/Users/raunak.agarwal/Desktop/test2.txt")
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
