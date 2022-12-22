package aws

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"img-svc/domain"
)

func GetPresignedUrl(name string) (string, error) {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1")},
	)

	if err != nil {
		log.Printf("Could not create session, %v\n", err)
		return "Could not create session", nil
	}

	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(domain.BucketName),
		Key:    aws.String(name),
	})
	urlStr, err := req.Presign(1 * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
		return "Failed to sign request", err
	}

	log.Println("The URL is---------->", urlStr)

	return urlStr, nil
}
