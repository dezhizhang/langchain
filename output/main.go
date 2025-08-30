package main

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/outputparser"
	"github.com/tmc/langchaingo/prompts"
)

var (
	apiKey             = "sk-83UzhDduYZygHMtlCC2TTsi4gjbUI90tjsEUO35kxt2XTpFI"
	url                = "https://poloai.top/v1"
	template           = "请你为{{.dep}} 部门新入职员工 {{.name}} 设计一个自我介绍"
	templateInputValue = []string{"dep", "name"}
)

func main() {
	ctx := context.Background()
	llm, err := openai.New(openai.WithToken(apiKey), openai.WithBaseURL(url))
	if err != nil {
		panic(err)
	}

	output := outputparser.NewStructured([]outputparser.ResponseSchema{
		{
			Name:        "content",
			Description: "介绍内容",
		},
		{
			Name:        "reason",
			Description: "为什么这么介绍",
		},
	})

	instructions := output.GetFormatInstructions()
	fmt.Println("output instructions:", instructions)

	promptStr := prompts.NewPromptTemplate(template+"\n"+instructions, templateInputValue)

	staff := map[string]any{
		"name": "数擎Ai",
		"dep":  "前端开发",
	}

	v, err := promptStr.FormatPrompt(staff)
	if err != nil {
		panic(err)
	}

	text, err := llms.GenerateFromSinglePrompt(ctx, llm, v.String())
	if err != nil {
		panic(err)
	}

	data, err := output.Parse(text)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)

}
