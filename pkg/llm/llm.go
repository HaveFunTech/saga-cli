package llm

import (
	"context"
	"errors"
	"os"
)

// LLM はAI言語モデルのインターフェース
type LLM interface {
	Complete(ctx context.Context, systemPrompt, userPrompt string) (string, error)
}

// GetLLM は設定に基づいて適切なLLMクライアントを返します
func GetLLM() (LLM, error) {
	// APIタイプに基づいて適切なLLMクライアントを返す
	apiType := os.Getenv("OPENAI_API_TYPE")

	switch apiType {
	case "openai", "":
		return NewOpenAIClient()
	case "claude":
		return NewClaudeClient()
	case "gemini":
		return NewGeminiClient()
	case "mcp":
		mcpConfigFile := os.Getenv("MCP_CONFIG_FILE")
		mcpServerName := os.Getenv("MCP_SERVER_NAME")
		return NewMCPClient(mcpConfigFile, mcpServerName)
	default:
		return nil, errors.New("サポートされていないAPI_TYPEです: " + apiType)
	}
}
