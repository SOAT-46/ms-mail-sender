package configuration

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/soat-46/ms-mail-sender/internal/global/infrastructure/clients/providers/entities"
)

func NewSQSClient(settings *entities.QueueSettings) (*sqs.Client, error) {
	staticProvider := credentials.NewStaticCredentialsProvider(settings.Key, settings.Secret, "")
	credentialProvider := config.WithCredentialsProvider(staticProvider)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		credentialProvider,
		config.WithRegion(settings.Region),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	client := sqs.NewFromConfig(cfg)
	return client, nil
}
