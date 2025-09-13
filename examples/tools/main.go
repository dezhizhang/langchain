package main

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/tools"
	"office-helper/examples/tools/weather"
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

	weather := weather.New()
	calculator := new(tools.Calculator)

	tools := []tools.Tool{
		weather, calculator,
	}

	executor := agents.NewExecutor(agents.NewOneShotAgent(llm, tools))
	run, err := chains.Run(ctx, executor, "我想知道广州与北京的气温差")
	if err != nil {
		panic(err)
	}
	fmt.Println(run)
}
