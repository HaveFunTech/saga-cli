package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var examplesCmd = &cobra.Command{
	Use:   "examples",
	Short: "SaGaCLIの使用例を表示します",
	Long:  "SaGaCLIの一般的な使用例とコマンドの例を表示します",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SaGaCLIの使用例:")
		fmt.Println()

		fmt.Println("# ファイルを英語に翻訳する")
		fmt.Println("echo input.txt | saga --translation --lang en")
		fmt.Println()

		fmt.Println("# ファイルを日本語に翻訳する")
		fmt.Println("echo input.txt | saga --translation --lang ja")
		fmt.Println()

		fmt.Println("# ファイルの内容を要約する")
		fmt.Println("echo document.txt | saga --summary --lang en")
		fmt.Println()

		fmt.Println("# ファイルの内容を日本語で解説する")
		fmt.Println("echo code.py | saga --explanation --lang ja")
		fmt.Println()

		fmt.Println("# クエリに基づいて情報を検索する")
		fmt.Println("echo question.txt | saga --search --lang en")
		fmt.Println()

		fmt.Println("# パイプを使って直接テキストを入力")
		fmt.Println("cat README.md | saga --summary --lang ja")
		fmt.Println()

		fmt.Println("# 標準入力からテキストを直接入力")
		fmt.Println(`echo "これはテストです。" | saga --translation --lang en`)
		fmt.Println()

		fmt.Println("# 環境変数設定を確認")
		fmt.Println("saga env")
	},
}

func init() {
	RootCmd.AddCommand(examplesCmd)
}
