package main

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
)

var (
	apiKey             = "sk-83UzhDduYZygHMtlCC2TTsi4gjbUI90tjsEUO35kxt2XTpFI"
	url                = "https://poloai.top/v1"
	format             = "要求：用中文输出，内容自然流畅，不要包含任何奇怪的符号、特殊空格或控制字符。"
	template           = "请你为{{.dep}}部门新入职员工{{.name}}设计一个自我介绍 "
	templateInputValue = []string{"dep", "name"}
)

func main() {
	ctx := context.Background()
	llm, err := openai.New(openai.WithToken(apiKey), openai.WithBaseURL(url), openai.WithModel("gpt-4o-mini"))
	if err != nil {
		panic(err)
	}

	// 生成提示词模板
	prompt := prompts.NewPromptTemplate(template+format, templateInputValue)
	// 创建链
	chain := chains.NewLLMChain(llm, prompt)
	res, err := chains.Call(ctx, chain, map[string]any{
		"name": "数擎AI",
		"dep":  "前端开发",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
