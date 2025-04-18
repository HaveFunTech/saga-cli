package services

import (
	"context"
	"os"
	"strings"
)

// Service はLLMを使用したサービスのインターフェース
type Service interface {
	Process(ctx context.Context, inputText string) (string, error)
}

// ReadFileContent はファイルパスからコンテンツを読み込みます
func ReadFileContent(filePath string) (string, error) {
	// ファイルパスからコンテンツを読み込む
	content, err := os.ReadFile(strings.TrimSpace(filePath))
	if err != nil {
		return "", err
	}
	return string(content), nil
}
