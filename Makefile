.PHONY: build clean test deps install

# ビルド変数
BINARY_NAME=saga
VERSION=$(shell git describe --tags --always || echo "dev")
COMMIT_SHA=$(shell git rev-parse HEAD || echo "unknown")
BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LD_FLAGS=-X github.com/sa-giga/saga-cli/cmd/saga/cmd.Version=$(VERSION) -X github.com/sa-giga/saga-cli/cmd/saga/cmd.CommitSHA=$(COMMIT_SHA) -X github.com/sa-giga/saga-cli/cmd/saga/cmd.BuildDate=$(BUILD_DATE)

# デフォルトゴール
all: deps test build

# 依存関係の取得
deps:
	@echo "依存関係を取得中..."
	go mod tidy
	go get -u github.com/sashabaranov/go-openai
	go get -u github.com/spf13/cobra
	go get -u github.com/spf13/viper
	go get -u github.com/anthropics/anthropic-sdk-go

# テストの実行
test:
	@echo "テストを実行中..."
	go test -v ./...

# アプリケーションのビルド
build:
	@echo "$(BINARY_NAME)をビルド中..."
	go build -ldflags "$(LD_FLAGS)" -o bin/$(BINARY_NAME) ./cmd/saga

# クリーンアップ
clean:
	@echo "クリーンアップ中..."
	rm -rf bin/
	go clean

# インストール
install: build
	@echo "$(BINARY_NAME)をインストール中..."
	go install -ldflags "$(LD_FLAGS)" ./cmd/saga