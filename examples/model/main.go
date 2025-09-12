package main

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

var (
	apiKey = "sk-UASfaZ1vN3ookOlmQaqnMUPJDPV6VXEjnkOPW0sWgYPmSlL6"
	url    = "https://poloai.top/v1"
)

func main() {
	llm, err := openai.New(openai.WithToken(apiKey), openai.WithBaseURL(url))
	if err != nil {
		panic(err)
	}

	rsp, err := llm.GenerateContent(context.Background(), []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "你是一个诗人"),
		llms.TextParts(llms.ChatMessageTypeHuman, "请用诗词描述我很帅气"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Choices[0].Content)
}
