package aws

// func UploadtoS3(name string, imgFile []byte) error {

// 	uploader := s3manager.NewUploader(conn.AwsClient.Sess)

// 	result, err := uploader.Upload(&s3manager.UploadInput{
// 		Bucket: aws.String(domain.BucketName),
// 		Key:    aws.String(name),
// 		Body:   bytes.NewReader(imgFile),
// 	})
// 	if err != nil {
// 		log.Printf("Failed to upload file, %v\n", err)
// 		return err
// 	}
// 	log.Printf("File uploaded to, %v\n", result.Location)

// 	return nil
// }
