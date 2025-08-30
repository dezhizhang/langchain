package main

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

var (
	apiKey = "sk-83UzhDduYZygHMtlCC2TTsi4gjbUI90tjsEUO35kxt2XTpFI"
	url    = "https://poloai.top/v1"
)

func main() {
	//llm, err := openai.New(openai.WithToken(apiKey), openai.WithBaseURL(url))
	//if err != nil {
	//	panic(err)
	//}
	//
	////text, err := llms.GenerateFromSinglePrompt(context.Background(), llm, "贵州数擎科技目前做什么")
	////if err != nil {
	////	panic(err)
	////}
	////
	////fmt.Println(text)
	//
	//rsp, err := llm.GenerateContent(context.Background(), []llms.MessageContent{
	//	llms.TextParts(llms.ChatMessageTypeSystem, "我是贵州数擎科技有限公司"),
	//	llms.TextParts(llms.ChatMessageTypeHuman, "请说出你需要的帮助"),
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(rsp.Choices[0].Content)

	llm, err := openai.New(openai.WithToken(apiKey), openai.WithBaseURL(url))
	if err != nil {
		panic(err)
	}

	rsp, err := llm.GenerateContent(context.Background(), []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "贵州数引擎科技有限公司"),
		llms.TextParts(llms.ChatMessageTypeHuman, "我有什么可以帮助你的吗"),
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Choices[0].Content)

}
