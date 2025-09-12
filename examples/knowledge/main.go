package main

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/textsplitter"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/redisvector"
	"os"
)

var (
	apiKey = "sk-83UzhDduYZygHMtlCC2TTsi4gjbUI90tjsEUO35kxt2XTpFI"
	url    = "https://poloai.top/v1"
)

func main() {
	//fb, err := os.Open("./employee.pdf")

	ctx := context.Background()

	f, err := os.Open("./employee.pdf")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		panic(err)
	}

	// 将文件进行加载
	p := documentloaders.NewPDF(f, stat.Size())

	// 将文件进行切分
	chunk, err := p.LoadAndSplit(ctx, textsplitter.NewRecursiveCharacter(
		textsplitter.WithChunkSize(200),
		textsplitter.WithChunkOverlap(1),
	))
	if err != nil {
		panic(err)
	}

	// 创建大模型
	llm, err := openai.New(openai.WithToken(apiKey), openai.WithBaseURL(url))
	if err != nil {
		panic(err)
	}

	embedder, err := embeddings.NewEmbedder(llm)
	if err != nil {
		panic(err)
	}

	// 介建向量数据库
	store, err := redisvector.New(
		ctx,
		redisvector.WithEmbedder(embedder),
		redisvector.WithConnectionURL("redis://106.15.74.79:6379"),
		redisvector.WithIndexName("knowledge", true),
	)
	if err != nil {
		panic(err)
	}

	_, err = store.AddDocuments(ctx, chunk)
	if err != nil {
		panic(err)
	}

	// 用户发起提问
	qa := chains.NewRetrievalQAFromLLM(llm, vectorstores.ToRetriever(store, 1))
	res, err := chains.Run(ctx, qa, "工作时间与考勤")
	if err != nil {
		panic(err)
	}
	fmt.Println("res=", res)

}
