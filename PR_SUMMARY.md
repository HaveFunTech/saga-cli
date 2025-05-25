# MCP (Model Context Protocol) 対応のPR

このPRでは、saga-cliにMCP (Model Context Protocol) サーバーを使用するための機能を追加しました。MCPサーバーを利用することで、専門的な機能を持つAIモデルにアクセスできるようになります。

## 主な変更点

1. **MCP対応クライアントの実装**
   - `/pkg/llm/mcp.go`: MCPクライアント実装
   - `/pkg/llm/llm.go`: GetLLM関数にMCPタイプ対応を追加

2. **コマンドラインオプションの追加**
   - `/cmd/saga/cmd/root.go`:
     - `--mcp-config`: MCP設定ファイルのパスを指定するオプション
     - `--mcp-server`: 使用するMCPサーバー名を指定するオプション

3. **ドキュメント更新**
   - `README.md`: MCP機能の使い方と設定方法を追加（英語版）
   - `README_ja.md`: MCP機能の使い方と設定方法を追加（日本語版）

4. **環境変数表示の更新**
   - `/cmd/saga/cmd/env.go`: MCP関連の環境変数表示を追加

5. **サンプルとテスト用ファイル**
   - `/examples/mcp-config.json`: MCP設定ファイルのサンプル
   - `/examples/mcp-test/`: テスト用ファイル

## 使用方法

MCPサーバーを使用するには、以下のいずれかの方法で設定します：

### コマンドラインオプションで指定

```bash
saga --mcp-config ./mcp-config.json --mcp-server server-name [その他のオプション]
```

### 環境変数で指定

```bash
export OPENAI_API_TYPE=mcp
export MCP_CONFIG_FILE=/path/to/mcp-config.json
export MCP_SERVER_NAME=server-name
saga [その他のオプション]
```

## 設定ファイル例

MCPサーバーの設定は以下のようなJSON形式で定義します：

```json
{
  "mcpServers": {
    "google-maps": {
      "command": "npx",
      "args": [
        "-y",
        "@modelcontextprotocol/server-google-maps"
      ],
      "env": {
        "GOOGLE_MAPS_API_KEY": "<YOUR_API_KEY>"
      }
    }
  }
}
```

## 今後の課題

- MCPサーバーとの通信エラー処理の改善
- 複数のMCPサーバーを同時に使用する機能