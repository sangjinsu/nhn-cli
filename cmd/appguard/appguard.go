package appguard

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"nhncli/cmd"
	"nhncli/internal/appguard"

	"github.com/spf13/cobra"
)

var AppGuardCmd = &cobra.Command{
	Use:   "appguard",
	Short: "AppGuard 서비스 관리",
	Long:  "NHN Cloud AppGuard 서비스를 관리합니다.",
}

var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "비정상 행위 탐지 현황 조회",
	RunE:  runDashboard,
}

var (
	targetDate string
	osType     int
	targetType int
)

func init() {
	cmd.GetRootCmd().AddCommand(AppGuardCmd)
	AppGuardCmd.AddCommand(dashboardCmd)
	AppGuardCmd.PersistentFlags().String("app-key", "", "AppGuard AppKey (프로필 설정 오버라이드)")

	dashboardCmd.Flags().StringVar(&targetDate, "target-date", "", "조회 날짜 (YYYY-MM-DD, 필수)")
	dashboardCmd.Flags().IntVar(&osType, "os", 1, "OS 타입 (1: Android, 2: iOS)")
	dashboardCmd.Flags().IntVar(&targetType, "target-type", 0, "대상 타입 (0: 전체, 1: 앱)")

	dashboardCmd.MarkFlagRequired("target-date")
}

func runDashboard(c *cobra.Command, args []string) error {
	appKey, _ := c.Flags().GetString("app-key")
	opts := appguard.ClientOption{AppKey: appKey}
	agClient, err := appguard.NewClient(cmd.GetProfile(), cmd.GetDebug(), opts)
	if err != nil {
		return err
	}

	entries, err := agClient.GetDashboard(targetDate, osType, targetType)
	if err != nil {
		return err
	}

	if cmd.GetOutput() == "json" {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(entries)
	}

	if len(entries) == 0 {
		fmt.Println("탐지 데이터가 없습니다.")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "날짜\t탐지 수\t차단 수")
	fmt.Fprintln(w, "----\t------\t------")
	for _, e := range entries {
		fmt.Fprintf(w, "%s\t%d\t%d\n", e.DetectedDate, e.DetectedCnt, e.BlockedCnt)
	}
	return w.Flush()
}
