package catchers

import (
	"bytes"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

var (
	ErrSlackCallFailed = errors.New("Slack Call failed")
)

type Slack struct {
	webhookUrl string
	httpClient *http.Client
}

func NewSlack(webhookurl string) *Slack {
	return &Slack{
		webhookUrl: webhookurl,
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (s Slack) Handle(message string) error {
	req, err := http.NewRequest("POST", s.webhookUrl, bytes.NewBuffer([]byte(message)))
	if err != nil {
		return err
	}
	_, err = s.httpClient.Do(req)
	return ErrSlackCallFailed
}
