package llm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// MCPConfig はMCPサーバーの設定情報
type MCPConfig struct {
	MCPServers map[string]MCPServerConfig `json:"mcpServers"`
}

// MCPServerConfig はMCPサーバーの設定詳細
type MCPServerConfig struct {
	Command string            `json:"command"`
	Args    []string          `json:"args"`
	Env     map[string]string `json:"env"`
}

// MCPClient はAnthropicのMCP (Model Control Plane) クライアント実装
type MCPClient struct {
	configFile string
	serverName string
	config     MCPConfig
}

// NewMCPClient は新しいMCPクライアントを作成します
func NewMCPClient(configFile, serverName string) (*MCPClient, error) {
	if configFile == "" {
		return nil, errors.New("MCP config file path is not specified")
	}

	// ファイルが存在するか確認
	_, err := os.Stat(configFile)
	if err != nil {
		return nil, fmt.Errorf("MCP config file does not exist: %v", err)
	}

	// 設定ファイルを読み込む
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read MCP config file: %v", err)
	}

	var config MCPConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse MCP config file: %v", err)
	}

	// 指定されたサーバーが存在するか確認
	if serverName != "" {
		if _, exists := config.MCPServers[serverName]; !exists {
			return nil, fmt.Errorf("specified MCP server '%s' not found in config", serverName)
		}
	}

	return &MCPClient{
		configFile: configFile,
		serverName: serverName,
		config:     config,
	}, nil
}

// Complete はMCPを通してテキスト補完を実行します
func (c *MCPClient) Complete(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
	// サーバーが指定されていない場合はエラー
	if c.serverName == "" {
		return "", errors.New("MCP server name is not specified")
	}

	// 指定されたサーバーの設定を取得
	serverConfig, exists := c.config.MCPServers[c.serverName]
	if !exists {
		return "", fmt.Errorf("MCP server '%s' not found in config", c.serverName)
	}

	// 一時ファイルを作成してプロンプトを書き込む
	tempDir, err := os.MkdirTemp("", "saga-mcp")
	if err != nil {
		return "", fmt.Errorf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	promptFile := filepath.Join(tempDir, "prompt.txt")
	content := ""
	if systemPrompt != "" {
		content += fmt.Sprintf("System: %s\n\n", systemPrompt)
	}
	content += fmt.Sprintf("User: %s\n\n", userPrompt)

	if err := os.WriteFile(promptFile, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("failed to write prompt file: %v", err)
	}

	// MCP サーバーコマンドを実行
	cmd := exec.CommandContext(ctx, serverConfig.Command, append(serverConfig.Args, promptFile)...)
	
	// 環境変数を設定
	cmd.Env = os.Environ()
	for key, value := range serverConfig.Env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
	}

	// 出力を取得
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute MCP server: %v, output: %s", err, string(output))
	}

	return string(output), nil
}