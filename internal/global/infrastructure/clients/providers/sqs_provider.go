package providers

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSClientProvider struct {
	client *sqs.Client
}

func NewSQSClientProvider(client *sqs.Client) *SQSClientProvider {
	return &SQSClientProvider{
		client: client,
	}
}

func (svc *SQSClientProvider) ReceiveMessages(sqsURL string) (*sqs.ReceiveMessageOutput, error) {
	results, err := svc.client.ReceiveMessage(
		context.Background(),
		&sqs.ReceiveMessageInput{
			QueueUrl:            &sqsURL,
			MaxNumberOfMessages: 10,
		},
	)
	return results, err
}

func (svc *SQSClientProvider) DeleteMessage(sqsURL string, receiptHandle *string) error {
	deleteParams := &sqs.DeleteMessageInput{
		QueueUrl:      &sqsURL,
		ReceiptHandle: receiptHandle,
	}

	_, err := svc.client.DeleteMessage(context.TODO(), deleteParams)
	if err != nil {
		return fmt.Errorf("failed to delete message,. Reason: %v", err)
	}
	return nil
}
