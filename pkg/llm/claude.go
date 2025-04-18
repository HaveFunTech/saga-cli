package llm

import (
	"context"
	"encoding/json"
	"errors"
	"os"

	anthropic "github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

// ClaudeClient は Claude API のラッパー
type ClaudeClient struct {
	client anthropic.Client
	model  string
}

// NewClaudeClient は Claude クライアントを作成します
func NewClaudeClient() (*ClaudeClient, error) {
	apiKey := os.Getenv("CLAUDE_API_KEY")
	if apiKey == "" {
		return nil, errors.New("CLAUDE_API_KEY が設定されていません")
	}

	model := os.Getenv("CLAUDE_API_MODEL")
	if model == "" {
		model = "claude-3-haiku-20240307" // デフォルトモデル
	}

	// 最新のSDK APIを使用してクライアントを初期化
	client := anthropic.NewClient(option.WithAPIKey(apiKey))

	return &ClaudeClient{
		client: client,
		model:  model,
	}, nil
}

// Complete はテキスト生成を実行します
func (c *ClaudeClient) Complete(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
	// SDK v0.2.0-beta.3 に合わせたメッセージを作成
	params := anthropic.MessageNewParams{
		Model:     c.model,
		MaxTokens: 1024,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(
				anthropic.NewTextBlock(userPrompt),
			),
		},
	}

	// システムプロンプトを設定（systemPromptを直接指定せず省略）
	// System フィールドは v0.2.0-beta.3 では必須ではない

	resp, err := c.client.Messages.New(ctx, params)
	if err != nil {
		return "", err
	}

	if len(resp.Content) == 0 {
		return "", errors.New("レスポンスが空です")
	}

	// レスポンスからテキストを抽出
	// v0.2.0-beta.3では、ContentBlockUnionを直接処理できないので
	// JSON変換してからマップに復元して処理する
	for i := range resp.Content {
		contentBytes, err := json.Marshal(resp.Content[i])
		if err != nil {
			continue
		}

		var contentMap map[string]interface{}
		if err := json.Unmarshal(contentBytes, &contentMap); err != nil {
			continue
		}

		// typeフィールドがtextであることを確認
		if t, ok := contentMap["type"].(string); ok && t == "text" {
			// テキストフィールドを取得
			if text, ok := contentMap["text"].(string); ok {
				return text, nil
			}
		}
	}

	return "", errors.New("テキストコンテンツが見つかりません")
}
