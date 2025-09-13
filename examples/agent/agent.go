package main

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/tools"
	"github.com/tmc/langchaingo/tools/serpapi"
	"log"
)

var (
	apiKey     = "sk-83UzhDduYZygHMtlCC2TTsi4gjbUI90tjsEUO35kxt2XTpFI"
	url        = "https://poloai.top/v1"
	serpapiKey = "74178d16d3b68d93b29454b6ba4ae118a2f0b39a0d9c51d701191cca14b8b8c7"
)

func handlePanic(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	ctx := context.Background()
	llm, err := openai.New(openai.WithToken(apiKey), openai.WithBaseURL(url))
	handlePanic(err)

	// 定义工具
	tool, err := serpapi.New(serpapi.WithAPIKey(serpapiKey))
	handlePanic(err)

	// 定义计算
	calculator := new(tools.Calculator)
	t := []tools.Tool{tool, calculator}

	executor := agents.NewExecutor(agents.NewOneShotAgent(llm, t))
	run, err := chains.Run(ctx, executor, "贵州数擎科技有限公司")
	handlePanic(err)

	fmt.Println(run)

}
