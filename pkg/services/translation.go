package services

import (
	"context"
	"fmt"

	"github.com/sa-giga/saga-cli/pkg/llm"
)

// TranslationService は翻訳サービスを提供します
type TranslationService struct {
	llmClient llm.LLM
	lang      string
}

// NewTranslationService は新しい翻訳サービスを作成します
func NewTranslationService(llmClient llm.LLM, lang string) *TranslationService {
	return &TranslationService{
		llmClient: llmClient,
		lang:      lang,
	}
}

// Process は入力テキストを翻訳します
func (s *TranslationService) Process(ctx context.Context, inputText string) (string, error) {
	systemPrompt := fmt.Sprintf("あなたは優秀な翻訳者です。入力されたテキストを%sに翻訳してください。元のテキストの意味と文脈を正確に維持してください。", s.lang)
	userPrompt := inputText

	return s.llmClient.Complete(ctx, systemPrompt, userPrompt)
}
