package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	c := slack.New(os.Getenv("SLACK_TOKEN"))

	err := Very色んな処理()
	msg, err := HandleErrorV3(err, c)
	if err != nil {
		panic(err)
	}

	fmt.Println(msg)
}

func HandleErrorV3(err error, slackClient slackClientInterface) (string, error) {
	var msg string
	if err == nil {
		msg = "Very色んな処理が完了しました\n"
	} else {
		msg = fmt.Sprintf("エラーが発生しました: %v\n", err)
	}

	if _, _, err = slackClient.PostMessage("#general", slack.MsgOptionText(msg, true)); err != nil {
		return "", fmt.Errorf("slackへのメッセージ送信に失敗しました: %v", err)
	}

	return msg, nil
}

type slackClientInterface interface {
	// 以下のシグネチャをinterfaceとして定義しただけ
	// https://pkg.go.dev/github.com/slack-go/slack#Client.PostMessage
	// Goは明示的なimplementsが不要なので、ライブラリに合うinterfaceをこちらで勝手に定義して使える
	PostMessage(channelID string, options ...slack.MsgOption) (string, string, error)
}

func Very色んな処理() error {
	// 実際にはめっちゃ色んな処理がここに書かれてることにしてください
	return nil
}
