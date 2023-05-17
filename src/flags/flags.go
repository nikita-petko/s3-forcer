package flags

import (
	"flag"
	"time"
)

var (
	// BindAddressIpv4 is the address to bind the prometheus metrics server to.
	BindAddressIpv4 = flag.String("bind-metrics-server", ":8080", "Address to bind the prometheus metrics server to.")

	// HelpFlag prints the usage.
	HelpFlag = flag.Bool("help", false, "Print usage.")

	// MinLength is the minimum character length of channel name combinations to bruteforce. (enviornment variable: MIN_CHANNEL_LENGTH)
	MinLength = flag.Int("min-length", 1, "Minimum character length of channel name combinations to bruteforce. (enviornment variable: MIN_CHANNEL_LENGTH)")

	// CharCombinations are the characters (in a single string) to use for channel name combinations. (environment variable: CHANNEL_CHAR_COMBINATIONS)
	CharCombinations = flag.String("char-combinations", "abcdefghijklmnopqrstuvwxyz1", "Characters (in a single string) to use for channel name combinations. (environment variable: CHANNEL_CHAR_COMBINATIONS)")

	// ChannelPrefix is the prefix for channel names (i.e. zwinplayer64). (environment variable: CHANNEL_PREFIX)
	ChannelPrefix = flag.String("channel-prefix", "z", "Prefix for channel names (i.e. zwinplayer64). (environment variable: CHANNEL_PREFIX)")

	// InvalidateCache invalidates the cache on cache initialization.
	InvalidateCache = flag.Bool("invalidate-cache", false, "Should the cache be invalidated on cache initialization.")

	// AttemptMilestone is the number of attempts between cache writes and milestone prints. (enviornment variable: ATTEMPT_MILESTONE)
	AttemptMilestone = flag.Int("attempt-milestone", 5000, "Number of attempts between cache writes and milestone prints. (enviornment variable: ATTEMPT_MILESTONE)")

	// UseS3Directly uses s3.amazonaws.com instead of the rbx-cdn-provider (rbxcdn.com.) (environment variable: USE_S3_DIRECTLY)
	UseS3Directly = flag.Bool("use-s3-directly", false, "Should we use S3 directly instead of the rbx-cdn-provider. (environment variable: USE_S3_DIRECTLY)")

	// WorkerThreads is the number of max semaphore thread workers to use. (environment variable: WORKER_THREADS)
	WorkerThreads = flag.Int64("worker-threads", 250, "Number of max thread workers to use. (environment variable: WORKER_THREADS)")

	// S3ClientTimeout is the timeout for the S3 HTTP client when receiving a response. (environment variable: S3_CLIENT_TIMEOUT)
	S3ClientTimeout = flag.Duration("s3-client-timeout", time.Second*30, "The timeout for the S3 HTTP client when receiving a response. (environment variable: S3_CLIENT_TIMEOUT)")

	// Alerting

	// SendGridApiKey is the SendGrid API key. This is optional. (environment variable: SENDGRID_API_KEY)
	SendGridApiKey = flag.String("sendgrid-api-key", "", "The SendGrid API key. This is optional. (environment variable: SENDGRID_API_KEY)")

	// SendGridFrom is the name to use as the sender. This is required if the API Key is specified. (environment variable: SENDGRID_FROM)
	SendGridFrom = flag.String("sendgrid-from", "", "The name to use as the sender. This is required if the API Key is specified. (environment variable: SENDGRID_FROM)")

	// SendGridFromEmail is the email address to use as the sender. This is required if the API Key is specified. (environment variable: SENDGRID_FROM_EMAIL)
	SendGridFromEmail = flag.String("sendgrid-from-email", "", "The email address to use as the sender. This is required if the API Key is specified. (environment variable: SENDGRID_FROM_EMAIL)")

	// SendGridMailingList is the mailing list to send the emails to. This is required if the API Key is specified. (environment variable: SENDGRID_MAILING_LIST)
	SendGridMailingList = flag.String("sendgrid-mailing-list", "", "The mailing list to send the emails to. This is required if the API Key is specified. (environment variable: SENDGRID_MAILING_LIST)")

	// SnsTopicArn is yhe ARN to the topic created in AWS SNS. This is optional. Needs AWS_ACCESS_KEY and AWS_SECRET_ACCESS_KEY. (environment variable: SNS_TOPIC_ARN)
	SnsTopicArn = flag.String("sns-topic-arn", "", "The ARN to the topic created in AWS SNS. This is optional. Needs AWS_ACCESS_KEY and AWS_SECRET_ACCESS_KEY. (environment variable: SNS_TOPIC_ARN)")

	// AwsCredentialsFromProfile will load the AWS credentials from the system profile instead of environment variables. (enviornment variable: AWS_CREDENTIALS_FROM_PROFILE)
	AwsCredentialsFromProfile = flag.Bool("aws-credentials-from-profile", false, "Is the AWS SNS Credentials coming from a profile file? If not use enviornment variables. (environment variable: AWS_CREDENTIALS_FROM_PROFILE)")

	// DiscordWebHookUri is the url that was generated when creating a Discord WebHook. (environment variable: DISCORD_WEBHOOK_URI)
	DiscordWebHookUri = flag.String("discord-webhook-uri", "", "The url that was generated when creating a Discord WebHook. (environment variable: DISCORD_WEBHOOK_URI)")
)

const FlagsUsageString string = `
	[-h|--help] [--min-length[=1]] [--char-combinations[=abcdefghijklmnopqrstuvwxyz1]] [--channel-prefix[=z]]
	[--invalidate-cache] [--attempt-milestone[=5000]] [--use-s3-directly] [--workers-threads[=250]]
	[--sendgrid-api-key[=]] [--sendgrid-from[=]] [--sendgrid-from-email[=]] [--sendgrid-mailing-list[=]]
	[--sns-topic-arn[=]] [--aws-credentials-from-profile]
	[--discord-webhook-uri[=]]`
