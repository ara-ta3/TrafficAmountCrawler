package main

import (
	"fmt"
	"log"

	"github.com/ara-ta3/trafic-costs-crawler/domains"
	"github.com/golang-module/carbon/v2"
	"github.com/slack-go/slack"
)

func main() {
	err := run(carbon.Now())
	if err != nil {
		log.Fatal(err)
	}
}

func run(now carbon.Carbon) error {
	c, err := domains.LoadEnv()
	if err != nil {
		return err
	}

	i, err := domains.FetchTraficAmount(c.NihonTsushinID, c.NihonTsushinPass)
	if err != nil {
		return err
	}

	a := domains.CalculateAmount(now, i)
	err = PostMessage(
		c.SlackToken,
		c.SlackChannelID,
		a,
	)
	if err != nil {
		return err
	}

	return nil
}

func PostMessage(token, slackChannelID string, a domains.Amount) error {
	api := slack.New(token)
	_, _, err := api.PostMessage(
		slackChannelID,
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
