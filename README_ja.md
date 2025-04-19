# SaGaCLI
Golangで書かれたコマンドラインインターフェースです。
コマンドラインインターフェースを使用して、LLMを利用できます。

## 特徴

- **複数のLLMサポート**: OpenAI（GPT-3.5/GPT-4）、AnthropicのClaude（Claude-3シリーズ）、GoogleのGemini（Gemini-2.5-pro）をサポート
- **多様な機能**: 翻訳、要約、解説、検索、カスタムメッセージなど、さまざまな機能を提供
- **柔軟な入力**: ファイルパスまたは標準入力からテキストを受け付け
- **多言語対応**: 出力言語を自由に指定可能

## インストール方法

### 前提条件

- Go 1.16以上

### 方法1: バイナリのインストール（推奨）

```bash
# Goのインストール（未インストールの場合）
$ brew install go  # macOS
# または
$ sudo apt install golang-go  # Ubuntu/Debian

# SaGaCLIのインストール 
$ go install github.com/sa-giga/saga-cli@latest

# SaGaCLIのパスを通す
$ export PATH=$PATH:$(go env GOPATH)/bin
```

### 方法2: ソースコードからのビルド

```bash
# リポジトリをクローン
$ git clone https://github.com/sa-giga/saga-cli.git
$ cd saga-cli

# ビルドとインストール
$ make build  # bin/sagaにバイナリが作成されます
$ make install  # システムにインストールします
```

## 詳細な環境設定

SaGaCLIは、OpenAIのGPT-3.5/GPT-4、AnthropicのClaudeモデル、またはGoogleのGeminiモデルを使用できます。使用したいサービスに応じて適切な環境変数を設定してください。

### OpenAI APIを使用する場合

```bash
# 必須の設定
export OPENAI_API_KEY=sk-...  # OpenAI APIキー
export OPENAI_API_TYPE=openai  # APIタイプを指定（デフォルト値）

# オプションの設定
export OPENAI_API_BASE_URL=https://api.openai.com/v1  # APIのベースURL（デフォルト値）
export OPENAI_API_VERSION=2023-05-15  # APIバージョン
export OPENAI_API_MODEL=gpt-3.5-turbo  # 使用するモデル（デフォルト: gpt-3.5-turbo）

# Azure OpenAI Serviceを使用する場合
export OPENAI_API_TYPE=azure
export OPENAI_API_BASE_URL=https://your-resource-name.openai.azure.com
export OPENAI_API_KEY=your-azure-api-key
export OPENAI_API_VERSION=2023-05-15
export OPENAI_API_MODEL=your-deployment-name
```

### Anthropic Claude APIを使用する場合

```bash
# 必須の設定
export CLAUDE_API_KEY=sk-ant-...  # Anthropic APIキー
export OPENAI_API_TYPE=claude  # APIタイプをclaudeに設定

# オプションの設定
export CLAUDE_API_MODEL=claude-3-haiku-20240307  # 使用するモデル（デフォルト値）
# 利用可能なモデル: claude-3-opus-20240229, claude-3-sonnet-20240229, claude-3-haiku-20240307 など
```

### Google Gemini APIを使用する場合

```bash
# 必須の設定
export GEMINI_API_KEY=your-gemini-api-key  # Gemini APIキー
export OPENAI_API_TYPE=gemini  # APIタイプをgeminiに設定

# オプションの設定
export GEMINI_API_MODEL=gemini-1.5-pro  # 使用するモデル（デフォルト値）
# 利用可能なその他のモデル: gemini-1.5-flash など
```

環境変数の設定が確認できない場合は、`saga env`コマンドを実行して現在の設定を確認できます。

## 使い方の詳細

### 基本的な使用方法

```bash
# ファイルの内容を処理
$ saga [オプション] < 入力ファイル

# または
$ cat 入力ファイル | saga [オプション]

# 標準入力から直接テキストを入力（Ctrl+Dで入力終了）
$ saga [オプション]
テキストを入力...
[Ctrl+D]
```

### 主要なオプション

```bash
--translation  # 翻訳モード
--summary      # 要約モード
--explanation  # 解説モード
--search       # 検索モード
--message      # カスタムメッセージモード（特定の指示を送信）
--lang [言語コード]  # 出力言語の指定（例: ja, en, fr, zh など）
```

### 使用例

```bash
# 英語のテキストを日本語に翻訳
$ echo "Hello world" | saga --translation --lang ja

# 長い文書を日本語で要約
$ cat document.txt | saga --summary --lang ja

# プログラムコードを解説
$ cat code.py | saga --explanation --lang ja

# 特定のトピックについて検索
$ echo "量子コンピューティングの基本原理" | saga --search --lang ja

# JSONデータからnameフィールドの値を抽出
$ cat document.json | saga --message "nameの値を取り出して" --lang ja

# CSVファイルの年齢列の平均を計算
$ cat document.csv | saga --message "年齢の列の平均を計算して" --lang ja

# 英語のドキュメントを翻訳して日本語で要約
$ cat english_doc.txt | saga --translation --lang ja | saga --summary --lang ja
```

### 使用例の一覧表示

```bash
# さまざまな使用例を確認
$ saga examples
```

## トラブルシューティング

### APIキーが設定されていないエラー

```
Error: OPENAI_API_KEY または CLAUDE_API_KEY または GEMINI_API_KEY が設定されていません
```

上記のエラーが発生した場合は、環境変数が正しく設定されているか確認してください。

### モデルの選択

処理が遅い場合や、より高品質な結果が必要な場合は、環境変数でより高性能なモデルを選択してください：

```bash
# OpenAIの場合
export OPENAI_API_MODEL=gpt-4

# Claudeの場合
export CLAUDE_API_MODEL=claude-3-opus-20240229

# Geminiの場合
export GEMINI_API_MODEL=gemini-1.5-pro
```

## ヘルプの表示

```bash
# コマンドヘルプの表示
$ saga --help

# バージョン情報の表示
$ saga --version

# 環境設定の確認
$ saga env
```

