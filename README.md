# SaGaCLI

SaGaCLI（サガCLI）は、LLM（Large Language Model）を活用するためのシンプルで強力なコマンドラインインターフェースです。OpenAIのGPT-3.5/GPT-4またはAnthropicのClaudeモデルを使用して、テキストの翻訳、要約、解説、検索などの機能を提供します。

## 特徴

- **複数のLLMサポート**: OpenAI（GPT-3.5/GPT-4）とClaude（Anthropic）の両方をサポート
- **多様な機能**: 翻訳、要約、解説、検索など、さまざまな機能を提供
- **柔軟な入力**: ファイルパスまたは標準入力からテキストを受け付け
- **多言語対応**: 出力言語を自由に指定可能

## インストール

### 前提条件

- Go 1.16以上

### ソースからのインストール

```bash
# リポジトリをクローン
git clone https://github.com/sa-giga/saga-cli.git
cd saga-cli

# ビルドとインストール
make install
```

## 使い方

SaGaCLIを使用する前に、必要な環境変数を設定する必要があります。

### 環境変数の設定

環境変数の設定方法については、`saga env`コマンドで詳細を確認できます。

```bash
# OpenAI APIを使用する場合
export OPENAI_API_KEY=your_api_key
export OPENAI_API_MODEL=gpt-3.5-turbo  # オプション、デフォルトはgpt-3.5-turbo

# Claude APIを使用する場合
export OPENAI_API_TYPE=claude
export CLAUDE_API_KEY=your_api_key
export CLAUDE_API_MODEL=claude-3-haiku-20240307  # オプション
```

### 基本的な使い方

```bash
# ファイルの内容を翻訳
echo input.txt | saga --translation --lang ja

# ファイルの内容を要約
echo document.txt | saga --summary --lang en

# ファイルの内容を解説
echo code.py | saga --explanation --lang ja

# クエリに基づいて情報を検索
echo query.txt | saga --search --lang en

# パイプを使用して直接テキストを入力
cat README.md | saga --summary --lang ja
```

より多くの使用例については、`saga examples`コマンドで確認できます。

## 機能一覧

- **翻訳** (`--translation`): 指定された言語にテキストを翻訳します
- **要約** (`--summary`): テキストの内容を簡潔にまとめます
- **解説** (`--explanation`): テキストの内容について詳しく説明します
- **検索** (`--search`): クエリに基づいて情報を提供します

## ライセンス

MIT License
