package main
import (
"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/aws/credentials"
	"fmt"
	"log"
"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	c := credentials.NewCredentials(&stubProvider{
		creds: credentials.Value{
			AccessKeyID:     "AKIAJRGV6GSEJYH3S55Q",
			SecretAccessKey: "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",
			SessionToken:    "",
		},
		expired: true,
	})

	svc := s3.New(session.New(&aws.Config{Region: aws.String("us-east-1"), Credentials: c}))

	result, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		log.Println("Failed to list buckets", err)
		return
	}
	log.Println("Buckets:")
	for _, bucket := range result.Buckets {
		log.Printf("%s : %s\n", aws.StringValue(bucket.Name), bucket.CreationDate)
		fmt.Println(aws.StringValue(bucket.Name), bucket.CreationDate)
	}

	//	ListObjects
}

type stubProvider struct {
creds   credentials.Value
expired bool
err     error
}
func (s *stubProvider) Retrieve() (credentials.Value, error) {
	s.expired = false
	return s.creds, s.err
}
func (s *stubProvider) IsExpired() bool {
	return s.expired
}
