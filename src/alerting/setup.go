package alerting

import (
	"net/url"
	"os"

	"github.com/golang/glog"
	"github.com/nikita-petko/s3-forcer/alerting/discord"
	"github.com/nikita-petko/s3-forcer/alerting/sg"
	"github.com/nikita-petko/s3-forcer/alerting/sns"
	"github.com/nikita-petko/s3-forcer/flags"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var alertingEnabled bool = false

// SetupAlerting sets up alerting.
func SetupAlerting() {
	if *flags.SnsTopicArn != "" {

		if !*flags.AwsCredentialsFromProfile {
			if _, exists := os.LookupEnv("AWS_ACCESS_KEY"); !exists {
				glog.Fatal("If you are using AWS SNS, you must specify AWS_ACCESS_KEY environment variable.\n\n")
			}

			if _, exists := os.LookupEnv("AWS_SECRET_ACCESS_KEY"); !exists {
				glog.Fatal("If you are using AWS SNS, you must specify AWS_SECRET_ACCESS_KEY environment variable.")
			}
		}

		sns.SetupSnsTopic()

		alertingEnabled = true
	}

	if *flags.SendGridApiKey != "" {
		if *flags.SendGridFrom == "" {
			glog.Fatal("If you are using SendGrid, you must specify SENDGRID_FROM or -sendgrid-from.")
		}

		if *flags.SendGridFromEmail == "" {
			glog.Fatal("If you are using SendGrid, you must specify SENDGRID_FROM_EMAIL or -sendgrid-from-email.")
		}

		if *flags.SendGridMailingList == "" {
			glog.Fatal("If you are using SendGrid, you must specify SENDGRID_MAILING_LIST or -sendgrid-mailing-list.")
		}

		_, err := mail.ParseEmail(*flags.SendGridFromEmail)

		if err != nil {
			glog.Fatalf("If you are using SendGrid, the from email you supply must be valid: %s", err)
		}

		sg.SetupSendGrid()

		alertingEnabled = true
	}

	if *flags.DiscordWebHookUri != "" {
		_, err := url.Parse(*flags.DiscordWebHookUri)

		if err != nil {
			glog.Fatalf("If you are using Discord, the from webhook url you supply must be valid: %s", err)
		}

		discord.SetupDiscordClient()

		alertingEnabled = true
	}

	if !alertingEnabled {
		glog.Warning("AWS SNS, SendGrid and Discord alerting is disabled!")
	}

	glog.Info("Alerting setup complete!")
}
