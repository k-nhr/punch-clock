package main

import (
	"github.com/nlopes/slack"
	"github.com/sclevine/agouti"
)

var login *Login

func (b *Bot) punchClock() ([]slack.Attachment, error) {

	// ブラウザはChromeを指定して起動
	driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	if err := driver.Start(); err != nil {
		return nil, err
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		return nil, err
	}
	if err := page.Navigate("https://qiita.com/"); err != nil {
		return nil, err
	}
	if err := page.Screenshot("c:\\Screenshot01.png"); err != nil {
		return nil, err
	}

	attachment := []slack.Attachment{slack.Attachment{
		Text: "打刻完了 :muscle:",
	}}

	return attachment, nil
}
