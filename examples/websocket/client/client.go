package client

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
)

var (
	apiKey = "sk-83UzhDduYZygHMtlCC2TTsi4gjbUI90tjsEUO35kxt2XTpFI"
	url    = "https://poloai.top/v1"
)

type Chat struct {
	client *openai.Client
}

// NewChat 创建客户端
func NewChat() *Chat {
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = url

	client := openai.NewClientWithConfig(config)
	return &Chat{client: client}
}

func (c *Chat) Message(ctx context.Context, msg string) (string, error) {
	//resp, err := c.client.CreateCompletion(ctx, openai.CompletionRequest{
	//	Model:     openai.GPT3Dot5Turbo,
	//	Prompt:    msg,
	//	MaxTokens: 10,
	//})
	//if err != nil {
	//	panic(err)
	//}

	resp, err := c.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: msg,
				},
			},
		},
	)

	if err != nil {
		panic(err)
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("no choices found")
	}

	return resp.Choices[0].Message.Content, nil
}
