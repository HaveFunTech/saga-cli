package services

import (
	"context"
	"fmt"

	"github.com/sa-giga/saga-cli/pkg/llm"
)

// SearchService は検索サービスを提供します
type SearchService struct {
	llmClient llm.LLM
	lang      string
}

// NewSearchService は新しい検索サービスを作成します
func NewSearchService(llmClient llm.LLM, lang string) *SearchService {
	return &SearchService{
		llmClient: llmClient,
		lang:      lang,
	}
}

// Process は入力テキストに基づいて情報を検索します
func (s *SearchService) Process(ctx context.Context, inputText string) (string, error) {
	systemPrompt := fmt.Sprintf("あなたはリサーチの専門家です。入力されたクエリに関する情報を%sで提供してください。関連性の高い事実に基づいた情報を提供し、必要に応じて複数の視点を含めてください。", s.lang)
	userPrompt := inputText

	return s.llmClient.Complete(ctx, systemPrompt, userPrompt)
}
