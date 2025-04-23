package contracts

import "github.com/aws/aws-sdk-go-v2/service/sqs"

type SQSClient interface {
	ReceiveMessages(sqsURL string) (*sqs.ReceiveMessageOutput, error)
	DeleteMessage(sqsURL string, receiptHandle *string) error
}
