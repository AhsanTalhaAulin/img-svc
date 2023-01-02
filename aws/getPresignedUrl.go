package aws

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"

	"img-svc/conn"
	"img-svc/domain"
)

func GetPresignedUrl(name string) (string, error) {

	// log.Println("Requesting Presigned Url for : ", name)
	svc := s3.New(conn.AwsClient.Sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(domain.BucketName),
		Key:    aws.String(name),
	})
	urlStr, err := req.Presign(5 * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
		return "Failed to sign request", err
	}

	// log.Println("The URL is---------->", urlStr)

	return urlStr, nil
}
