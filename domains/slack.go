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
	t := fmt.Sprintf("日本通信SIMの利用データ量: %dMB", a.CurrentAmount)
	m1 := ""
	m2 := ""
	if a.ExpireByTheLastDay() {
		m1 = fmt.Sprintf("平均使用 %.1fMB(%d日)", a.Average, a.CurrentDays)
		m2 = fmt.Sprintf("残り %d日~%s", a.RestDays, a.ExpectedEnd.ToDateString())
	} else {
		v := float64(a.CurrentAmount) / float64(a.RestDaysUntilEnd)
		m1 = fmt.Sprintf("平均使用可能量 %.1fMB(%d日)", v, a.RestDaysUntilEnd)
		m2 = fmt.Sprintf("残り %d日~%s", a.RestDaysUntilEnd, a.End.ToDateString())
	}
	_, _, err := s.api.PostMessage(
		s.slackChannelID,
		slack.MsgOptionBlocks(
			slack.NewSectionBlock(
				&slack.TextBlockObject{Type: "plain_text", Text: t},
				[]*slack.TextBlockObject{
					{Type: "plain_text", Text: m1},
					{Type: "plain_text", Text: m2},
				},
				nil,
			),
		),
	)
	return err
}
