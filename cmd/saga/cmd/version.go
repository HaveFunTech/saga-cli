package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version   = "0.1.0"
	CommitSHA = "unknown"
	BuildDate = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "SaGaCLIのバージョン情報を表示",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("SaGaCLI バージョン: %s\n", Version)
		fmt.Printf("コミットハッシュ: %s\n", CommitSHA)
		fmt.Printf("ビルド日時: %s\n", BuildDate)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
