package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "CLI 버전 정보 출력",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("nhn version %s (built: %s)\n", Version, BuildDate)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
