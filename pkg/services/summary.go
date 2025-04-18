package services

import (
	"context"
	"fmt"

	"github.com/sa-giga/saga-cli/pkg/llm"
)

// SummaryService は要約サービスを提供します
type SummaryService struct {
	llmClient llm.LLM
	lang      string
}

// NewSummaryService は新しい要約サービスを作成します
func NewSummaryService(llmClient llm.LLM, lang string) *SummaryService {
	return &SummaryService{
		llmClient: llmClient,
		lang:      lang,
	}
}

// Process は入力テキストを要約します
func (s *SummaryService) Process(ctx context.Context, inputText string) (string, error) {
	systemPrompt := fmt.Sprintf("あなたは優秀な要約者です。入力されたテキストの要点を簡潔に%sでまとめてください。重要な情報を漏らさず、冗長な表現は避けてください。", s.lang)
	userPrompt := inputText

	return s.llmClient.Complete(ctx, systemPrompt, userPrompt)
}
