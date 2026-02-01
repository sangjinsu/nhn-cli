package gamebase

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"nhncli/cmd"
	gb "nhncli/internal/gamebase"

	"github.com/spf13/cobra"
)

var banCmd = &cobra.Command{
	Use:   "ban",
	Short: "이용 정지 관리",
}

var banCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "이용 정지",
	RunE:  runBanCreate,
}

var banListCmd = &cobra.Command{
	Use:   "list",
	Short: "이용 정지 목록 조회",
	RunE:  runBanList,
}

var banReleaseCmd = &cobra.Command{
	Use:   "release",
	Short: "이용 정지 해제",
	RunE:  runBanRelease,
}

var (
	banUserID    string
	banType      string
	banBeginDate string
	banEndDate   string
	banReason    string
	banMessage   string
)

func init() {
	GamebaseCmd.AddCommand(banCmd)
	banCmd.AddCommand(banCreateCmd)
	banCmd.AddCommand(banListCmd)
	banCmd.AddCommand(banReleaseCmd)

	banCreateCmd.Flags().StringVar(&banUserID, "user-id", "", "사용자 ID (필수)")
	banCreateCmd.Flags().StringVar(&banType, "type", "", "정지 타입")
	banCreateCmd.Flags().StringVar(&banBeginDate, "begin-date", "", "시작일 (필수)")
	banCreateCmd.Flags().StringVar(&banEndDate, "end-date", "", "종료일 (필수)")
	banCreateCmd.Flags().StringVar(&banReason, "reason", "", "사유")
	banCreateCmd.Flags().StringVar(&banMessage, "message", "", "메시지")
	banCreateCmd.MarkFlagRequired("user-id")
	banCreateCmd.MarkFlagRequired("begin-date")
	banCreateCmd.MarkFlagRequired("end-date")

	banListCmd.Flags().StringVar(&banUserID, "user-id", "", "사용자 ID (필수)")
	banListCmd.MarkFlagRequired("user-id")

	banReleaseCmd.Flags().StringVar(&banUserID, "user-id", "", "사용자 ID (필수)")
	banReleaseCmd.MarkFlagRequired("user-id")
}

func runBanCreate(c *cobra.Command, args []string) error {
	appKey, _ := c.Flags().GetString("app-key")
	secretKey, _ := c.Flags().GetString("secret-key")
	opts := gb.ClientOption{AppKey: appKey, SecretKey: secretKey}
	gbClient, err := gb.NewClient(cmd.GetProfile(), cmd.GetDebug(), opts)
	if err != nil {
		return err
	}

	req := &gb.BanCreateRequest{
		UserID:    banUserID,
		BanType:   banType,
		BeginDate: banBeginDate,
		EndDate:   banEndDate,
		Reason:    banReason,
		Message:   banMessage,
	}

	if err := gbClient.CreateBan(req); err != nil {
		return err
	}

	fmt.Printf("사용자 '%s'이(가) 이용 정지되었습니다.\n", banUserID)
	return nil
}

func runBanList(c *cobra.Command, args []string) error {
	appKey, _ := c.Flags().GetString("app-key")
	secretKey, _ := c.Flags().GetString("secret-key")
	opts := gb.ClientOption{AppKey: appKey, SecretKey: secretKey}
	gbClient, err := gb.NewClient(cmd.GetProfile(), cmd.GetDebug(), opts)
	if err != nil {
		return err
	}

	bans, err := gbClient.ListBans(banUserID)
	if err != nil {
		return err
	}

	if cmd.GetOutput() == "json" {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(bans)
	}

	if len(bans) == 0 {
		fmt.Println("이용 정지 내역이 없습니다.")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "User ID\t타입\t시작일\t종료일\t사유")
	fmt.Fprintln(w, "-------\t----\t-----\t-----\t----")
	for _, b := range bans {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", b.UserID, b.BanType, b.BeginDate, b.EndDate, b.Reason)
	}
	return w.Flush()
}

func runBanRelease(c *cobra.Command, args []string) error {
	appKey, _ := c.Flags().GetString("app-key")
	secretKey, _ := c.Flags().GetString("secret-key")
	opts := gb.ClientOption{AppKey: appKey, SecretKey: secretKey}
	gbClient, err := gb.NewClient(cmd.GetProfile(), cmd.GetDebug(), opts)
	if err != nil {
		return err
	}

	if err := gbClient.ReleaseBan(banUserID); err != nil {
		return err
	}

	fmt.Printf("사용자 '%s'의 이용 정지가 해제되었습니다.\n", banUserID)
	return nil
}
