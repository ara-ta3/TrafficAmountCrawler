package main

import (
	"log"
	"traffic-costs-crawler/pkg/domains"

	"github.com/golang-module/carbon/v2"
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
	err = domains.NewSlackAPI(c.SlackToken, c.SlackChannelID).Send(a)
	if err != nil {
		return err
	}

	return nil
}
