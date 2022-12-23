package domain

import "os"

var BucketName = "talha-test-image-s3"

var AccessKeyId = os.Getenv("AWS_ACCESS_KEY_ID_TES")
var SecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY_TEST")
var RegionName = os.Getenv("AWS_DEFAULT_REGION_TEST")
