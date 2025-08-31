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
		"员工:{{.name}}\n: {{.dep}}\n介绍: {{.introduce}}",
		[]string{"name", "dep", "introduce"},
	)

	example := []map[string]string{
		{
			"name":      "tom",
			"dep":       "前端开发",
			"introduce": "大家好，我是lili 具有5年前端开发经验，很高兴加入",
		},
		{
			"name":      "数擎Ai",
			"dep":       "前端开发",
			"introduce": "大家好，我是数擎Ai 具有10年前端开发经验，很高兴加入,希大家多多关照",
		},
	}

	p, err := prompts.NewFewShotPrompt(
		template,
		example,
		nil,
		"请根据如下示例参考输出用户的个人介绍",
		"请你为:{{.sdep}}部门新入职员工:{{.sname}}设计一个自我介绍",
		[]string{"prefixId", "sdep", "sname"},
		map[string]interface{}{
			//"prefixId":  func() string { return "id" },
			//"prefixCtx": "测试",
		},
		"\n",
		prompts.TemplateFormatGoTemplate,
		false,
	)

	if err != nil {
		panic(err)
	}

	v, err := p.FormatPrompt(map[string]any{
		"sname": "数擎Ai",
		"sdep":  "前端开发",
	})
	if err != nil {
		panic(err)
	}

	text, err := llms.GenerateFromSinglePrompt(ctx, llm, v.String())
	if err != nil {
		panic(err)
	}
	fmt.Println(text)

}
