package sg

import (
	"github.com/golang/glog"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendMail sends an email to the configured mailing list.
func SendMail(subject string, body string) {
	if mailClient == nil {
		return
	}

	if len(mailingList) == 1 {
		message := mail.NewSingleEmail(from, subject, mailingList[0], body, body)

		response, err := mailClient.Send(message)
		if err != nil {
			if response != nil {
				if response.StatusCode == 401 || response.StatusCode > 500 {
					return
				}
			}

			// Some other error, needs to be reported!
			glog.Errorf("Error sending mail: %v", err)
		}
	} else {
		for _, email := range mailingList {
			message := mail.NewSingleEmail(from, subject, email, body, body)

			response, err := mailClient.Send(message)
			if err != nil {
				if response != nil {
					// If it's an access denied error, then just skip it
					if response.StatusCode == 401 || response.StatusCode > 500 {
						return
					}
				}

				// Some other error, needs to be reported!
				glog.Errorf("Error sending email to %s, %v", email, err)
			}
		}
	}
}
