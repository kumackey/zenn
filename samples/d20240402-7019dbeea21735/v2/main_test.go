package main

import (
	"errors"
	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandleErrorV2(t *testing.T) {
	mockPostMessage := func(channel string, msgOption slack.MsgOption) (string, string, error) {
		// 実際の通知をさせず、何も起こさない
		return "", "", nil
	}

	msg, err := HandleErrorV2(errors.New("何かしらのエラー"), mockPostMessage)
	assert.Nil(t, err)
	assert.Equal(t, "エラーが発生しました: 何かしらのエラー\n", msg)
}
