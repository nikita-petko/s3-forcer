package flags

func applyEnvironmentVariableFlags() {
	GetEnvironmentVariableOrFlag("MIN_CHANNEL_LENGTH", MinLength)
	GetEnvironmentVariableOrFlag("CHANNEL_CHAR_COMBINATIONS", CharCombinations)
	GetEnvironmentVariableOrFlag("CHANNEL_PREFIX", ChannelPrefix)
	GetEnvironmentVariableOrFlag("ATTEMPT_MILESTONE", AttemptMilestone)
	GetEnvironmentVariableOrFlag("USE_S3_DIRECTLY", UseS3Directly)
	GetEnvironmentVariableOrFlag("WORKER_THREADS", WorkerThreads)

	GetEnvironmentVariableOrFlag("SEND_GRID_API_KEY", SendGridApiKey)
	GetEnvironmentVariableOrFlag("SEND_GRID_FROM", SendGridFrom)
	GetEnvironmentVariableOrFlag("SEND_GRID_FROM_EMAIL", SendGridFromEmail)
	GetEnvironmentVariableOrFlag("SEND_GRID_MAILING_LIST", SendGridMailingList)

	GetEnvironmentVariableOrFlag("SNS_TOPIC_ARN", SnsTopicArn)
	GetEnvironmentVariableOrFlag("AWS_CREDENTIALS_FROM_PROFILE", AwsCredentialsFromProfile)

	GetEnvironmentVariableOrFlag("DISCORD_WEBHOOK_URI", DiscordWebHookUri)
}
