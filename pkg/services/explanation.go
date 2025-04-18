package services

import (
	"context"
	"fmt"

	"github.com/sa-giga/saga-cli/pkg/llm"
)

// ExplanationService は解説サービスを提供します
type ExplanationService struct {
	llmClient llm.LLM
	lang      string
}

// NewExplanationService は新しい解説サービスを作成します
func NewExplanationService(llmClient llm.LLM, lang string) *ExplanationService {
	return &ExplanationService{
		llmClient: llmClient,
		lang:      lang,
	}
}

// Process は入力テキストを解説します
func (s *ExplanationService) Process(ctx context.Context, inputText string) (string, error) {
	systemPrompt := fmt.Sprintf("あなたは優秀な解説者です。入力されたテキストの内容を%sで詳しく説明してください。複雑な概念をわかりやすく解説し、必要に応じて例を挙げてください。", s.lang)
	userPrompt := inputText

	return s.llmClient.Complete(ctx, systemPrompt, userPrompt)
}
