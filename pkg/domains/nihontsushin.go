package domains

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/playwright-community/playwright-go"
)

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
		Channel:  playwright.String("chrome"),
		Headless: playwright.Bool(true),
		Args:     []string{"--headless=new"},
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

	if err = GoToLoginPage(page); err != nil {
		return 0, err
	}

	if err = Login(page, id, pass); err != nil {
		return 0, err
	}

	page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
		State: playwright.LoadStateNetworkidle,
	})

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
	if err := page.Locator("input[name=username]").Fill(id); err != nil {
		return fmt.Errorf("fill username failed: %v", err)
	}
	if err := page.Locator("input[name=password]").Fill(pass); err != nil {
		return fmt.Errorf("fill password failed: %v", err)
	}
	if err := page.Locator("input[type=submit]").Click(); err != nil {
		return fmt.Errorf("click submit failed: %v", err)
	}
	return nil
}

func GoToLoginPage(page playwright.Page) error {
	if _, err := page.Goto("https://mypage.bmobile.ne.jp/"); err != nil {
		return fmt.Errorf("playwright goto failed: %v", err)
	}
	page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
		State: playwright.LoadStateNetworkidle,
	})
	return nil
}
