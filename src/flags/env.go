package flags

func applyEnvironmentVariableFlags() {
	getEnvironmentVariableOrFlag("MIN_CHANNEL_LENGTH", MinLength)
	getEnvironmentVariableOrFlag("CHANNEL_CHAR_COMBINATIONS", CharCombinations)
	getEnvironmentVariableOrFlag("CHANNEL_PREFIX", ChannelPrefix)
	getEnvironmentVariableOrFlag("ATTEMPT_MILESTONE", AttemptMilestone)
	getEnvironmentVariableOrFlag("USE_S3_DIRECTLY", UseS3Directly)
	getEnvironmentVariableOrFlag("WORKER_THREADS", WorkerThreads)

	getEnvironmentVariableOrFlag("SEND_GRID_API_KEY", SendGridApiKey)
	getEnvironmentVariableOrFlag("SEND_GRID_FROM", SendGridFrom)
	getEnvironmentVariableOrFlag("SEND_GRID_FROM_EMAIL", SendGridFromEmail)
	getEnvironmentVariableOrFlag("SEND_GRID_MAILING_LIST", SendGridMailingList)

	getEnvironmentVariableOrFlag("SNS_TOPIC_ARN", SnsTopicArn)
	getEnvironmentVariableOrFlag("AWS_CREDENTIALS_FROM_PROFILE", AwsCredentialsFromProfile)

	getEnvironmentVariableOrFlag("DISCORD_WEBHOOK_URI", DiscordWebHookUri)
}
