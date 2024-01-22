package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	c "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func SetupDynoConfig() *dynamodb.Client {
	cfg, err := c.LoadDefaultConfig(context.Background(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{URL: "http://localhost:8000"}, nil
		})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     "apaja",
				SecretAccessKey: "apaja",
				SessionToken:    "apaja",
				Source:          "apaja",
			},
		}),
	)

	if err != nil {
		log.Fatal("Failed iniitalize credentials configuration", err)
	}

	return dynamodb.NewFromConfig(cfg)
}
