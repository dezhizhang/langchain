package main

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
)

var (
	apiKey = "sk-83UzhDduYZygHMtlCC2TTsi4gjbUI90tjsEUO35kxt2XTpFI"
	url    = "https://poloai.top/v1"
)

func main() {

	ctx := context.Background()
	llm, err := openai.New(openai.WithToken(apiKey), openai.WithBaseURL(url))
	if err != nil {
		panic(err)
	}

	template := prompts.NewPromptTemplate(
		"请你为{{.dep}}部门新入职员工{{.name}} 设计一个自我介绍",
		[]string{"dep", "name"},
	)

	staff := map[string]any{
		"name": "数擎Ai",
		"dep":  "产品研发",
	}

	prompt, err := template.FormatPrompt(staff)
	if err != nil {
		panic(err)
	}
	text, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt.String())
	if err != nil {
		panic(err)
	}
	fmt.Println(text)

}
