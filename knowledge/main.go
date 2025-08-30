package main

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/textsplitter"
	"github.com/tmc/langchaingo/vectorstores/redisvector"
	"os"
)

var (
	apiKey = "sk-83UzhDduYZygHMtlCC2TTsi4gjbUI90tjsEUO35kxt2XTpFI"
	url    = "https://poloai.top/v1"
)

func main() {
	ctx := context.Background()

	fb, err := os.Open("./employee.pdf")
	if err != nil {
		panic(err)
	}
	defer fb.Close()
	content, err := fb.Stat()
	handleError(err)

	p := documentloaders.NewPDF(fb, content.Size())
	split, err := p.LoadAndSplit(
		ctx,
		textsplitter.NewRecursiveCharacter(
			textsplitter.WithChunkSize(200),
			textsplitter.WithChunkOverlap(1),
		),
	)

	handleError(err)

	// 创建embedding
	llm, err := openai.New(openai.WithToken(apiKey), openai.WithBaseURL(url))
	handleError(err)

	embedder, err := embeddings.NewEmbedder(llm)
	handleError(err)

	// 设置向量数据库
	store, err := redisvector.New(ctx,
		redisvector.WithEmbedder(embedder),
		redisvector.WithConnectionURL("redis://localhost:6379"),
		redisvector.WithIndexName("knowledge", true),
	)
	handleError(err)
	documents, err := store.AddDocuments(ctx, split)
	handleError(err)
	fmt.Println("documents:", documents)

}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
