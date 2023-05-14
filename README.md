# s3-forcer

Brute forcer for channel names on Roblox's S3 client server.

## Running

This repository provides [releases](https://github.com/nikita-petko/s3-forcer/releases) and [docker images](https://hub.docker.com/repository/docker/mfdlabs/s3-forcer)

## Building

Ensure you have [Go 1.20.3+](https://go.dev/dl/)

1. Clone the repository via `git`:

    ```txt
    git clone git@github.com:nikita-petko/s3-forcer.git
    cd s3-forcer
    ```

2. Build via [make](https://www.gnu.org/software/make/)

    ```txt
    make build-debug WITH_STDERR=1
    ```

## Usage

`cd src && go run main.go --help` (use the build binary found in the bin directory if you downloaded a prebuilt or built it yourself)

```txt
Usage: s3-forcer
Build Mode: <build_mode>
Commit: <commit_sha> 
        [-h|--help] [--min-length[=1]] [--char-combinations[=abcdefghijklmnopqrstuvwxyz1]] [--channel-prefix[=z]]
        [--invalidate-cache] [--attempt-milestone[=5000]] [--use-s3-directly] [--workers-threads[=250]]
        [--sendgrid-api-key[=]] [--sendgrid-from[=]] [--sendgrid-from-email[=]] [--sendgrid-mailing-list[=]]
        [--sns-topic-arn[=]] [--aws-credentials-from-profile]
        [--discord-webhook-uri[=]]

  -alsologtostderr
        log to standard error as well as files
  -attempt-milestone int
        Number of attempts between cache writes and milestone prints. (enviornment variable: ATTEMPT_MILESTONE) (default 5000)
  -aws-credentials-from-profile
        Is the AWS SNS Credentials coming from a profile file? If not use enviornment variables. (environment variable: AWS_CREDENTIALS_FROM_PROFILE)
  -bind-metrics-server string
        Address to bind the prometheus metrics server to. (default ":8080")
  -channel-prefix string
        Prefix for channel names (i.e. zwinplayer64). (environment variable: CHANNEL_PREFIX) (default "z")
  -char-combinations string
        Characters (in a single string) to use for channel name combinations. (environment variable: CHANNEL_CHAR_COMBINATIONS) (default "abcdefghijklmnopqrstuvwxyz1")
  -discord-webhook-uri string
        The url that was generated when creating a Discord WebHook. (environment variable: DISCORD_WEBHOOK_URI)
  -help
        Print usage.
  -invalidate-cache
        Should the cache be invalidated on cache initialization.
  -log_backtrace_at value
        when logging hits line file:N, emit a stack trace
  -log_dir string
        If non-empty, write log files in this directory
  -log_link string
        If non-empty, add symbolic links in this directory to the log files
  -logbuflevel int
        Buffer log messages logged at this level or lower (-1 means don't buffer; 0 means buffer INFO only; ...). Has limited applicability on non-prod platforms.
  -logtostderr
        log to standard error instead of files
  -min-length int
        Minimum character length of channel name combinations to bruteforce. (enviornment variable: MIN_CHANNEL_LENGTH) (default 1)
  -sendgrid-api-key string
        The SendGrid API key. This is optional. (environment variable: SENDGRID_API_KEY)
  -sendgrid-from string
        The name to use as the sender. This is required if the API Key is specified. (environment variable: SENDGRID_FROM)
  -sendgrid-from-email string
        The email address to use as the sender. This is required if the API Key is specified. (environment variable: SENDGRID_FROM_EMAIL)
  -sendgrid-mailing-list string
        The mailing list to send the emails to. This is required if the API Key is specified. (environment variable: SENDGRID_MAILING_LIST)
  -sns-topic-arn string
        The ARN to the topic created in AWS SNS. This is optional. Needs AWS_ACCESS_KEY and AWS_SECRET_ACCESS_KEY. (environment variable: SNS_TOPIC_ARN)
  -stderrthreshold value
        logs at or above this threshold go to stderr (default 2)
  -use-s3-directly
        Should we use S3 directly instead of the rbx-cdn-provider. (environment variable: USE_S3_DIRECTLY)
  -v value
        log level for V logs
  -vmodule value
        comma-separated list of pattern=N settings for file-filtered logging
  -worker-threads int
        Number of max thread workers to use. (environment variable: WORKER_THREADS) (default 250)
```

### Extra Notes

* To quit the daemon and stop the request process, you need to use a keyboard interrupt (`Ctrl+C`) in your terminal, or just cancel the process directly through whatever you're using.

* If you get immediate request errors, try toggling `--workers-threads` to a number lower than the default of `250`! You may need to mess with this flag a bit to get the most desirable result for your own personal machine/network.

* If you're annoyed at how often the "Milestone of (n) reached.." message is, you can toggle `--attempt-milestone` (which is the multiplier this uses) to a higher number! *This flag is also used for how often the session should write to the cache, which is what s3-forcer uses to pick up (roughly) where you left off on a certain length cycle of your current prefix+char configuration, incase you exit the program or there's an uncaught error for whatever reason.*

<sub>

*P.S., just try to figure everything else out yourself from the `--help` prompt; this is provided with zero extra support or warranty <3*

</sub>

## License

```txt
Copyright 2023 Nikita Petko <petko@vmminfra.net>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```