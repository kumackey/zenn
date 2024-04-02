package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	c := slack.New(os.Getenv("SLACK_TOKEN"))

	err := Very色んな処理()
	msg, err := HandleErrorV1(c, err)
	if err != nil {
		panic(err)
	}

	fmt.Println(msg)
}

func HandleErrorV1(c *slack.Client, err error) (string, error) {
	var msg string
	if err == nil {
		msg = "Very色んな処理が完了しました\n"
	} else {
		msg = fmt.Sprintf("エラーが発生しました: %v\n", err)
	}

	if _, _, err = c.PostMessage("#general", slack.MsgOptionText(msg, true)); err != nil {
		return "", fmt.Errorf("slackへのメッセージ送信に失敗しました: %v", err)
	}

	return msg, nil
}

func Very色んな処理() error {
	// 実際にはめっちゃ色んな処理がここに書かれてることにしてください
	return nil
}
