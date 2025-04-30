package aws_s3

import (
	"aws-learning/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type ClientS3 struct {
	cfg           config.Config
	client        *s3.Client
	presignClinet *s3.PresignClient
}

func NewClientS3(cfg config.Config, awsCfg aws.Config) *ClientS3 {
	clientS3 := s3.NewFromConfig(awsCfg)
	presignClientS3 := s3.NewPresignClient(clientS3)
	return &ClientS3{
		cfg:           cfg,
		client:        clientS3,
		presignClinet: presignClientS3,
	}
}
