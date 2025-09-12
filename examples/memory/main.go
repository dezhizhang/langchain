package main

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/memory"
)

var (
	apiKey = "sk-83UzhDduYZygHMtlCC2TTsi4gjbUI90tjsEUO35kxt2XTpFI"
	url    = "https://poloai.top/v1"
	format = "要求：用中文输出，内容自然流畅，不要包含任何奇怪的符号、特殊空格或控制字符。"
)

func main() {

	ctx := context.Background()
	llm, err := openai.New(openai.WithToken(apiKey), openai.WithBaseURL(url))
	if err != nil {
		panic(err)
	}

	c := chains.NewConversation(llm, memory.NewConversationBuffer())
	_, err = chains.Run(ctx, c, "你好，我是贵州数擎科技有限公司")
	if err != nil {
		panic(err)
	}

	s, err := chains.Run(ctx, c, "我是谁，叫什么")
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

}
