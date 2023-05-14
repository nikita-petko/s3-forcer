package alerting

import (
	"fmt"

	"github.com/nikita-petko/s3-forcer/alerting/discord"
	"github.com/nikita-petko/s3-forcer/alerting/sg"
	"github.com/nikita-petko/s3-forcer/alerting/sns"
)

// Alert alerts to any of the 3 different alerting methods.
func Alert(title, description string) {
	if !alertingEnabled {
		return
	}

	sns.PublishToSnsTopic(fmt.Sprintf("%s\n\n%s", title, description))
	sg.SendMail(title, description)
	discord.PublishEmbed(title, description, 0x0)
}
