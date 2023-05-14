package discord

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golang/glog"
	"github.com/nikita-petko/s3-forcer/flags"
)

// PublishEmbed publishes an embed to the configured
// WebHook. If color is 0x0, it will default to
// 0x3498db.
func PublishEmbed(title, description string, color int) {
	if discordHttpClient == nil {
		return // not configured
	}

	if color == 0 {
		color = 0x3498db
	}

	strColor := strconv.FormatInt(int64(color), 16)

	data := &postData{
		Embeds: []*embed{
			{
				Title:       title,
				Description: description,
				Color:       strColor,
			},
		},
	}

	marshalledData, err := json.Marshal(data)
	if err != nil {
		glog.Fatalf("Error while marshalling request: %v", err)
	}

	request, err := http.NewRequest("POST", *flags.DiscordWebHookUri, bytes.NewBuffer(marshalledData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		glog.Fatalf("Error while creating request: %v", err)
	}

	response, err := discordHttpClient.Do(request)
	if err != nil {
		glog.Errorf("Error while sending request: %v", err)
		return
	}

	if response.StatusCode != 204 {
		glog.Errorf("Error when sending Discord embed: %s", response.Status)
	}
}
