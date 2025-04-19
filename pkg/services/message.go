// filepath: /home/soudai/work/hft/saga-cli/pkg/services/message.go
package services

import (
	"context"
	"fmt"

	"github.com/sa-giga/saga-cli/pkg/llm"
)

// MessageService はユーザー指定メッセージに基づいた処理サービスを提供します
type MessageService struct {
	llmClient llm.LLM
	lang      string
	message   string
}

// NewMessageService は新しいメッセージサービスを作成します
func NewMessageService(llmClient llm.LLM, lang string, message string) *MessageService {
	return &MessageService{
		llmClient: llmClient,
		lang:      lang,
		message:   message,
	}
}

// Process はユーザー指定のメッセージに基づいて入力テキストを処理します
func (s *MessageService) Process(ctx context.Context, inputText string) (string, error) {
	systemPrompt := fmt.Sprintf("以下の入力データに対して、指示に従ってタスクを実行してください。回答は%sで提供してください。", s.lang)
	userPrompt := fmt.Sprintf("【指示】\n%s\n\n【入力データ】\n%s", s.message, inputText)

	return s.llmClient.Complete(ctx, systemPrompt, userPrompt)
}
