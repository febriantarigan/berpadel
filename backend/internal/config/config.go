package config

import "time"

type DynamoDB struct {
	Table        string
	BatchSize    int
	RetryBackoff time.Duration
	NumRetries   int
}
