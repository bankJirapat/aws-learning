package config

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	aws_config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	App      App      `env:"APP"`
	Database Database `env:"DATABASE"`
	Redis    Redis    `env:"REDIS"`
	Aws      Aws      `env:"AWS"`
	Secret   Secret   `env:"SECRET"`
}

type Aws struct {
	Region string `env:"REGION"`
	S3     S3     `env:"S3"`
}

type Secret struct {
	Aws AwsSecret `env:"AWS"`
}

type App struct {
	Name        string `env:"NAME"`
	Port        string `env:"PORT"`
	Environment string `env:"ENVIRONMENT"`
}

type AwsSecret struct {
	AccessKey       string `env:"ACCESS_KEY"`
	SecretAccessKey string `env:"SECRET_ACCESS_KEY"`
}

type Database struct{}

type Redis struct{}

type S3 struct {
	BucketName string `env:"BUCKET_NAME"`
}

func Init() (Config, error) {
	godotenv.Load(CONFIG_ENV, SECRET_ENV)

	config := Config{}
	if err := env.Parse(&config); err != nil {
		return Config{}, fmt.Errorf("failed to parse env: %w", err)
	}
	return config, nil
}

func (cfg Config) InitAwsConfig() (aws.Config, error) {
	// Manually set AWS credentials from environment variables
	customResolver := aws.NewCredentialsCache(aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		return aws.Credentials{
			AccessKeyID:     cfg.Secret.Aws.AccessKey,
			SecretAccessKey: cfg.Secret.Aws.SecretAccessKey,
			Source:          AWS_SOURCE,
		}, nil
	}))

	// Load AWS configuration
	awsCfg, err := aws_config.LoadDefaultConfig(context.TODO(),
		aws_config.WithCredentialsProvider(customResolver),
		aws_config.WithRegion(cfg.Aws.Region),
	)
	if err != nil {
		return aws.Config{}, fmt.Errorf("failed to load AWS config: %w", err)
	}
	return awsCfg, nil
}
