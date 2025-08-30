package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"net/http"
)

var (
	apiKey = "sk-bL635gjwMN0Ft5VFdLhwPPAo3CMr4Rd0ncSQc19C59O0VuNT"
	url    = "https://poloai.top/v1"
)

func main() {
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = url

	config.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	client := openai.NewClientWithConfig(config)
	rsp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo, // 或者 openai.GPT4oMini
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: "hello world",
				},
			},
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(rsp.Choices)

}
