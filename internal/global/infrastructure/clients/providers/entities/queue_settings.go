package entities

type QueueSettings struct {
	Key    string
	Secret string
	Region string
}

func NewQueueSettings(key, secret, region string) *QueueSettings {
	return &QueueSettings{
		Key:    key,
		Secret: secret,
		Region: region,
	}
}
