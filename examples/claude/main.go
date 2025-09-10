package main

import (
	"fmt"
	"gitee.com/dn-jinmin/claude"
	"log"
)

var (
	apiKey = "sk-UASfaZ1vN3ookOlmQaqnMUPJDPV6VXEjnkOPW0sWgYPmSlL6"
	url    = "https://poloai.top"
)

func main() {

	client, err := claude.NewClient(apiKey, claude.WithClientBaseURL(url))
	if err != nil {
		panic(err)
	}
	fmt.Println(client)
	req := &claude.CompletionRequest{
		Prompt: "\n\nHumab: hello world\n\nAssistant",
		Model:  claude.ClaudeV2,
	}
	resp, err := client.Complete(req)
	if err != nil {
		log.Fatalf("error sending completion request: %+v", err) // 打印完整错误
	}
	fmt.Println(resp)
}
