package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/playwright-community/playwright-go"
	"github.com/slack-go/slack"
	"github.com/spf13/viper"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	c, err := LoadEnv()
	if err != nil {
		return err
	}

	i, err := FetchTraficAmount(c.NihonTsushinID, c.NihonTsushinPass)
	if err != nil {
		return err
	}

	api := slack.New(c.SlackToken)
	_, _, err = api.PostMessage(
		c.SlackChannelID,
		slack.MsgOptionText(fmt.Sprintf("日本通信SIMの利用データ量: %dMB", i), false),
	)
	if err != nil {
		return err
	}

	return nil
}

func FetchTraficAmount(id, pass string) (int, error) {
	runOption := &playwright.RunOptions{
		SkipInstallBrowsers: true,
	}

	err := playwright.Install(runOption)
	if err != nil {
		return 0, fmt.Errorf("playwright install failed: %v", err)
	}

	pw, err := playwright.Run()
	if err != nil {
		return 0, fmt.Errorf("playwright run failed: %v", err)
	}

	option := playwright.BrowserTypeLaunchOptions{
		Channel: playwright.String("chrome"),
	}

	browser, err := pw.Chromium.Launch(option)
	if err != nil {
		return 0, fmt.Errorf("playwright launch failed: %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	defer page.Close()
	if err != nil {
		return 0, fmt.Errorf("playwright new page failed: %v", err)
	}

	if err = GoToMyPage(page); err != nil {
		return 0, err
	}
	page.Goto(page.URL())

	if err = Login(page, id, pass); err != nil {
		return 0, err
	}
	page.Goto(page.URL())

	return FindAcount(page)
}

func FindAcount(page playwright.Page) (int, error) {
	entries, err := page.Locator("dl.mdesign > dd > span").All()
	if err != nil {
		return 0, fmt.Errorf("locator failed: %v", err)
	}
	for _, entry := range entries {
		t, _ := entry.InnerText()
		if strings.Contains(t, "高速通信") && strings.Contains(t, "MB") {
			x := strings.TrimSpace(strings.ReplaceAll(
				strings.ReplaceAll(t, "MB", ""), "高速通信", ""))
			i, e := strconv.Atoi(x)
			if e != nil {
				return 0, e
			}
			return i, nil
		}
	}
	return 0, nil
}

func Login(page playwright.Page, id, pass string) error {
	entries, err := page.Locator("input").All()
	if err != nil {
		return fmt.Errorf("locator failed: %v", err)
	}

	for _, entry := range entries {
		n, _ := entry.GetAttribute("name")
		if n == "josso_username" {
			entry.Fill(id)
		}
		if n == "josso_password" {
			entry.Fill(pass)
		}
	}
	ss, err := page.Locator("input[type=submit]").First().All()
	if err != nil {
		return fmt.Errorf("locator failed: %v", err)
	}

	for _, entry := range ss {
		entry.Click()
	}
	return nil
}

func GoToMyPage(page playwright.Page) error {
	if _, err := page.Goto(" https://www.nihontsushin.com/"); err != nil {
		return fmt.Errorf("playwright goto failed: %v", err)
	}

	entries, err := page.Locator("a").All()
	if err != nil {
		return fmt.Errorf("locator failed: %v", err)
	}

	for _, entry := range entries {
		t, _ := entry.InnerText()
		if t == "マイページ" {
			entry.Click()
			return nil
		}
	}
	return fmt.Errorf("MyPage Not Found")
}

func LoadEnv() (config *EnvConfigs, err error) {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}
	return
}

type EnvConfigs struct {
	NihonTsushinID   string `mapstructure:"NIHON_TSUSHIN_ID"`
	NihonTsushinPass string `mapstructure:"NIHON_TSUSHIN_PASS"`
	SlackToken       string `mapstructure:"SLACK_TOKEN"`
	SlackChannelID   string `mapstructure:"SLACK_CHANNEL_ID"`
}
