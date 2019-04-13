package catchers

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

var (
	ErrSlackCallFailed = errors.New("Slack Call failed")
)

type slackMessage struct {
	Message string `json:"text"`
}

//Slack is a catcher for your panic handler middleware.
// You shouldn't instantiate one directly, you should use NewSlack()
type Slack struct {
	webhookUrl string
	httpClient *http.Client
}

//Newslack takes a webhook url that you want to write to when a panic happens.
//Returns a *Slack.
func NewSlack(webhookurl string) *Slack {
	return &Slack{
		webhookUrl: webhookurl,
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

//HandlePanic is the function that will be called in the panic handle middleware if your program panics.
// It takes a message that will be written to the webhook if your server does panic.
// HandlePanic can return an error if it cannot marshall the message into JSON, building a request fails, or the call to Slack fails.
func (s Slack) HandlePanic(message string) error {
	b, err := json.Marshal(slackMessage{Message: message})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", s.webhookUrl, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	res, err := s.httpClient.Do(req)

	if err != nil || res.StatusCode >= 400 {
		return ErrSlackCallFailed
	}
	return err
}
