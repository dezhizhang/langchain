package book

import (
	"context"
	"encoding/json"
	"fmt"
	"office-helper/examples/book/pkg"
	"office-helper/examples/book/prompts"
	"os"
)

type BookGpt struct {
	gpt *pkg.Chat
}

func NewBookGpt() *BookGpt {
	return &BookGpt{gpt: &pkg.Chat{}}
}

// Abstract 生成摘要
func (b *BookGpt) Abstract(ctx context.Context, title string, sectionStr string) (*BodyCount, error) {
	prompt := fmt.Sprintf("主题:%s\n 所有章节：%s", title, sectionStr)
	return b.BodyContent(ctx, prompts.Abstract, prompt)
}

// Summary 生成提示词
func (b *BookGpt) Summary(ctx context.Context, title string, sectionStr string) (*BodyCount, error) {
	prompt := fmt.Sprintf("主题:%s\n 所有章节：%s", title, sectionStr)
	return b.BodyContent(ctx, prompts.Summary, prompt)
}

func (b *BookGpt) SectionContent(ctx context.Context, title string, sections []*Section) ([]*BodyCount, error) {
	res := make([]*BodyCount, 0, len(sections))

	for i, _ := range sections {
		fmt.Println("生成章节", sections[i].Title)

		prompt := fmt.Sprintf("主题:%s\n 章节:%s \n内容字数要求：%s", title, sections[i].Title, sections[i].Count)

		bodyContent, err := b.BodyContent(ctx, prompts.SectionContent, prompt)
		if err != nil {
			return nil, err
		}
		res = append(res, bodyContent)

		// 判断生成的内容是不是完成了
		count := bodyContent.Count

		for sections[i].Count > count {
			fmt.Println("继续完善章节", sections[i].Title)

			//perfectPrompt := fmt.Sprintf("")
		}

		fmt.Println("生成章节", sections[i].Title, "完成")
	}
	return res, nil
}

// BodyContent 生成文章内容
func (b *BookGpt) BodyContent(ctx context.Context, system, prompt string) (*BodyCount, error) {
	resp, err := b.gpt.GenResponse(ctx, system, prompt)
	if err != nil {
		return nil, err
	}

	var res *BodyCount
	if err := json.Unmarshal([]byte(resp), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// Section 生成文章章阳数
func (b *BookGpt) Section(ctx context.Context, title string) ([]*Section, error) {
	resp, err := b.gpt.GenResponse(ctx, prompts.Section, title)
	if err != nil {
		return nil, err
	}

	var sections []*Section
	if err := json.Unmarshal([]byte(resp), &sections); err != nil {
		return nil, err
	}
	return sections, nil
}

// SaveContent 保存文件
func (b *BookGpt) SaveContent(content, filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return os.WriteFile(filePath, []byte(content), 0644)
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
