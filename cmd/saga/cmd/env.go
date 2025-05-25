package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "環境変数の設定方法を表示",
	Long:  "SaGaCLIで使用する環境変数の設定方法を表示します。",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SaGaCLIで使用する環境変数:")
		fmt.Println()

		fmt.Println("# 使用するAPIタイプの選択（デフォルトはopenai）")
		fmt.Println("export OPENAI_API_TYPE=openai  # OpenAIを使用")
		fmt.Println("export OPENAI_API_TYPE=claude  # Claudeを使用")
		fmt.Println("export OPENAI_API_TYPE=gemini  # Geminiを使用")
		fmt.Println("export OPENAI_API_TYPE=mcp     # MCPサーバーを使用")
		fmt.Println()

		fmt.Println("# OpenAI API設定")
		fmt.Println("export OPENAI_API_KEY=sk-xxxxxxxxxxxx  # APIキーを設定")
		fmt.Println("export OPENAI_API_BASE_URL=https://api.openai.com/v1  # カスタムAPIエンドポイント（オプション）")
		fmt.Println("export OPENAI_API_MODEL=gpt-3.5-turbo  # 使用するモデル（デフォルトはgpt-3.5-turbo）")
		fmt.Println()

		fmt.Println("# Claude API設定")
		fmt.Println("export CLAUDE_API_KEY=sk-xxxxxxxxxxxx  # Anthropic APIキーを設定")
		fmt.Println("export CLAUDE_API_MODEL=claude-3-haiku-20240307  # 使用するモデル（デフォルトはclaude-3-haiku-20240307）")
		fmt.Println()

		fmt.Println("# Gemini API設定")
		fmt.Println("export GEMINI_API_KEY=your-api-key  # Gemini APIキーを設定")
		fmt.Println("export GEMINI_API_MODEL=gemini-1.5-pro  # 使用するモデル（デフォルトはgemini-1.5-pro）")
		fmt.Println()

		fmt.Println("# MCP (Model Context Protocol) 設定")
		fmt.Println("export MCP_CONFIG_FILE=/path/to/mcp-config.json  # MCP設定ファイルのパス")
		fmt.Println("export MCP_SERVER_NAME=server-name  # 使用するMCPサーバー名")
	},
}

func init() {
	RootCmd.AddCommand(envCmd)
}
