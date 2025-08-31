package openai

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

var (
	apiKey = "sk-bL635gjwMN0Ft5VFdLhwPPAo3CMr4Rd0ncSQc19C59O0VuNT"
	url    = "https://poloai.top/v1"
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

func (c *Chat) RoleMessage(ctx context.Context, prompts string) (string, error) {
	rsp, err := c.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: prompts,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}
	return rsp.Choices[0].Message.Content, nil
}
