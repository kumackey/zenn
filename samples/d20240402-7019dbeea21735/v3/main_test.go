package main

import (
	"errors"
	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockSlackClient struct{}

func (m *mockSlackClient) PostMessage(channelID string, options ...slack.MsgOption) (string, string, error) {
	// 実際の通知をさせず、何も起こさない
	return "", "", nil
}

func TestHandleErrorV2(t *testing.T) {
	msg, err := HandleErrorV3(errors.New("何かしらのエラー"), &mockSlackClient{})
	assert.Nil(t, err)
	assert.Equal(t, "エラーが発生しました: 何かしらのエラー\n", msg)
}
