package client

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

var (
	apiKey = "sk-83UzhDduYZygHMtlCC2TTsi4gjbUI90tjsEUO35kxt2XTpFI"
	url    = "https://poloai.top/v1"
)

type Chat struct {
	client *openai.Client
}

func NewChat() *Chat {
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = url

	client := openai.NewClientWithConfig(config)
	return &Chat{client: client}
}

func (c *Chat) Message(ctx context.Context, msg string) (string, error) {
	resp, err := c.client.CreateChatCompletion(
		ctx, openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: msg,
				},
			},
		})
	if err != nil {
		panic(err)
	}

	return resp.Choices[0].Message.Content, nil
}
