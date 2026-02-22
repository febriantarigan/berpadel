package config

func GetDynamoDB() *DynamoDB {
	return &DynamoDB{
		Table:        configManager.GetString("dynamodb.table"),
		BatchSize:    configManager.GetInt("dynamodb.batch_size"),
		RetryBackoff: configManager.GetDuration("dynamodb.retry_backoff"),
		NumRetries:   configManager.GetInt("dynamodb.num_retries"),
	}
}
