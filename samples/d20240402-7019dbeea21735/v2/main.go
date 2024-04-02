package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	err := Very色んな処理()
	msg, err := HandleErrorV2(err, PostMessage) // ここで注入
	if err != nil {
		panic(err)
	}

	fmt.Println(msg)
}

// 引数のfuncにてDIっぽいことを実現する
func HandleErrorV2(err error, postMessage func(string, slack.MsgOption) (string, string, error)) (string, error) {
	var msg string
	if err == nil {
		msg = "Very色んな処理が完了しました\n"
	} else {
		msg = fmt.Sprintf("エラーが発生しました: %v\n", err)
	}

	if _, _, err = postMessage("#general", slack.MsgOptionText(msg, true)); err != nil {
		return "", fmt.Errorf("slackへのメッセージ送信に失敗しました: %v", err)
	}

	return msg, nil
}

func PostMessage(channel string, msgOption slack.MsgOption) (string, string, error) {
	return slack.New(os.Getenv("SLACK_TOKEN")).PostMessage(channel, msgOption)
}

func Very色んな処理() error {
	// 実際にはめっちゃ色んな処理がここに書かれてることにしてください
	return nil
}
