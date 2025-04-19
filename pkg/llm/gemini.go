// filepath: /home/soudai/work/hft/saga-cli/pkg/llm/gemini.go
package llm

import (
	"context"
	"errors"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// GeminiClient は Gemini API のラッパー
type GeminiClient struct {
	client *genai.Client
	model  string
}

// NewGeminiClient は Gemini クライアントを作成します
func NewGeminiClient() (*GeminiClient, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("GEMINI_API_KEY が設定されていません")
	}

	model := os.Getenv("GEMINI_API_MODEL")
	if model == "" {
		model = "gemini-1.5-pro" // デフォルトモデル
	}

	// Gemini APIクライアントを初期化
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return &GeminiClient{
		client: client,
		model:  model,
	}, nil
}

// Complete はテキスト生成を実行します
func (c *GeminiClient) Complete(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
	// モデルを取得
	model := c.client.GenerativeModel(c.model)

	// チャットセッションを開始
	chat := model.StartChat()

	// システムプロンプトがある場合、システムメッセージとして追加
	if systemPrompt != "" {
		chat.History = append(chat.History, &genai.Content{
			Parts: []genai.Part{genai.Text(systemPrompt)},
			Role:  "system",
		})
	}

	// ユーザープロンプトでレスポンスを生成
	resp, err := chat.SendMessage(ctx, genai.Text(userPrompt))
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", errors.New("レスポンスが空です")
	}

	// テキスト応答を取得
	result, ok := resp.Candidates[0].Content.Parts[0].(genai.Text)
	if !ok {
		return "", errors.New("テキストレスポンスを取得できませんでした")
	}

	return string(result), nil
}
