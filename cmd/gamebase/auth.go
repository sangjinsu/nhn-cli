package gamebase

import (
	"fmt"

	"nhncli/cmd"
	gb "nhncli/internal/gamebase"

	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "인증 관리",
}

var authValidateCmd = &cobra.Command{
	Use:   "validate [user-id] [access-token]",
	Short: "토큰 검증",
	Args:  cobra.ExactArgs(2),
	RunE:  runAuthValidate,
}

func init() {
	GamebaseCmd.AddCommand(authCmd)
	authCmd.AddCommand(authValidateCmd)
}

func runAuthValidate(c *cobra.Command, args []string) error {
	gbClient, err := gb.NewClient(cmd.GetProfile(), cmd.GetDebug())
	if err != nil {
		return err
	}

	valid, err := gbClient.ValidateToken(args[0], args[1])
	if err != nil {
		return err
	}

	if valid {
		fmt.Println("토큰이 유효합니다.")
	} else {
		fmt.Println("토큰이 유효하지 않습니다.")
	}
	return nil
}
