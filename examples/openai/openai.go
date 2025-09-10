package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

var (
	apiKey = "sk-83UzhDduYZygHMtlCC2TTsi4gjbUI90tjsEUO35kxt2XTpFI"
	url    = "https://poloai.top/v1"
)

func main() {
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = url

	client := openai.NewClientWithConfig(config)

	rsp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "hello world",
				},
			},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Choices[0].Message.Content)

}
