package cdn

import (
	"fmt"
	"strings"

	"nhncli/cmd"
	"nhncli/internal/cdn"

	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
	Use:   "purge [domain]",
	Short: "캐시 퍼지",
	Args:  cobra.ExactArgs(1),
	RunE:  runPurge,
}

var (
	purgeType  string
	purgeItems string
)

func init() {
	CDNCmd.AddCommand(purgeCmd)

	purgeCmd.Flags().StringVar(&purgeType, "type", "ALL", "퍼지 타입 (ALL, ITEM)")
	purgeCmd.Flags().StringVar(&purgeItems, "items", "", "퍼지 대상 경로 (쉼표 구분)")
}

func runPurge(c *cobra.Command, args []string) error {
	domain := args[0]

	cdnClient, err := cdn.NewClient(cmd.GetProfile(), cmd.GetDebug())
	if err != nil {
		return err
	}

	var items []string
	if purgeItems != "" {
		items = strings.Split(purgeItems, ",")
	}

	if err := cdnClient.Purge(domain, purgeType, items); err != nil {
		return err
	}

	fmt.Printf("캐시 퍼지가 요청되었습니다. (도메인: %s, 타입: %s)\n", domain, purgeType)
	return nil
}
