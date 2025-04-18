package llm

import (
	"context"
	"errors"
	"os"

	"github.com/sashabaranov/go-openai"
)

// OpenAIClient は OpenAI API のラッパー
type OpenAIClient struct {
	client *openai.Client
	model  string
}

// NewOpenAIClient は OpenAI クライアントを作成します
func NewOpenAIClient() (*OpenAIClient, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("OPENAI_API_KEY が設定されていません")
	}

	baseURL := os.Getenv("OPENAI_API_BASE_URL")

	model := os.Getenv("OPENAI_API_MODEL")
	if model == "" {
		model = "gpt-3.5-turbo" // デフォルトモデル
	}

	config := openai.DefaultConfig(apiKey)
	if baseURL != "" {
		config.BaseURL = baseURL
	}

	client := openai.NewClientWithConfig(config)
	return &OpenAIClient{
		client: client,
		model:  model,
	}, nil
}

// Complete はテキスト生成を実行します
func (c *OpenAIClient) Complete(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
	resp, err := c.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: c.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userPrompt,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("レスポンスが空です")
	}

	return resp.Choices[0].Message.Content, nil
}
