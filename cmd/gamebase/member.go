package gamebase

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"nhncli/cmd"
	gb "nhncli/internal/gamebase"

	"github.com/spf13/cobra"
)

var memberCmd = &cobra.Command{
	Use:   "member",
	Short: "회원 관리",
}

var memberDescribeCmd = &cobra.Command{
	Use:   "describe [user-id]",
	Short: "회원 조회",
	Args:  cobra.ExactArgs(1),
	RunE:  runMemberDescribe,
}

var memberListCmd = &cobra.Command{
	Use:   "list",
	Short: "회원 일괄 조회",
	RunE:  runMemberList,
}

var memberWithdrawCmd = &cobra.Command{
	Use:   "withdraw [user-id]",
	Short: "회원 탈퇴",
	Args:  cobra.ExactArgs(1),
	RunE:  runMemberWithdraw,
}

var userIDs string

func init() {
	GamebaseCmd.AddCommand(memberCmd)
	memberCmd.AddCommand(memberDescribeCmd)
	memberCmd.AddCommand(memberListCmd)
	memberCmd.AddCommand(memberWithdrawCmd)

	memberListCmd.Flags().StringVar(&userIDs, "user-ids", "", "사용자 ID 목록 (쉼표 구분, 필수)")
	memberListCmd.MarkFlagRequired("user-ids")
}

func runMemberDescribe(c *cobra.Command, args []string) error {
	appKey, _ := c.Flags().GetString("app-key")
	secretKey, _ := c.Flags().GetString("secret-key")
	opts := gb.ClientOption{AppKey: appKey, SecretKey: secretKey}
	gbClient, err := gb.NewClient(cmd.GetProfile(), cmd.GetDebug(), opts)
	if err != nil {
		return err
	}

	member, err := gbClient.GetMember(args[0])
	if err != nil {
		return err
	}

	if cmd.GetOutput() == "json" {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(member)
	}

	fmt.Printf("User ID: %s\n", member.UserID)
	fmt.Printf("Valid: %s\n", member.Valid)
	fmt.Printf("App ID: %s\n", member.AppID)
	fmt.Printf("등록일: %s\n", member.RegDate)
	fmt.Printf("최근 로그인: %s\n", member.LastLoginDate)
	return nil
}

func runMemberList(c *cobra.Command, args []string) error {
	appKey, _ := c.Flags().GetString("app-key")
	secretKey, _ := c.Flags().GetString("secret-key")
	opts := gb.ClientOption{AppKey: appKey, SecretKey: secretKey}
	gbClient, err := gb.NewClient(cmd.GetProfile(), cmd.GetDebug(), opts)
	if err != nil {
		return err
	}

	ids := strings.Split(userIDs, ",")
	members, err := gbClient.ListMembers(ids)
	if err != nil {
		return err
	}

	if cmd.GetOutput() == "json" {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(members)
	}

	if len(members) == 0 {
		fmt.Println("회원이 없습니다.")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "User ID\t상태\t등록일\t최근 로그인")
	fmt.Fprintln(w, "-------\t----\t-----\t---------")
	for _, m := range members {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", m.UserID, m.Valid, m.RegDate, m.LastLoginDate)
	}
	return w.Flush()
}

func runMemberWithdraw(c *cobra.Command, args []string) error {
	appKey, _ := c.Flags().GetString("app-key")
	secretKey, _ := c.Flags().GetString("secret-key")
	opts := gb.ClientOption{AppKey: appKey, SecretKey: secretKey}
	gbClient, err := gb.NewClient(cmd.GetProfile(), cmd.GetDebug(), opts)
	if err != nil {
		return err
	}

	if err := gbClient.WithdrawMember(args[0]); err != nil {
		return err
	}

	fmt.Printf("회원 '%s'이(가) 탈퇴 처리되었습니다.\n", args[0])
	return nil
}
