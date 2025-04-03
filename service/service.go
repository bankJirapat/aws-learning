package service

import "aws-learning/config"

type Service struct {
	awsEC2 AwsEC2
	awsS3  AwsS3

	condig config.Config
}

func NewService(awsEC2 AwsEC2, awsS3 AwsS3, config config.Config) *Service {
	return &Service{
		awsEC2: awsEC2,
		awsS3:  awsS3,
		condig: config,
	}
}
