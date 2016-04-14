package main
import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	config := aws.NewConfig().WithRegion("us-east-1")
//	sess := session.New(&aws.Config{
//		Region: aws.String("us-east-1")})
	sess := session.New(config)
	svc := s3.New(sess)

	i := 0
	err := svc.ListObjectsPages(&s3.ListObjectsInput{
		Bucket: &os.Args[1],
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		fmt.Println("Page,", i)

		for _, obj := range p.Contents {
			fmt.Println("Object:", *obj.Key)
		}
		return true
	})
	if err != nil {
		fmt.Println("failed to list objects", err)
		return
	}
}
