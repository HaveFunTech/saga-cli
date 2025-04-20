package cmd

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/sa-giga/saga-cli/pkg/llm"
	"github.com/sa-giga/saga-cli/pkg/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// フラグ変数
	isTranslation bool
	isSummary     bool
	isExplanation bool
	isSearch      bool
	messageFlag   string
	langFlag      string

	fileFlags []string
	dirFlags  []string

	// RootCmd はCLIツールのルートコマンド
	RootCmd = &cobra.Command{
		Use:   "saga",
		Short: "SaGaCLI - LLMを活用するコマンドラインツール",
		Long: `SaGaCLI は翻訳、要約、解説、検索などのタスクにLLMを活用するコマンドラインツールです。
環境変数を設定して、OpenAIやClaudeのAPIを利用することができます。`,
		Run: rootRun,
	}
)

func init() {
	// フラグの設定
	RootCmd.PersistentFlags().BoolVarP(&isTranslation, "translation", "t", false, "入力されたテキストを翻訳します")
	RootCmd.PersistentFlags().BoolVarP(&isSummary, "summary", "s", false, "入力されたテキストを要約します")
	RootCmd.PersistentFlags().BoolVarP(&isExplanation, "explanation", "e", false, "入力されたテキストを解説します")
	RootCmd.PersistentFlags().BoolVarP(&isSearch, "search", "S", false, "入力されたテキストに基づいて情報を検索します")
	RootCmd.PersistentFlags().StringVarP(&messageFlag, "message", "m", "", "入力データに対して実行したい操作を指定します")
	RootCmd.PersistentFlags().StringVarP(&langFlag, "lang", "l", "en", "出力言語を指定します（例: en, ja, fr）")
	RootCmd.PersistentFlags().StringSliceVarP(&fileFlags, "file", "f", nil, "コンテキストとして利用するファイルを指定します（複数指定可）")
	RootCmd.PersistentFlags().StringSliceVarP(&dirFlags, "dir", "d", nil, "コンテキストとして利用するディレクトリを指定します（再帰的に全ファイルを読み込み）")

	// 環境変数の設定を読み込む
	viper.AutomaticEnv()
}

// ディレクトリ内の全ファイルを再帰的に取得
func getAllFilesRecursive(dir string) ([]string, error) {
	var files []string
	err := walkDir(dir, &files)
	return files, err
}

func walkDir(dir string, files *[]string) error {
	d, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range d {
		path := dir + string(os.PathSeparator) + entry.Name()
		if entry.IsDir() {
			if err := walkDir(path, files); err != nil {
				return err
			}
		} else {
			*files = append(*files, path)
		}
	}
	return nil
}

// ファイルリストから内容を結合
func readFilesContent(filePaths []string) (string, error) {
	var buf bytes.Buffer
	for _, f := range filePaths {
		content, err := services.ReadFileContent(f)
		if err != nil {
			return "", err
		}
		buf.WriteString("【ファイル: " + f + "】\n")
		buf.WriteString(content)
		buf.WriteString("\n\n")
	}
	return buf.String(), nil
}

func rootRun(cmd *cobra.Command, args []string) {
	var content string
	var err error

	// --- コンテキストファイル・ディレクトリの内容を先に読み込む ---
	var contextContents string
	var contextFiles []string
	if len(fileFlags) > 0 {
		contextFiles = append(contextFiles, fileFlags...)
	}
	if len(dirFlags) > 0 {
		for _, dir := range dirFlags {
			files, err := getAllFilesRecursive(dir)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ディレクトリの読み取りに失敗しました: %v\n", err)
				os.Exit(1)
			}
			contextFiles = append(contextFiles, files...)
		}
	}
	if len(contextFiles) > 0 {
		contextContents, err = readFilesContent(contextFiles)
		if err != nil {
			fmt.Fprintf(os.Stderr, "コンテキストファイルの読み取りに失敗しました: %v\n", err)
			os.Exit(1)
		}
	}

	// 標準入力が端末から来ているか確認
	stat, _ := os.Stdin.Stat()
	isPipe := (stat.Mode() & os.ModeCharDevice) == 0

	if isPipe {
		// パイプまたはリダイレクトからの入力
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, os.Stdin); err != nil {
			fmt.Fprintf(os.Stderr, "標準入力の読み取りに失敗しました: %v\n", err)
			os.Exit(1)
		}
		content = buf.String()
	} else {
		// 対話モード: ファイルパスを読み取る
		scanner := bufio.NewScanner(os.Stdin)
		var filePath string
		if scanner.Scan() {
			filePath = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "入力の読み取りに失敗しました: %v\n", err)
			os.Exit(1)
		}

		// ファイルの内容を読み込む
		content, err = services.ReadFileContent(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ファイルの読み取りに失敗しました: %v\n", err)
			os.Exit(1)
		}
	}

	// --- コンテキストをuserPromptの先頭に付与 ---
	if contextContents != "" {
		content = "【コンテキスト】\n" + contextContents + "\n【入力データ】\n" + content
	}

	// LLMクライアントの取得
	llmClient, err := llm.GetLLM()
	if err != nil {
		fmt.Fprintf(os.Stderr, "LLMクライアントの初期化に失敗しました: %v\n", err)
		os.Exit(1)
	}

	ctx := context.Background()
	var service services.Service

	// フラグに基づいて適切なサービスを選択
	switch {
	case isTranslation:
		service = services.NewTranslationService(llmClient, langFlag)
	case isSummary:
		service = services.NewSummaryService(llmClient, langFlag)
	case isExplanation:
		service = services.NewExplanationService(llmClient, langFlag)
	case isSearch:
		service = services.NewSearchService(llmClient, langFlag)
	case messageFlag != "":
		service = services.NewMessageService(llmClient, langFlag, messageFlag)
	default:
		fmt.Fprintf(os.Stderr, "機能フラグが指定されていません。--translation, --summary, --explanation, --search, --messageのいずれかを指定してください。\n")
		os.Exit(1)
	}

	// サービスを実行
	result, err := service.Process(ctx, content)
	if err != nil {
		fmt.Fprintf(os.Stderr, "処理に失敗しました: %v\n", err)
		os.Exit(1)
	}

	// 結果を表示
	fmt.Println(result)
}
