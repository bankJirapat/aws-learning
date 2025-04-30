package aws_s3

import (
	"bytes"
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

func (c ClientS3) PutObject(ctx context.Context, awsCfg aws.Config, fileName string, data []byte) error {
	_, err := c.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &c.cfg.Aws.S3.BucketName,
		Key:         &fileName,
		Body:        bytes.NewReader(data),
		ContentType: aws.String(c.detectContentType(data, fileName)),
	})
	return err
}

func (c ClientS3) PresignUrlObject(ctx context.Context, awsCfg aws.Config) {
	req, err := c.presignClinet.PresignPostObject(ctx, &s3.PutObjectInput{}, func(opts *s3.PresignPostOptions) {})

}
