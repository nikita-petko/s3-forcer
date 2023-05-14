package sg

import (
	"strings"
	"sync"

	"github.com/golang/glog"
	"github.com/nikita-petko/s3-forcer/flags"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	mailingList []*mail.Email    = []*mail.Email{}
	mailClient  *sendgrid.Client = nil
	from        *mail.Email      = nil
	setupOnce   sync.Once
)

// SetupSendGrid sets up the SendGrid mailing client.
// Can only occur once.
func SetupSendGrid() {
	setupOnce.Do(func() {
		// Setup the sendgrid client
		mailClient = sendgrid.NewSendClient(*flags.SendGridApiKey)

		// Setup the from email
		from = mail.NewEmail(*flags.SendGridFrom, *flags.SendGridFromEmail)

		// Setup the mailing list
		for _, email := range strings.Split(*flags.SendGridMailingList, ",") {
			email, err := mail.ParseEmail(email)
			if err != nil {
				glog.Fatalf("Invalid sendgrid mailing list email address: %s!", email)
			}

			mailingList = append(mailingList, email)
		}

		if len(mailingList) == 0 {
			glog.Fatal("Invalid sendgrid mailing list, no emails found!")
		}

		glog.Info("SendGrid setup complete!")
	})
}
