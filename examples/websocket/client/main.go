package main

import "github.com/sashabaranov/go-openai"

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
