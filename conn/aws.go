package conn

import (
	"img-svc/domain"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type awsSession struct {
	Sess *session.Session
}

var AwsClient awsSession

func ConnectAWS() error {

	sess, err := session.NewSession(
		&aws.Config{
			Credentials: credentials.NewStaticCredentials(domain.AccessKeyId, domain.SecretAccessKey, ""),
			Region:      aws.String(domain.RegionName),
		},
	)
	if err != nil {
		log.Fatal(err)
		return err
	}

	AwsClient.Sess = sess

	return nil
}
