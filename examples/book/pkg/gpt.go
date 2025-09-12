package pkg

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

type Chat struct {
	client *openai.Client
}

func NewChat() *Chat {
	cfg := openai.DefaultConfig(apiKey)
	cfg.BaseURL = url

	client := openai.NewClientWithConfig(cfg)
	return &Chat{client: client}
}

func (c *Chat) GenResponse(ctx context.Context, system, prompts string) (string, error) {
	resp, err := c.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:       openai.GPT4,
		Temperature: 0.4,
		N:           1,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: system,
			},
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: prompts,
			},
		},
	})
	if err != nil {
		panic(err)
	}

	return resp.Choices[0].Message.Content, err

}
