package aws_s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c ClientS3) GetObject(ctx context.Context, awsCfg aws.Config, fileName string) ([]byte, string, error) {
	resp, err := c.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &c.cfg.Aws.S3.BucketName,
		Key:    &fileName,
	})
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	contentType := c.detectContentType(body, fileName)
	return body, contentType, nil
}
