package catchers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"net/http/httputil"
	"time"
)

var (
	ErrSlackCallFailed = errors.New("Slack Call failed")
)

type slackMessage struct {
	Message string `json:"text"`
}

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

func (s Slack) HandlePanic(message string) error {
	b, err := json.Marshal(slackMessage{Message: message})
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	req, err := http.NewRequest("POST", s.webhookUrl, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	res, err := s.httpClient.Do(req)

	br,_:= httputil.DumpRequest(req, true)
	fmt.Println(string(br))

	st, _ := httputil.DumpResponse(res, true)
	fmt.Println(string(st))
	if err != nil {
		return ErrSlackCallFailed
	}
	return err
}
