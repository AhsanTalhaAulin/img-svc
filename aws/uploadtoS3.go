package aws

import (
	"bytes"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"img-svc/domain"
)

func UploadtoS3(name string, imgFile []byte) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1")},
	)
	if err != nil {
		log.Printf("Could not create session, %v\n", err)
		return err
	}

	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(domain.BucketName),
		Key:    aws.String(name),
		Body:   bytes.NewReader(imgFile),
	})
	if err != nil {
		log.Printf("Failed to upload file, %v\n", err)
		return err
	}
	log.Printf("File uploaded to, %v\n", result.Location)

	return nil
}
