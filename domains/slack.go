package domains

import (
	"fmt"

	"github.com/slack-go/slack"
)

type SlackAPI struct {
	api            *slack.Client
	slackChannelID string
}

func NewSlackAPI(token, slackChannelID string) SlackAPI {
	return SlackAPI{
		api:            slack.New(token),
		slackChannelID: slackChannelID,
	}
}

func (s SlackAPI) Send(a Amount) error {
	_, _, err := s.api.PostMessage(
		s.slackChannelID,
		slack.MsgOptionBlocks(
			slack.NewSectionBlock(
				&slack.TextBlockObject{Type: "plain_text", Text: fmt.Sprintf("日本通信SIMの利用データ量: %dMB", a.CurrentAmount)},
				[]*slack.TextBlockObject{
					{Type: "plain_text", Text: fmt.Sprintf("平均使用 %.1fMB(%d日)", a.Average, a.CurrentDays)},
					{Type: "plain_text", Text: fmt.Sprintf("残り %d日~%s", a.RestDays, a.ExpectedEndDate)},
				},
				nil,
			),
		),
	)
	return err
}
